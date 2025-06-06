package models

import "github.com/mitchellh/mapstructure"

var Configs map[string]interface{}

var Wed_configs Wed_configs_t

var Database_configs map[string]interface{}

var User_configs map[string]interface{}

var Allowed_avatar_ext = map[string]bool{
	".jpg":  true,
	".jpeg": true,
	".png":  true,
}

var Allowed_avatar_mime = map[string]bool{
	"image/jpeg": true,
	"image/png":  true,
}

func init() {

}

func All_config_init() {
	//读取web配置

	err := mapstructure.Decode(Configs["web"].(map[string]interface{}), &Wed_configs)
	if err != nil {
		panic(err)
	}

	//初始化数据库
	Database_init()

	//初始化user config
	User_configs = Configs["user"].(map[string]interface{})

}
