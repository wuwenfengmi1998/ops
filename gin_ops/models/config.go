package models

import (
	"fmt"
	"os"

	"github.com/mitchellh/mapstructure"
)

var Configs map[string]interface{}

var Configs_wed Configs_web_t
var Configs_user Configs_user_t
var Configs_file Configs_file_t

var Database_configs map[string]interface{}



func init() {

}

func All_config_init() {

	//初始化数据库
	Init_database()

	//读取web配置
	err := mapstructure.Decode(Configs["web"].(map[string]interface{}), &Configs_wed)
	if err != nil {
		panic(err)
	}

	//初始化user config
	err = mapstructure.Decode(Configs["user"].(map[string]interface{}), &Configs_user)
	if err != nil {
		panic(err)
	}

	//初始化file config
	err = mapstructure.Decode(Configs["file"].(map[string]interface{}), &Configs_file)
	if err != nil {
		panic(err)
	}

	//fmt.Println(Configs_file)
	//fmt.Println(Allowed_avatar_mime)

	//创建file的关键文件夹
	for _, value := range Configs_file.Pahts {
		err := os.MkdirAll(value, 0755)
		if err != nil {
			fmt.Printf("创建文件夹失败: %v\n", err)
			panic("创建文件夹失败")

		}
	}

}
