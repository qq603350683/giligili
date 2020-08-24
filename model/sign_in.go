package model

import "time"

type SignIn struct {
	SiID int `json:"si_id" gorm:"column:si_id;type:int(10) unsigned auto_increment;not null;primary_key;comment:'签到主键ID'"`
	UID int `json:"u_id" gorm:"column:u_id;type:int(10) unsigned;not null;default:0;index:idx_u_id;comment:'用户ID 关联来自 users 表的 u_id'"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;type:datetime;not null;comment:'签到时间'"`
}
