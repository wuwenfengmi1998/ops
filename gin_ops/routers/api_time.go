package routers

import "github.com/gin-gonic/gin"

func V1_time_api(r *gin.RouterGroup) {
	r.GET("/now", func(ctx *gin.Context) {
		
	})

}
