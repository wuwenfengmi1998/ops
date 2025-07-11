package routers

type From_user_loggin struct {
	Is_keep_login bool   `json:"is_keep_login"`
	Username      string `json:"username"`
	Password      string `json:"password"`
}

type From_user_add struct {
	Useremail string `json:"useremail"`
	Username  string `json:"username"`
	Userpass  string `json:"userpass"`
}

type From_user_updata_info struct {
	Avatar_id  uint   `json:"avatar_id"`
	Username   string `json:"username"`
	First_name string `json:"first_name"`
	Birthday   string `json:"birthday"`
}
