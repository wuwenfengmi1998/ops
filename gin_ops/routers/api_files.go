package routers

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"path"
	"path/filepath"
	"saas/models"

	"github.com/gin-gonic/gin"
)

func V1_file_api(r *gin.RouterGroup) {
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, map[string]interface{}{
			"error": "you need use Post",
		})
	})
	r.POST("/upload", func(ctx *gin.Context) {
		file, err := ctx.FormFile("file")
		if err == nil {
			dst := path.Join("./data/upload", file.Filename)
			ctx.SaveUploadedFile(file, dst)
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
