package user_model

import "nexwind.net/admin/server/api/model/basemodel"

// Role 角色模型
type Role struct {
	Name        string `json:"name" gorm:"column:name"`
	Description string `json:"description" gorm:"column:description"`
	basemodel.BaseModel
}

func (r Role) TableName() string {
	return "sys_admin_role"
}
