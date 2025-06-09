package routers

import "github.com/gin-gonic/gin"

func Return_json(ctx *gin.Context, err_msg string, data map[string]interface{}) {
	var err_code = Error_code[err_msg]
	return_data := map[string]interface{}{}

	cookie, have_cookie := ctx.Get("cookie")
	if have_cookie {
		return_data["cookie"] = cookie
	}

	return_data["err_code"] = err_code
	return_data["err_msg"] = err_msg
	if data != nil {
		return_data["return"] = data
	}

	ctx.JSON(200, &return_data)

}
