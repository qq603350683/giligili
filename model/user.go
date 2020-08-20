package model

import "time"

// 全局u_id
var UID int

type User struct {
	UID int `json:"u_id" gorm:"column:u_id;type:int(10) unsigned auto_increment;primary_key;"`
	PID int `json:"p_id" gorm:"column:p_id;type:int(10) unsigned; default:0; comment:'飞机ID'"`
	CreatedAt time.Time `json:"created_at" gorm:"type:datetime;not null; comment:'创建时间'"`
	UpdatedAt time.Time `json:"update_at" gorm:"type:datetime;not null; comment:'更新时间'"`
	DelAt time.Time `json:"del_at" gorm:"type:datetime;not null;default:'1000-01-01 00:00:00'; comment:'删除时间'"`
}

