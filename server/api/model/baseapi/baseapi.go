package baseapi

import "nexwind.net/admin/server/api/model/basemodel"

// SysBaseApi 系统api
type SysBaseApi struct {
	basemodel.BaseModel
	Title     string `gorm:"type:varchar(64);not null;default:'';comment:'接口名称'" json:"title"`              // 接口名称
	Scope     string `gorm:"type:varchar(32);not null;default:'';comment:'范围'" json:"scope"`                // 范围
	Group     string `gorm:"type:varchar(64);not null;default:'';comment:'分组'" json:"group"`                // 分组
	Method    string `gorm:"type:varchar(24);not null;default:'GET';comment:'请求类型:GET,POST'" json:"method"` // 请求类型:GET,POST
	Path      string `gorm:"type:varchar(255);not null;comment:'路径'" json:"path"`                           // 路径
	NeedLogin uint8  `gorm:"type:tinyint;not null;default:1;comment:'是否需要登录'" json:"need_login"`            // 是否需要登录
	NeedAuth  uint8  `gorm:"type:tinyint;not null;default:1;comment:'是否需要鉴权'" json:"need_auth"`             // 是否需要鉴权
}

// TableName 返回表名
func (SysBaseApi) TableName() string {
	return "sys_base_api"
}
