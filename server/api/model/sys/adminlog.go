package sys

import "nexwind.net/admin/server/api/model/basemodel"

// AdminLog 系统管理员操作日志模型
type AdminLog struct {
	ID       int    `gorm:"primaryKey;autoIncrement;comment:'id'" json:"id"`
	UID      int    `gorm:"not null;default:0;comment:'用户id'" json:"uid"`
	Nickname string `gorm:"size:32;not null;comment:'管理员名称'" json:"nickname"`
	APIName  string `gorm:"size:255;not null;default:'';comment:'接口名称'" json:"api_name"`
	Path     string `gorm:"size:255;not null;default:'';comment:'path'" json:"path"`
	Method   string `gorm:"size:24;not null;default:'';comment:'method'" json:"method"`
	ApiID    int    `gorm:"not null;default:0;comment:'接口 id'" json:"api_id"`
	Param    string `gorm:"type:text;comment:'query 参数'" json:"param"`
	ClientIp string `gorm:"size:45;not null;default:'';comment:'客户端 ip'" json:"client_ip"`
	basemodel.BaseModel
}

// TableName 设置模型对应的表名
func (AdminLog) TableName() string {
	return "sys_admin_log"
}
