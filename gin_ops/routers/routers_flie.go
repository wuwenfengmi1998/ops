package routers

//文件路由

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Router_file(r *gin.RouterGroup) {

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

	r.POST("/upload", func(ctx *gin.Context) {
		file, err := ctx.FormFile("file")
		if err == nil {
			fmt.Println("ok")
		} else {
			fmt.Println("err:", err)
			fmt.Println("file:", file)
		}
		Return_json(ctx, "api_ok", nil)
	})

}
