package routers

import "github.com/gin-gonic/gin"

func V1_cookie_api(r *gin.RouterGroup) {
	r.GET("/test", func(ctx *gin.Context) {
		ctx.SetCookie("test", "testcookie", 100, "/", "127.0.0.1", false, true)
	})
}
