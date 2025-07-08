package routers

import (
	"bytes"
	"fmt"
	"io"
	"time"

	"github.com/gin-gonic/gin"
)

// 中间件：打印原始请求数据
func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1. 记录基础信息
		start := time.Now()
		reqMethod := c.Request.Method
		reqURL := c.Request.URL.String()

		// 2. 读取并复制请求体
		bodyBytes, _ := io.ReadAll(c.Request.Body)
		defer c.Request.Body.Close()

		// 重要：将 body 内容写回，供后续使用
		c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		bodyString := string(bodyBytes)
		if bodyString == "" {
			bodyString = "[空请求体]"
		}

		// 3. 打印请求信息
		fmt.Printf("\n[请求开始] %s\n", start.Format(time.RFC3339))
		fmt.Printf("[方法] %s\n[URL] %s\n", reqMethod, reqURL)

		// 打印请求头（过滤敏感信息）
		fmt.Println("[请求头]:")
		for name, values := range c.Request.Header {
			// 简化打印，实际使用可添加敏感头过滤
			fmt.Printf("  %s: %s\n", name, values)
		}

		fmt.Printf("[请求体]:\n%s\n", bodyString)
		fmt.Printf("[请求结束] 耗时: %v\n\n", time.Since(start))

		c.Next()
	}
}

func V1_file_api(r *gin.RouterGroup) {
	r.GET("/", func(ctx *gin.Context) {
		Return_json(ctx, "json_error", nil)
	})

	//文件api是一定要登录的，直接用中间件判断登录状态
	// r.Use(func(ctx *gin.Context) {
	// 	Use_is_login(ctx)
	// })

	r.Use(func(ctx *gin.Context) {
		fmt.Println("file_api")

	})

	//r.Use(RequestLogger())

	r.POST("/upload", func(ctx *gin.Context) {
		//先判断有没有登录
		//获取中间件处理的结果
		_, is_login := ctx.Get("user_info")
		if is_login {
			cookie := ctx.PostForm("cookie")
			fmt.Println(cookie)

			//fmt.Println(cookie)
			file, err := ctx.FormFile("file")
			if err == nil {
				fmt.Println(file)
			} else {
				fmt.Println("err:", err)
				fmt.Println("file:", file)
			}
			Return_json(ctx, "api_ok", nil)
		} else {
			Return_json(ctx, "user_no_sign", nil)
		}

	})

	//接收头像的接口，
	r.POST("/avatar", func(ctx *gin.Context) {

	})
}
