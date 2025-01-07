package basemodel

import (
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	ID            int       `gorm:"primaryKey;autoIncrement;comment:'id'" json:"id"`
	CreatedTime   time.Time `json:"-" gorm:"column:created_time;type:timestamp;default:CURRENT_TIMESTAMP;comment:创建时间"`
	UpdatedTime   time.Time `json:"-" gorm:"column:updated_time;type:timestamp;default:CURRENT_TIMESTAMP;comment:修改时间" `
	DeletedTime   time.Time `json:"-" gorm:"column:deleted_time;type:timestamp;default:null;comment:删除时间" `
	FormatCreated string    `json:"created_time" gorm:"-"`
	FormatUpdated string    `json:"updated_time" gorm:"-" `
	FormatDeleted string    `json:"deleted_time" gorm:"-" `
}

// AfterFind 钩子函数，在查询数据后自动格式化时间字段
func (b *BaseModel) AfterFind(tx *gorm.DB) (err error) {
	b.FormatCreated = b.CreatedTime.Format(time.DateTime)
	b.FormatUpdated = b.UpdatedTime.Format(time.DateTime)
	if b.DeletedTime.IsZero() {
		b.FormatDeleted = "0"
	}
	return nil
}
