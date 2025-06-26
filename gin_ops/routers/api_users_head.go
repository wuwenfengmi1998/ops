package routers

type Login_from struct {
	Is_keep_login bool   `json:"is_keep_login"`
	Username      string `json:"username"`
	Password      string `json:"password"`
}

type Add_user_from struct {
	Useremail string `json:"useremail"`
	Username  string `json:"username"`
	Userpass  string `json:"userpass"`
}
