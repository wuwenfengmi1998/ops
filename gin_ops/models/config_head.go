package models

//mime信息转换位拓展名

type Configs_web_t struct {
	Host              string `mapstructure:"host"`
	Port              string `mapstructure:"port"`
	Tls               bool   `mapstructure:"tls"`
	Cert_private_path string `mapstructure:"cert_private_path"`
	Cert_public_path  string `mapstructure:"cert_public_path"`
}

type Configs_user_t struct {
	Cookie_timeout        int    `mapstructure:"cookie_timeout"`
	Pass_hash_type        string `mapstructure:"pass_hash_type"`
	Avatar_save_path      string `mapstructure:"avatar_save_path"`
	Avatar_ginrouter_path string `mapstructure:"avatar_ginrouter_path"`
	Avatar_path           string `mapstructure:"avatar_path"`
}

type Configs_file_t struct {
	Max_size         uint64            `mapstructure:"max_size"`
	Pahts            map[string]string `mapstructure:"pahts"`
	Allow_image_mime map[string]string `mapstructure:"allow_image_mime"`
	Allow_video_mime map[string]string `mapstructure:"allow_video_mime"`
	Allow_music_mime map[string]string `mapstructure:"allow_music_mime"`
	Allow_pdf_mime   map[string]string `mapstructure:"allow_pdf_mime"`
}
