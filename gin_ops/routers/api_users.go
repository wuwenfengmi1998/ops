package routers

import (
	"fmt"
	"path"
	"saas/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
)

func user_logout(ctx *gin.Context) {
	//返回前端的数据

	//先判断是否已经登录
	//获取中间件处理的结果
	_, is_login := ctx.Get("user_info")
	//fmt.Println(is_login)
	//fmt.Println(user_info)
	if is_login {
		//fmt.Println("loged")
		cookie_any, _ := ctx.Get("cookie") //这个cookie在中间件已经判断为有效的，否则is_login不可能为true，所以直接在数据库删除应该是安全的
		//删除数据库里的cookie
		var cookie models.Cookie
		if err := mapstructure.Decode(cookie_any, &cookie); err == nil {
			models.DB.Where("ID=?", cookie.ID).Delete(&cookie)
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
}

func V1_user_api(r *gin.RouterGroup) {

	r.POST("/add", func(ctx *gin.Context) {
		//返回前端的数据

		//转换传进来的数据
		//转换传进来的数据
		var jsonData From_user_add
		data, is_have_data := ctx.Get("data")
		//fmt.Println(data)
		if is_have_data {
			if err := mapstructure.Decode(data, &jsonData); err == nil {
				//转换字段
				newUser := models.User{
					Name:  jsonData.Username,
					Email: jsonData.Useremail,
					Pass:  jsonData.Userpass, // 实际应替换为哈希值
					Date:  time.Now(),
					// Date 字段无需赋值，数据库会自动填充默认值
				}
				if newUser.Name != "" && newUser.Pass != "" && newUser.Email != "" {
					//fmt.Println(newUser)
					//对用户的密码进行哈希替换
					newUser.Pass = models.Hash_user_pass(newUser.Pass)

					//用户名是唯一的，先读取是否有这个用户名
					var user models.User
					user.Name = newUser.Name

					if models.DB.Where("Name=?", user.Name).First(&user).Error == nil {
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

			} else {
				Return_json(ctx, "json_error", nil)
			}
		} else {
			Return_json(ctx, "json_error", nil)
		}

	})

	r.POST("/login", func(ctx *gin.Context) {
		//返回前端的数据

		//转换传进来的数据
		var jsonData From_user_loggin
		data, is_have_data := ctx.Get("data")
		if is_have_data {
			if err := mapstructure.Decode(data, &jsonData); err == nil {
				if jsonData.Username != "" {
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
					if models.DB.Where("Name=?", user.Name).First(&user).Error == nil {
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

							models.DB.Where("ID=?", user_info.ID).First(&user_info)

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
					Return_json(ctx, "user_name_nofind", nil)
				}

			} else {
				Return_json(ctx, "json_error", nil)
			}
		} else {
			Return_json(ctx, "json_error", nil)
		}

	})

	r.POST("/logout", func(ctx *gin.Context) {
		user_logout(ctx)
	})

	r.POST("/updata_info", func(ctx *gin.Context) {
		user_info_any, is_login := ctx.Get("user_info")
		if is_login {

			//转换传进来的数据
			var json_data From_user_updata_info
			data, is_have_data := ctx.Get("data")
			if is_have_data {
				if err := mapstructure.Decode(data, &json_data); err == nil {
					//fmt.Println(json_data)

					updata_user_info := models.User_info{
						FirstName: json_data.First_name,
						Username:  json_data.Username,
						Birthdate: models.Time_date_str_to_time(json_data.Birthday),
					}

					sele := []string{"FirstName", "Username", "Birthdate"}

					user_info := user_info_any.(*models.User_info)
					//先判断头像是否合法
					if json_data.Avatar_id != 0 {
						file_info := models.File_info{}
						file_info.ID = json_data.Avatar_id
						if models.DB.Where("ID=?", file_info.ID).First(&file_info).Error == nil {
							//读取到文件，判断是不是图片
							if file_info.Type == "image" && file_info.UserID == user_info.UserID {
								file_id_str := fmt.Sprintf("%d", file_info.ID)
								url_preview := path.Join(Url_flie_preview_from_id_head, file_id_str)
								updata_user_info.AvatarPath = url_preview
								sele = append(sele, "AvatarPath")
							}

						}
					}

					if models.DB.Where("ID=?", user_info.ID).Select(sele).Updates(&updata_user_info).Error == nil {
						Return_json(ctx, "api_ok", nil)
					} else {
						Return_json(ctx, "DB_err", nil)
					}

				} else {
					Return_json(ctx, "json_error", nil)
				}
			} else {
				Return_json(ctx, "json_error", nil)
			}

			//var user_info models.User_info{}{}

		} else {
			Return_json(ctx, "user_no_sign", nil)
		}

	})

	r.POST("/change_email", func(ctx *gin.Context) {
		user_info_any, is_login := ctx.Get("user_info")
		if is_login {
			data, is_have_data := ctx.Get("data")
			if is_have_data {
				datas := data.(*map[string]interface{})
				email_str := (*datas)["new_email"].(string)
				if models.Is_email_valid(email_str) {
					user_updata := models.User{
						Email: email_str,
					}
					user_info := user_info_any.(*models.User_info)
					user_fund := models.User{
						ID: user_info.UserID,
					}
					if models.DB.Where("ID=?", user_fund.ID).Updates(&user_updata).Error == nil {
						Return_json(ctx, "api_ok", nil)
					} else {
						Return_json(ctx, "DB_err", nil)
					}

				} else {
					Return_json(ctx, "email_error", nil)
				}

			} else {
				Return_json(ctx, "json_error", nil)
			}

		} else {
			Return_json(ctx, "user_no_sign", nil)
		}

	})

	r.POST("/change_pass", func(ctx *gin.Context) {
		user_info_any, is_login := ctx.Get("user_info")
		if is_login {
			data, is_have_data := ctx.Get("data")
			if is_have_data {
				datas := data.(*map[string]interface{})
				pass_old := (*datas)["pass_old"].(string)
				pass_new := (*datas)["pass_new"].(string)
				pass_old = models.Hash_user_pass(pass_old)
				pass_new = models.Hash_user_pass(pass_new)
				user_info := user_info_any.(*models.User_info)
				user_fund := models.User{
					ID: user_info.UserID,
				}
				models.DB.Where("ID=?", user_fund.ID).First(&user_fund)
				if user_fund.Pass == pass_old {
					user_new := models.User{
						Pass: pass_new,
					}
					if models.DB.Where("ID=?", user_fund.ID).Updates(&user_new).Error == nil {
						user_logout(ctx)
						//Return_json(ctx, "api_ok", nil)
					} else {
						Return_json(ctx, "DB_err", nil)
					}
				} else {
					Return_json(ctx, "user_password_err", nil)
				}

			} else {
				Return_json(ctx, "json_error", nil)
			}

		} else {
			Return_json(ctx, "user_no_sign", nil)
		}

	})
}
