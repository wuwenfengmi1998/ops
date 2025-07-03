package routers

import (
	"saas/models"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
)

func V1_user_api(r *gin.RouterGroup) {

	var err_code = Error_code["api_ok"]
	var err_msg string

	r.POST("/add", func(ctx *gin.Context) {
		//返回前端的数据

		//转换传进来的数据
		//转换传进来的数据
		var jsonData Add_user_from
		data, _ := ctx.Get("data")
		if err := mapstructure.Decode(data, &jsonData); err == nil {
			//转换字段
			newUser := models.User{
				Name:  jsonData.Username,
				Email: jsonData.Useremail,
				Pass:  jsonData.Userpass, // 实际应替换为哈希值
				Date:  time.Now(),
				// Date 字段无需赋值，数据库会自动填充默认值
			}
			//fmt.Println(newUser)
			//对用户的密码进行哈希替换
			newUser.Pass = models.Hash_user_pass(newUser.Pass)

			//用户名是唯一的，先读取是否有这个用户名
			var user models.User
			user.Name = newUser.Name

			if models.DB.Where(&user).First(&user).Error == nil {
				//fmt.Println("找到用户:", user.ID)
				Return_json(ctx, "user_name_dup", nil)
			} else {
				//fmt.Println("用户不存在")
				models.DB.Create(&newUser) // 传入指针

				//创建info
				var user_info models.User_info
				user_info.AvatarPath = models.Configs_user.Avatar_path
				user_info.UserID = newUser.ID
				models.DB.Create(&user_info) // 传入指针

				Return_json(ctx, "api_ok", nil)
			}

		} else {
			Return_json(ctx, "json_error", nil)
		}
		//Return_json(ctx, "api_err", nil)

	})

	r.POST("/login", func(ctx *gin.Context) {
		//返回前端的数据

		//转换传进来的数据
		var jsonData Login_from
		data, _ := ctx.Get("data")
		if err := mapstructure.Decode(data, &jsonData); err == nil {
			//转换字段
			newUser := models.User{
				Name: jsonData.Username,
				Pass: jsonData.Password, // 实际应替换为哈希值
				// Date 字段无需赋值，数据库会自动填充默认值
			}

			//对用户的密码进行哈希替换
			newUser.Pass = models.Hash_user_pass(newUser.Pass)

			var user models.User
			user.Name = newUser.Name
			if models.DB.Where(&user).First(&user).Error == nil {
				// 有数据
				//fmt.Println(user)
				//fmt.Println(newUser)

				if user.Pass == newUser.Pass {
					//成功登录
					//发送cookie
					//cookie时间
					var cookie_time = 0
					if jsonData.Is_keep_login {
						cookie_time = models.Configs_user.Cookie_timeout
					}

					cookie := models.Rand_str_32() //生成32字节cookie
					//cookie := "testcookie"
					//fmt.Println(cookie)
					//将cookie写进数据库
					new_cookie := models.Cookie{}
					new_cookie.Domain = models.Configs_wed.Host
					new_cookie.Name = "user"
					new_cookie.Value = cookie
					new_cookie.UserID = user.ID

					//cookie时间
					new_cookie.CreatedAt = time.Now()
					new_cookie.UpdatedAt = new_cookie.CreatedAt
					//计算cookie失效时间
					new_cookie.ExpiresAt = time.Now().Add(time.Duration(models.Configs_user.Cookie_timeout) * time.Second) //计算过期时间
					new_cookie.SecureFlag = models.Configs_wed.Tls

					//发送到前端
					ctx.SetCookie("user", cookie, cookie_time, "/", models.Configs_wed.Host, models.Configs_wed.Tls, true)
					ctx.Set("cookie", new_cookie)
					//写到数据库
					models.DB.Create(&new_cookie) // 传入指针

					//获取用户info
					user_info := models.User_info{
						UserID: user.ID,
					}

					models.DB.Where(&user_info).First(&user_info)

					red := map[string]interface{}{
						"cookie":    new_cookie,
						"user_info": user_info,
					}

					Return_json(ctx, "api_ok", red)

				} else {
					//密码错误

					Return_json(ctx, "user_password_err", nil)
				}

			} else {
				//fmt.Println("用户不存在")

				Return_json(ctx, "user_name_nofind", nil)
			}

		} else {
			Return_json(ctx, "json_error", nil)
		}

		//Return_json(ctx, "api_err", nil)

	})

	r.POST("/logout", func(ctx *gin.Context) {
		//返回前端的数据

		//先判断是否已经登录
		//获取中间件处理的结果
		_, is_login := ctx.Get("user_info")
		//fmt.Println(is_login)
		//fmt.Println(user_info)
		if is_login == true {
			//fmt.Println("loged")
			cookie_any, _ := ctx.Get("cookie") //这个cookie在中间件已经判断为有效的，否则is_login不可能为true，所以直接在数据库删除应该是安全的
			//删除数据库里的cookie
			var cookie models.Cookie
			if err := mapstructure.Decode(cookie_any, &cookie); err == nil {
				models.DB.Where(&cookie).Delete(&cookie)
				//删除前端cookie
				ctx.SetCookie("user", "", -1, "/", models.Configs_wed.Host, models.Configs_wed.Tls, true)
				ctx.Set("cookie", nil)

				Return_json(ctx, "api_ok", nil)
			} else {

				Return_json(ctx, "json_error", nil)

			}

		} else {
			ctx.SetCookie("user", "", -1, "/", models.Configs_wed.Host, models.Configs_wed.Tls, true)
			Return_json(ctx, "user_no_sign", nil)
		}

	})

	r.POST("/updata_info", func(ctx *gin.Context) {
		//返回前端的数据
		err_msg = "user_api_error"
		err_code = Error_code[err_msg]

		//先判断是否已经登录
		//获取中间件处理的结果
		is_login, _ := ctx.Get("is_login")
		if is_login == true {
			//转换传进来的数据
			var jsonData map[string]interface{}
			if err := ctx.ShouldBindJSON(&jsonData); err == nil {

				user_info_, _ := ctx.Get("user_info")
				user_info, _ := user_info_.(*models.User_info) //这个数据本身就是从数据库读出来的，理论上结构转换不会出错
				user_info_find := models.User_info{
					ID: user_info.ID,
				}

				new_user_info := models.User_info{
					AvatarPath: jsonData["avatar"].(string),
					FirstName:  jsonData["first_name"].(string),
					Username:   jsonData["username"].(string),
					Birthdate:  models.Time_date_str_to_time(jsonData["birthday"].(string)),
				}

				//需要验证传入数据的合法性 例如头像url是否站内的
				if strings.HasPrefix(new_user_info.AvatarPath, models.Configs_user.Avatar_ginrouter_path) {

				} else {
					new_user_info.AvatarPath = models.Configs_user.Avatar_path
				}

				//fmt.Printf("%%#v: %#v\n", new_user_info)
				models.DB.Where(&user_info_find).Updates(&new_user_info)

				err_msg = "api_ok"
				err_code = Error_code[err_msg]

			} else {
				err_msg = "json_error"
				err_code = Error_code[err_msg]
			}

		} else {
			//fmt.Println("no loged")
			err_msg = "user_no_sign"
			err_code = Error_code[err_msg]
		}

		ctx.JSON(200, map[string]interface{}{
			"api":      "ok",
			"err_code": err_code,
			"err_msg":  err_msg,
		})
	})

	r.POST("/change_email", func(ctx *gin.Context) {
		//返回前端的数据
		err_msg = "user_api_error"
		err_code = Error_code[err_msg]

		//先判断是否已经登录
		//获取中间件处理的结果
		is_login, _ := ctx.Get("is_login")
		if is_login == true {
			//转换传进来的数据
			var jsonData map[string]interface{}
			if err := ctx.ShouldBindJSON(&jsonData); err == nil {

				//需要验证传入数据的合法性
				if models.Is_email_valid(jsonData["new_email"].(string)) {
					user_, _ := ctx.Get("user")
					user, _ := user_.(*models.User)
					user_find := models.User{
						ID: user.ID,
					}
					user_new := models.User{
						Email: jsonData["new_email"].(string),
					}
					models.DB.Where(&user_find).Updates(&user_new)
					err_msg = "api_ok"
					err_code = Error_code[err_msg]

				} else {
					err_msg = "email_error"
					err_code = Error_code[err_msg]
				}

			} else {
				err_msg = "json_error"
				err_code = Error_code[err_msg]
			}

		} else {
			//fmt.Println("no loged")
			err_msg = "user_no_sign"
			err_code = Error_code[err_msg]
		}

		ctx.JSON(200, map[string]interface{}{
			"api":      "ok",
			"err_code": err_code,
			"err_msg":  err_msg,
		})
	})

	r.POST("/change_pass", func(ctx *gin.Context) {
		//返回前端的数据
		err_msg = "user_api_error"
		err_code = Error_code[err_msg]

		//先判断是否已经登录
		//获取中间件处理的结果
		is_login, _ := ctx.Get("is_login")
		if is_login == true {
			//转换传进来的数据
			var jsonData map[string]interface{}
			if err := ctx.ShouldBindJSON(&jsonData); err == nil {

				//需要验证传入数据的合法性

				//读取已登录的用户信息
				user_, _ := ctx.Get("user")
				user, _ := user_.(*models.User)
				user_find := models.User{
					ID: user.ID,
				}

				models.DB.Where(&user_find).First(&user_find)

				pass_old := jsonData["pass_old"].(string)
				pass_new := jsonData["pass_new"].(string)
				//对用户的密码进行哈希替换
				pass_old = models.Hash_user_pass(pass_old)
				pass_new = models.Hash_user_pass(pass_new)

				if user_find.Pass == pass_old {

					new_user := models.User{
						Pass: pass_new,
					}

					//修改密码
					models.DB.Where(&user_find).Updates(&new_user)

					//密码修改后所有cookie都应该失效
					cookie_find := models.Cookie{
						UserID: user.ID,
					}
					models.DB.Where(&cookie_find).Delete(&cookie_find)

					err_msg = "api_ok"
					err_code = Error_code[err_msg]

				} else {
					err_msg = "user_password_err"
					err_code = Error_code[err_msg]
				}

			} else {
				err_msg = "json_error"
				err_code = Error_code[err_msg]
			}

		} else {
			//fmt.Println("no loged")
			err_msg = "user_no_sign"
			err_code = Error_code[err_msg]
		}

		ctx.JSON(200, map[string]interface{}{
			"api":      "ok",
			"err_code": err_code,
			"err_msg":  err_msg,
		})
	})
}
