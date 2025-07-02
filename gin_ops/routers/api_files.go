package routers

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"path"
	"path/filepath"
	"saas/models"
	"time"

	"github.com/gin-gonic/gin"
)

// 中间件：打印原始请求数据
func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1. 记录基础信息
		start := time.Now()
		reqMethod := c.Request.Method
		reqURL := c.Request.URL.String()

		// 2. 读取并复制请求体
		bodyBytes, _ := io.ReadAll(c.Request.Body)
		defer c.Request.Body.Close()

		// 重要：将 body 内容写回，供后续使用
		c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		bodyString := string(bodyBytes)
		if bodyString == "" {
			bodyString = "[空请求体]"
		}

		// 3. 打印请求信息
		fmt.Printf("\n[请求开始] %s\n", start.Format(time.RFC3339))
		fmt.Printf("[方法] %s\n[URL] %s\n", reqMethod, reqURL)

		// 打印请求头（过滤敏感信息）
		fmt.Println("[请求头]:")
		for name, values := range c.Request.Header {
			// 简化打印，实际使用可添加敏感头过滤
			fmt.Printf("  %s: %s\n", name, values)
		}

		fmt.Printf("[请求体]:\n%s\n", bodyString)
		fmt.Printf("[请求结束] 耗时: %v\n\n", time.Since(start))

		c.Next()
	}
}

func V1_file_api(r *gin.RouterGroup) {
	r.GET("/", func(ctx *gin.Context) {
		Return_json(ctx, "json_error", nil)
	})

	//文件api是一定要登录的，直接用中间件判断登录状态
	// r.Use(func(ctx *gin.Context) {
	// 	Use_is_login(ctx)
	// })

	r.Use(func(ctx *gin.Context) {
		fmt.Println("file_api")

	})

	//r.Use(RequestLogger())

	r.POST("/upload", func(ctx *gin.Context) {
		//先判断有没有登录
		//获取中间件处理的结果
		_, is_login := ctx.Get("user_info")
		if is_login {
			cookie := ctx.PostForm("cookie")
			fmt.Println(cookie)

			//fmt.Println(cookie)
			file, err := ctx.FormFile("file")
			if err == nil {
				fmt.Println(file)
			} else {
				fmt.Println("err:", err)
				fmt.Println("file:", file)
			}
			Return_json(ctx, "api_ok", nil)
		} else {
			Return_json(ctx, "user_no_sign", nil)
		}

	})

	//接收头像的接口，
	r.POST("/avatar", func(ctx *gin.Context) {
		//返回前端的数据
		err_msg := "user_api_error"
		err_code := Error_code[err_msg]
		data := map[string]interface{}{}

		//判断权限是否可以接收
		//先判断是否已经登录
		//获取中间件处理的结果
		is_login, _ := ctx.Get("is_login")
		if is_login == true {
			//读取用户id信息
			user, _ := ctx.Get("user")
			//保存这个头像
			file, err := ctx.FormFile("file")
			if err == nil {
				//限制文件大小
				if file.Size > 512 {
					if file.Size < 1024000 {
						// 2. 安全获取文件名并处理路径问题
						filename := filepath.Base(file.Filename) // 防御性处理路径分隔符
						// 3. 获取标准后缀名（含点）
						extWithDot := filepath.Ext(filename)
						//判断后缀名类型是否是允许的
						if models.Allowed_avatar_ext[extWithDot] {
							//判断文件mime是否合法
							// 打开文件流
							src_mime, _ := file.Open()
							defer src_mime.Close()
							// 读取前512字节用于MIME检测
							buffer := make([]byte, 512)
							io.ReadFull(src_mime, buffer)
							// 检测MIME类型
							mimeType := http.DetectContentType(buffer)
							if models.Allowed_avatar_mime[mimeType] {
								// 打开文件流
								src, _ := file.Open()
								defer src.Close()
								// 创建SHA256哈希器
								hasher := sha256.New()

								// 计算哈希值
								io.Copy(hasher, src)
								// 获取哈希结果
								hashBytes := hasher.Sum(nil)
								hashString := hex.EncodeToString(hashBytes)

								new_filename := fmt.Sprintf("%d_%s%s", user.(*models.User).ID, hashString, extWithDot)
								file.Filename = new_filename

								//这是上传的真实路径
								dst := path.Join("./data/avatar", file.Filename)

								//这是经过gin路由的路径
								gin_dat := path.Join("/avatar", file.Filename)

								//判断文件是否存在避免重复保存
								if models.File_exists(dst) {
									//fmt.Println("文件存在")
									err_msg = "api_ok"
									err_code = Error_code[err_msg]
									//返回gin路由的路径
									data["path"] = gin_dat
									data["new_path"] = false
								} else {
									//fmt.Println("文件no存在")
									ferr := ctx.SaveUploadedFile(file, dst)
									if ferr == nil {
										//文件保存成功
										err_msg = "api_ok"
										err_code = Error_code[err_msg]
										//返回gin路由的路径
										data["path"] = gin_dat
										data["new_path"] = true

									} else {
										err_msg = "file_save_err"
										err_code = Error_code[err_msg]
										fmt.Println(ferr)
										data["err"] = ferr
									}
								}
							} else {
								err_msg = "file_mime_err"
								err_code = Error_code[err_msg]
							}

						} else {
							err_msg = "file_type_err"
							err_code = Error_code[err_msg]
						}
					} else {
						err_msg = "file_size_err"
						err_code = Error_code[err_msg]
					}
				} else {
					err_msg = "file_size_err"
					err_code = Error_code[err_msg]
				}

			} else {
				err_msg = "file_get_err"
				err_code = Error_code[err_msg]
			}
		} else {
			err_msg = "user_no_sign"
			err_code = Error_code[err_msg]
		}
		ctx.JSON(200, map[string]interface{}{
			"api":      "ok",
			"err_code": err_code,
			"err_msg":  err_msg,
			"data":     data,
		})
	})
}
