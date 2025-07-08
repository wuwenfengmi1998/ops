package models

//mime信息转换位拓展名
var Mime_to_extname = map[string]string{

	"image/jpeg": ".jpeg",
	"image/png":  ".png",
	"image/gif":  ".gif",
	"image/bmp":  ".bmp",

	"video/mp4":       ".mp4",
	"video/x-msvideo": ".",
	"video/quicktime": ".",
	"video/x-flv":     ".flv",
	"video/mpeg":      ".mpeg",

	"audio/mpeg": ".mp3",
	"audio/aac":  ".acc",
	"audio/wav":  ".wav",
	"audio/flac": ".flac",

	"application/pdf": ".pdf",
}

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
	Allow_image_mime map[string]bool   `mapstructure:"allow_image_mime"`
	Allow_video_mime map[string]bool   `mapstructure:"allow_video_mime"`
	Allow_music_mime map[string]bool   `mapstructure:"allow_music_mime"`
	Allow_pdf_mime   map[string]bool   `mapstructure:"allow_pdf_mime"`
}
