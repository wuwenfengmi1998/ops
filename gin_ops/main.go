package main

import (
	"fmt"
	"os"
	"saas/models"
	"saas/routers"

	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v3"
)

func init() {

}

func main() {
	fmt.Println("main_init!")
	data_path := "./data/upload/"
	avatar_path := "./data/avatar/"
	config_file_path := "./data/config.yaml"
	config_temp_path := "./def_config/config_temp.yaml"
	// 直接尝试创建所有必要的目录
	err := os.MkdirAll(data_path, 0755)
	if err != nil {
		fmt.Printf("创建文件夹失败: %v\n", err)
		return
	}
	err = os.MkdirAll(avatar_path, 0755)
	if err != nil {
		fmt.Printf("创建文件夹失败: %v\n", err)
		return
	}
	//fmt.Println("文件夹创建成功或已存在")
	//尝试读取配置
	if !models.File_exists(config_file_path) {
		fmt.Println("读取配置失败")

		//复制配置模板
		fmt.Println("复制配置模板")
		input, err := os.ReadFile(config_temp_path)
		if err != nil {
			fmt.Println(err)
			return
		}
		err = os.WriteFile(config_file_path, input, 0644)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("需要修改此配置:%s\n", config_file_path)
		return
	}
	//读取默认配置
	data, err := os.ReadFile(config_file_path)
	if err != nil {
		fmt.Println(err)
		return
	}

	if err := yaml.Unmarshal(data, &models.Configs); err != nil {
		fmt.Println(err)
		return
	}

	if models.Configs["configed"] == false {
		fmt.Printf("需要将:%s 内的configed设置为true", config_file_path)
		return
	}

	//统一初始化
	models.All_config_init()

	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	//
	r.StaticFile("/favicon.ico", "./static/favicon.ico")
	r.StaticFile("/robots.txt", "./static/robots.txt")
	r.Static("/static/", "./static/static/")
	r.Static("/dist/", "./static/dist/")

	//静态用户上传的文件
	r.Static("/avatar/", avatar_path)
	//store := cookie.NewStore([]byte("secret"))

	// 自定义 404 页面（需要提前加载模板）
	r.NoRoute(func(ctx *gin.Context) {
		ctx.HTML(404, "error_404.html", gin.H{})
	})

	r.Use(func(ctx *gin.Context) {
		routers.Fitst_use(ctx)
	})
	routers.Def_router(r.Group("/"))     //分组路由传递到def_routers。go
	routers.Api_router(r.Group("/api/")) //分组路由传递到api_routers。go

	var http_port = models.Wed_configs.Host + ":" + models.Wed_configs.Port
	var gin_port = "0.0.0.0" + ":" + models.Wed_configs.Port
	if models.Wed_configs.Tls {
		if models.Wed_configs.Cert_public_path == "" || models.Wed_configs.Cert_private_path == "" {
			fmt.Printf("需要配置证书路径")
			return
		} else {
			fmt.Println("https://" + http_port)
			r.RunTLS(gin_port, models.Wed_configs.Cert_public_path, models.Wed_configs.Cert_private_path)
		}
	} else {
		fmt.Println("http://" + http_port)
		r.Run(gin_port)
	}

}
