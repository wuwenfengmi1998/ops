package routers

import (
	"github.com/gin-gonic/gin"
)

func Api_router(r *gin.RouterGroup) {

	r.GET("/", func(ctx *gin.Context) {
		Return_json(ctx, "api_ok", nil)
	})

	v1_api := r.Group("/v1/")
	{
		V1_time_api(v1_api.Group("/time/"))
		V1_user_api(v1_api.Group("/user/"))
		V1_file_api(v1_api.Group("/file/"))
		V1_cookie_api(v1_api.Group("/cookie/"))
		V1_warehouses_api(v1_api.Group("/warehouses_api/"))
	}

}
