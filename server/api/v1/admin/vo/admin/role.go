package admin

// RoleListReq 角色列表请求参数
type RoleListReq struct {
	Current  int    `json:"current" form:"current"`
	PageSize int    `json:"pageSize" form:"pageSize"`
	Name     string `json:"name" form:"name"`
}

// RoleAuthReq 角色授权请求参数
type RoleAuthReq struct {
	RoleId int   `json:"role_id" form:"role_id"`
	Add    []any `json:"add" form:"add"`
	Remove []any `json:"remove" form:"remove"`
}
type SetRoleReq struct {
	Uid    int   `json:"uid" form:"uid"`
	Add    []any `json:"add"`
	Remove []any `json:"remove"`
}
