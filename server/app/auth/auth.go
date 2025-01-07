package auth

import (
	"gorm.io/gorm"
	"time"
)

type Auth struct {
	UserInfo *UserInfo
}

// UserInfo 用户信息
type UserInfo struct {
	ID                     int       `gorm:"primaryKey;autoIncrement;comment:'id'" json:"id"`
	AccountID              string    `gorm:"type:varchar(32);not null;default:'';comment:'账号id'" json:"accountId"`
	Name                   string    `gorm:"type:varchar(64);not null;default:'';comment:'昵称'" json:"name"`
	Account                string    `gorm:"type:varchar(32);not null;default:'';comment:'账号'" json:"account"`
	Email                  string    `gorm:"type:varchar(255);not null;default:'';comment:'邮箱'" json:"email"`
	Avatar                 string    `gorm:"type:varchar(255);not null;default:'';comment:'头像'" json:"avatar"`
	Organization           string    `gorm:"type:varchar(32);not null;default:'';comment:'部门(en)'" json:"organization"`
	OrganizationName       string    `gorm:"type:varchar(32);default:null;comment:'部门(zh)'" json:"organizationName"`
	Job                    string    `gorm:"type:varchar(32);not null;default:'';comment:'职位(en)'" json:"job"`
	JobName                string    `gorm:"type:varchar(32);not null;default:'';comment:'职位(zh)'" json:"jobName"`
	Location               string    `gorm:"type:varchar(32);not null;default:'';comment:'地区(en)'" json:"location"`
	LocationName           string    `gorm:"type:varchar(32);not null;default:'';comment:'地区(zh)'" json:"locationName"`
	Introduction           string    `gorm:"type:varchar(255);not null;default:'';comment:'个人介绍'" json:"introduction"`
	PersonalWebsite        string    `gorm:"type:varchar(255);not null;default:'';comment:'个人主页'" json:"personalWebsite"`
	Phone                  string    `gorm:"type:varchar(11);not null;default:'';comment:'手机号'" json:"phone"`
	Status                 int8      `gorm:"not null;default:1;comment:'状态:[1=正常,2=禁用中]'" json:"status"`
	Password               string    `gorm:"type:varchar(64);not null;default:'';comment:'密码串'" json:"-"`
	RegistrationDate       time.Time `gorm:"not null;default:CURRENT_TIMESTAMP;comment:'注册时间'""`
	DeleteDate             string    `gorm:"default:null;comment:'删除日期'" json:"deleteDate"`
	FormatRegistrationDate string    `json:"registrationDate" gorm:"-"`
}

// TableName 指定表名
func (*UserInfo) TableName() string {
	return "sys_admin_user"
}

func (u *UserInfo) AfterFind(tx *gorm.DB) (err error) {
	u.FormatRegistrationDate = u.RegistrationDate.Format(time.DateTime)
	return nil
}
