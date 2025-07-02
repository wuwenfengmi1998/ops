package routers

import (
	"fmt"
	"net/http"
	"saas/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Def_router(r *gin.RouterGroup) {

	r.Use(func(ctx *gin.Context) {
		cookie_vel := ""
		//读取用户cookie，判断用户是否已登录
		cookie_s, is_have_cookie := ctx.Cookie("user")
		if is_have_cookie == nil {
			cookie_vel = cookie_s
		}

		//fmt.Println(cookie_vel)
		if cookie_vel != "" {
			ctx.Set("cookie_value", cookie_vel)

		}

		Use_login_from_cookie(ctx)
	})

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
