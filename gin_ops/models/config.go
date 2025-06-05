package models

var Configs map[string]interface{}

var Wed_configs map[string]interface{}

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
	Wed_configs = Configs["web"].(map[string]interface{})

	//初始化数据库
	Database_init()

	//初始化user config
	User_configs = Configs["user"].(map[string]interface{})

}
