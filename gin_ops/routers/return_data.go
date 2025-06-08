package routers

import "github.com/gin-gonic/gin"

func Return_json(ctx *gin.Context, err_msg string, data map[string]interface{}) {
	var err_code = Error_code[err_msg]
	ctx.JSON(200, map[string]interface{}{
		"err_code": err_code,
		"err_msg":  err_msg,
		"return":   data,
	})

}
