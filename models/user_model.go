package models

import "time"

// User 用户模型
type User struct {
	Id         uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Phone      string    `json:"phone" gorm:"size:20;uniqueIndex;not null;comment:手机号"`
	Password   string    `json:"-" gorm:"size:255;not null;comment:密码"` // json:"-" 表示不序列化到JSON
	CreateTime time.Time `json:"create_time" gorm:"autoCreateTime;comment:创建时间"`
	UpdateTime time.Time `json:"update_time" gorm:"autoUpdateTime;comment:更新时间"`
	Yn         bool      `json:"yn" gorm:"default:true;comment:软删除标记,true=正常,false=删除"`
}

// TableName 指定表名
func (User) TableName() string {
	return "users"
}
