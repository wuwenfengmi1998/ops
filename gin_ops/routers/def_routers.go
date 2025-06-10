package routers

import (
	"fmt"
	"net/http"
	"saas/models"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
)

func Def_router(r *gin.RouterGroup) {
	r.Use(func(ctx *gin.Context) {

		cookie_vel := ""
		//读取用户cookie，判断用户是否已登录
		cookie_s, _ := ctx.Cookie("user")
		cookie_vel = cookie_s
		//转换传进来的数据
		var jsonData map[string]interface{}
		if err := ctx.ShouldBindJSON(&jsonData); err == nil {
			//分离数据
			var cookie_t models.Cookie
			if err = mapstructure.Decode(jsonData["cookie"], &cookie_t); err == nil {
				cookie_vel = cookie_t.Value
			}
			var data_t map[string]interface{}
			if err = mapstructure.Decode(jsonData["data"], &data_t); err == nil {
				ctx.Set("data", &data_t)
			}

		}

		//fmt.Println(cookie_vel)
		if cookie_vel != "" {
			var cookie models.Cookie
			cookie.Value = cookie_vel
			if models.DB.Where(&cookie).First(&cookie).Error == nil {
				// 有数据
				//有cookie，判断cookie有效性
				if cookie.ExpiresAt.After(time.Now()) {
					// ExpiresAt 在当前时间之后（未过期）
					//fmt.Println("Cookie 未过期")

					//每次调用都更新cookie的最新状态 ，用于计算在线
					var cookie_up models.Cookie
					cookie_up.UpdatedAt = time.Now()
					cookie_up.ExpiresAt = time.Now().Add(time.Duration(models.User_configs["cookie_timeout"].(int)) * time.Second) //计算过期时间
					models.DB.Model(&models.Cookie{}).Where(&cookie).Updates(&cookie_up)
					//更新前端cookie
					ctx.SetCookie("user", cookie.Value, models.User_configs["cookie_timeout"].(int), "/", models.Wed_configs.Host, models.Wed_configs.Tls, true)
					cookie.UpdatedAt = cookie_up.UpdatedAt
					cookie.ExpiresAt = cookie_up.ExpiresAt
					ctx.Set("cookie", cookie)
					//读取用户权限信息
					var user models.User
					user.ID = cookie.UserID
					if models.DB.Where(&user).First(&user).Error == nil {
						//找到登录权限
						//清除一些重要数据，避免传递的时候泄露
						user.Pass = ""
						// 读取用户info
						var user_info models.User_info
						user_info.UserID = cookie.UserID
						if models.DB.Where(&user_info).First(&user_info).Error == nil {
							// 有数据
							//fmt.Println(user_info)

						} else {
							// 无数据
							//创建一个默认info
							user_info.AvatarPath = models.User_configs["def_avatar_path"].(string)
							user_info.UserID = cookie.UserID
							models.DB.Create(&user_info) // 传入指针
						}
						//写入当前登录的用户信息 传递给下一个组件
						ctx.Set("user_info", &user_info)
						ctx.Set("user", &user)
					} else {
						//找不到登录权限？？ 可能被封号？
						//删除前端cookie
						ctx.SetCookie("user", "", -1, "/", models.Wed_configs.Host, models.Wed_configs.Tls, true)
						cookie.Value = ""
						cookie.ExpiresAt = time.Now()
						ctx.Set("cookie", cookie)
					}

				} else {
					// ExpiresAt 在当前时间之前或等于（已过期）
					//fmt.Println("Cookie 已过期")
					//删除数据库的cookie
					models.DB.Delete(&cookie)
					//删除前端cookie
					ctx.SetCookie("user", "", -1, "/", models.Wed_configs.Host, models.Wed_configs.Tls, true)
					cookie.Value = ""
					cookie.ExpiresAt = time.Now()
					ctx.Set("cookie", cookie)
				}
			} else {
				//找不到cookie，未登录
				//删除前端cookie
				ctx.SetCookie("user", "", -1, "/", models.Wed_configs.Host, models.Wed_configs.Tls, true)
				cookie.Value = ""
				cookie.ExpiresAt = time.Now()
				ctx.Set("cookie", cookie)
			}

		}

	})

	Api_router(r.Group("/api/")) //分组路由传递到api_routers。go

	//无需权限的页面
	r.GET("/", func(ctx *gin.Context) {
		user_info, is_login := ctx.Get("user_info")
		ctx.HTML(http.StatusOK, "index.html", gin.H{
			"is_login":  is_login,
			"user_info": user_info,
		})
	})

	r.GET("/sign-up", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "sign-up.html", gin.H{})
	})
	r.GET("/sign-in", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "sign-in-01.html", gin.H{})
	})
	r.GET("/test", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "empty.html", gin.H{})
	})

	//需要权限的页面
	// r.Use(func(ctx *gin.Context) {
	// 	is_login, _ := ctx.Get("is_login")
	// 	//判断是否登录
	// 	if is_login == true {

	// 	} else {
	// 		ctx.Redirect(302, "/sign-in")
	// 	}

	// })
	r.GET("/workorders", func(ctx *gin.Context) {
		is_login, _ := ctx.Get("is_login")
		user_info, _ := ctx.Get("user_info")
		//判断是否登录
		if is_login == true {

			ctx.HTML(http.StatusOK, "workorders.html", gin.H{
				"is_login":  is_login,
				"user_info": user_info,
			})
		} else {
			ctx.Redirect(302, "/sign-in")
		}
	})

	r.GET("/warehouses", func(ctx *gin.Context) {
		is_login, _ := ctx.Get("is_login")
		user_info, _ := ctx.Get("user_info")
		//判断是否登录
		if is_login == true {
			total_pages := models.Warehouse_get_total_pages() //获取总页数
			disabled_next_page := false
			if total_pages <= 1 {
				disabled_next_page = true
			} else {
				disabled_next_page = false
			}

			ctx.HTML(http.StatusOK, "warehouses.html", gin.H{
				"total_pages":        total_pages,
				"this_page":          1,
				"prev_page":          fmt.Sprintf("/warehouses/%d", 1),
				"next_page":          fmt.Sprintf("/warehouses/%d", 2),
				"disabled_prev_page": true,
				"disabled_next_page": disabled_next_page,
				"page_range":         models.Page_range(1, total_pages, 1, "/warehouses/"),
				"Warehouses":         models.Warehouse_get_warehouses(1),
				"is_login":           is_login,
				"user_info":          user_info,
			})
		} else {
			ctx.Redirect(302, "/sign-in")
		}
	})
	r.GET("/warehouses/:id", func(ctx *gin.Context) {
		is_login, _ := ctx.Get("is_login")
		user_info, _ := ctx.Get("user_info")
		//判断是否登录
		if is_login == true {
			id := ctx.Param("id")
			id_int64, err := strconv.ParseInt(id, 10, 64)
			if err == nil {
				if id_int64 > 0 {
					total_pages := models.Warehouse_get_total_pages() //获取总页数
					disabled_next_page := false
					disabled_prev_page := false
					if total_pages == id_int64 {
						disabled_next_page = true
					} else {
						disabled_next_page = false
					}
					if id_int64 > 1 {
						disabled_prev_page = false
					} else {
						disabled_prev_page = true
					}

					//制作页码标签

					ctx.HTML(http.StatusOK, "warehouses.html", gin.H{
						"total_pages":        total_pages,
						"this_page":          id_int64,
						"prev_page":          fmt.Sprintf("/warehouses/%d", id_int64-1),
						"next_page":          fmt.Sprintf("/warehouses/%d", id_int64+1),
						"disabled_prev_page": disabled_prev_page,
						"disabled_next_page": disabled_next_page,
						"page_range":         models.Page_range(1, total_pages, id_int64, "/warehouses/"),
						"Warehouses":         models.Warehouse_get_warehouses(id_int64),
						"is_login":           is_login,
						"user_info":          user_info,
					})
				} else {
					ctx.Redirect(302, "/warehouses")
				}

			} else {
				ctx.Redirect(302, "/warehouses")
			}

		} else {
			ctx.Redirect(302, "/sign-in")
		}
	})

	r.GET("/warehouse/:id", func(ctx *gin.Context) {
		is_login, _ := ctx.Get("is_login")
		user_info, _ := ctx.Get("user_info")
		//判断是否登录
		if is_login == true {
			id := ctx.Param("id")
			id_int64, err := strconv.ParseInt(id, 10, 64)
			if err == nil {
				if id_int64 > 0 {

					ctx.HTML(http.StatusOK, "warehouses_show_item.html", gin.H{
						"warehouse_id": id_int64,
						"is_login":     is_login,
						"user_info":    user_info,
						"items":        models.Warehouse_get_items_from_whid(uint(id_int64)),
					})
				} else {
					ctx.Redirect(302, "/warehouses")
				}

			} else {
				ctx.Redirect(302, "/warehouses")
			}

		} else {
			ctx.Redirect(302, "/sign-in")
		}
	})
	r.GET("/setting-my", func(ctx *gin.Context) {

		user_info, is_login := ctx.Get("user_info")

		//判断是否登录
		if is_login == true {

			ctx.HTML(http.StatusOK, "setting-my.html", gin.H{
				"is_login":  is_login,
				"user_info": user_info,
			})
		} else {
			ctx.HTML(404, "error_404.html", gin.H{})
		}

	})

	r.GET("/setting-security", func(ctx *gin.Context) {
		user_info, is_login := ctx.Get("user_info")
		user, _ := ctx.Get("user")
		//判断是否登录
		if is_login == true {

			ctx.HTML(http.StatusOK, "setting-security.html", gin.H{
				"is_login":  is_login,
				"user_info": user_info,
				"user":      user,
			})
		} else {
			ctx.HTML(404, "error_404.html", gin.H{})
		}

	})
}
