package routers

import (
	"saas/models"
	"time"

	"github.com/gin-gonic/gin"
)

// 从cookie登录用户
func Use_login_from_cookie(ctx *gin.Context) {
	//先从缓存获取cookie值
	cookie_value, is_have_cookie := ctx.Get("cookie_value")
	if is_have_cookie {
		//fmt.Println(cookie_value)

		var cookie models.Cookie
		cookie.Value = cookie_value.(string)
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
					//fmt.Println(user_info)
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

}

func Use_is_login(ctx *gin.Context) {

}
