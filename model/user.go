package model

import "time"

type User struct {
	UID int `json:"u_id" gorm:"column:u_id;type:int(10) unsigned auto_increment;primary_key;"`
	UpID int `json:"up_id" gorm:"column:up_id;type:int(10) unsigned; not null; default:0; comment:'飞机ID'"`
	Plan *UserPlan `json:"plan" comment:"飞机详情"`
	CreatedAt time.Time `json:"created_at" gorm:"type:datetime;not null; comment:'创建时间'"`
	UpdatedAt time.Time `json:"-" gorm:"type:datetime;not null; comment:'更新时间'"`
	DelAt time.Time `json:"-" gorm:"type:datetime;not null;default:'1000-01-01 00:00:00'; comment:'删除时间'"`
}

