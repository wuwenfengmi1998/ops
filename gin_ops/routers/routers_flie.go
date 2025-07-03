package routers

//文件路由

import (
	"github.com/gin-gonic/gin"
)

func Router_file(r *gin.RouterGroup) {

	//无需权限，可以直接下载的接口
	r.GET("/download/:id", func(ctx *gin.Context) {

	})

	//先在中间件判断有没有登录
	r.Use(func(ctx *gin.Context) {
		cookie_value := ctx.PostForm("cookie")
		//fmt.Println(cookie_value)

		ctx.Set("cookie_value", cookie_value)
		Use_login_from_cookie(ctx)

		//先判断有没有登录
		_, is_login := ctx.Get("user_info")
		if is_login {

		} else {
			Return_json(ctx, "user_no_sign", nil)
		}
	})

	upload := r.Group("/upload") //定义上传组
	//4大媒体上传接口，严格判断文件类型，可以直接被前端引用
	upload.POST("/image", func(ctx *gin.Context) {

	})
	upload.POST("/video", func(ctx *gin.Context) {

	})
	upload.POST("/music", func(ctx *gin.Context) {

	})
	upload.POST("/pdf", func(ctx *gin.Context) {

	})
	//其他文件，只能通过用户报告的类型定义，不能直接被前端引用
	upload.POST("/other", func(ctx *gin.Context) {

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
