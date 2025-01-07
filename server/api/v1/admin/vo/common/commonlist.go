package common

// ListReq 通用列表请求参数
type ListReq struct {
	Current  int    `json:"current" form:"current"`
	PageSize int    `json:"pageSize" form:"pageSize"`
	Nickname string `json:"nickname" form:"nickname"`
	ApiName  string `json:"api_name" form:"api_name"`
	Path     string `json:"path" form:"path"`
}
