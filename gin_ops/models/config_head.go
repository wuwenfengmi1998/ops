package models

type Wed_configs_t struct {
	Host              string `mapstructure:"host"`
	Port              string `mapstructure:"port"`
	Tls               bool   `mapstructure:"tls"`
	Cert_private_path string `mapstructure:"cert_private_path"`
	Cert_public_path  string `mapstructure:"cert_public_path"`
}
