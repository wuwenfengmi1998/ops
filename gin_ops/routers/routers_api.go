package routers

import (
	"saas/models"

	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
)

func Router_api(r *gin.RouterGroup) {

	r.GET("/", func(ctx *gin.Context) {
		Return_json(ctx, "api_ok", nil)
	})

	r.Use(func(ctx *gin.Context) {
		//转换传进来的数据
		var jsonData map[string]interface{}
		if err := ctx.ShouldBindJSON(&jsonData); err == nil {
			//分离数据
			var cookie_t models.Cookie
			if err = mapstructure.Decode(jsonData["cookie"], &cookie_t); err == nil {
				if cookie_t.Value != "" {
					cookie_vel := cookie_t.Value
					ctx.Set("cookie_value", cookie_vel)
				}

			}

			var data_t map[string]interface{}
			if err = mapstructure.Decode(jsonData["data"], &data_t); err == nil {
				ctx.Set("data", &data_t)
			}

		}

		Use_login_from_cookie(ctx)
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
