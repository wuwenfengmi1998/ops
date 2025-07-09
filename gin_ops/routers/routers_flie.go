package routers

//文件路由

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

func Router_file(r *gin.RouterGroup) {

	//无需权限，可以直接下载的接口
	r.GET("/download/:id", func(ctx *gin.Context) {

	})

	//先在中间件判断有没有传进cookie
	r.Use(func(ctx *gin.Context) {
		cookie_value := ctx.PostForm("cookie")
		//fmt.Println(cookie_value)

		if cookie_value != "" {
			ctx.Set("cookie_value", cookie_value)
		}

		Use_login_from_cookie(ctx) //因为需要同步前端cookie，所有只能用ctx

	})

	upload := r.Group("/upload") //定义上传组
	//4大媒体上传接口，严格判断文件类型，可以直接被前端引用
	upload.POST("/image", func(ctx *gin.Context) {

		//先判断有没有登录
		user_info_any, is_login := ctx.Get("user_info") //因为需要读取user_info，避免重复调用 这一步就不在中间件操作了
		if is_login {
			user_info := user_info_any.(*models.User_info) //直接断言类型，这个值是从数据库直接读取的 理论不会出错
			file, err := ctx.FormFile("file")
			if err == nil {

				//限制文件大小
				if file.Size > 512 {
					if file.Size < int64(models.Configs_file.Max_size) {
						// 2. 安全获取文件名并处理路径问题
						filename := filepath.Base(file.Filename) // 防御性处理路径分隔符
						//fmt.Println(filename)
						// 3. 获取标准后缀名（含点）
						//extWithDot := filepath.Ext(filename)

						//判断文件mime是否合法
						// 打开文件流
						src_mime, _ := file.Open()
						defer src_mime.Close()
						// 读取前512字节用于MIME检测
						buffer := make([]byte, 512)
						io.ReadFull(src_mime, buffer)
						// 检测MIME类型
						mimeType := http.DetectContentType(buffer)
						file_extname := models.Configs_file.Allow_image_mime[mimeType]
						if file_extname != "" {
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

							new_filename := fmt.Sprintf("%s%s", hashString, file_extname)
							file.Filename = new_filename

							//fmt.Println(user_info)

							//这是上传的真实路径
							dst := path.Join(models.Configs_file.Pahts["image"], file.Filename)

							//判断文件是否存在避免重复保存
							if models.File_exists(dst) {
								//fmt.Println("文件存在")

							} else {
								//fmt.Println("文件no存在")
								ferr := ctx.SaveUploadedFile(file, dst)
								if ferr == nil {
									//文件保存成功

								} else {

									Return_json(ctx, "file_save_err", nil)
									ctx.Abort() //end
								}
							}
							//记录到数据库

							//先检查数据库有没有数据
							fund_file_info := models.File_info{
								Name:   filename,
								Sha256: hashString,
								Mime:   mimeType,
								Type:   "image",
								UserID: user_info.UserID,
							}
							fund_file_info2 := models.File_info{}

							models.DB.Where(&fund_file_info).Find(&fund_file_info2)

							if fund_file_info2.ID != 0 {
								fmt.Println(fund_file_info2)
								fund_file_info2.Const += 1
								models.DB.Where(&fund_file_info).Updates(&fund_file_info2)
							} else {
								models.DB.Create(&fund_file_info) // 传入指针
								fund_file_info2 = fund_file_info
							}

							red := map[string]interface{}{
								"data": fund_file_info2,
							}

							Return_json(ctx, "api_ok", red)

						} else {

							Return_json(ctx, "file_mime_err", nil)
						}

					} else {

						Return_json(ctx, "file_size_err", nil)
					}
				} else {

					Return_json(ctx, "file_size_err", nil)
				}
			} else {
				Return_json(ctx, "file_get_err", nil)
			}

		} else {
			Return_json(ctx, "user_no_sign", nil)
		}

	})
	upload.POST("/video", func(ctx *gin.Context) {
		Return_json(ctx, "api_ok", nil)
	})
	upload.POST("/music", func(ctx *gin.Context) {
		Return_json(ctx, "api_ok", nil)
	})
	upload.POST("/pdf", func(ctx *gin.Context) {
		Return_json(ctx, "api_ok", nil)
	})
	//其他文件，只能通过用户报告的类型定义，不能直接被前端引用
	upload.POST("/other", func(ctx *gin.Context) {
		Return_json(ctx, "api_ok", nil)
	})

	// r.POST("/upload", func(ctx *gin.Context) {
	// 	file, err := ctx.FormFile("file")
	// 	if err == nil {
	// 		fmt.Println("ok")
	// 	} else {
	// 		fmt.Println("err:", err)
	// 		fmt.Println("file:", file)
	// 	}
	// 	Return_json(ctx, "api_ok", nil)
	// })

}
