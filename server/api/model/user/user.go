package user_model

import (
	"nexwind.net/admin/server/api/model/basemodel"
	"time"
)

type UserInfo struct {
	Id       int    `json:"id"`
	NickName string `json:"nickName"`
	Pwd      string `json:"pwd"`
}

// Logout 用户退出
type Logout struct {
	Id          int       `json:"id"`
	Uid         int       `json:"uid"`
	Token       string    `json:"token"`
	ExpiredTime time.Time `json:"-" gorm:"column:expired_time;type:timestamp;comment:过期时间"`
	basemodel.BaseModel
}

func (Logout) TableName() string {
	return "sys_admin_logout"
}
