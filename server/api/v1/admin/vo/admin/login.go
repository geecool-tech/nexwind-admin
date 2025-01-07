package admin

// UserLoginReq 用户登录请求参数
type UserLoginReq struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}
