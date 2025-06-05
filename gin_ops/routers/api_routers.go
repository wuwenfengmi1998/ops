package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Api_router(r *gin.RouterGroup) {

	r.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "api")
	})

	v1_api := r.Group("/v1/")
	{
		V1_user_api(v1_api.Group("/user/"))
		V1_file_api(v1_api.Group("/file/"))
		V1_cookie_api(v1_api.Group("/cookie/"))
		V1_warehouses_api(v1_api.Group("/warehouses_api/"))
	}

}
