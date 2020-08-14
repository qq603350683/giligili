package model

import "time"

type UserTokens struct {
	UtID int `json:"ut_id" gorm:"column:ut_id;type:int(10) unsigned auto_increment; primary_key;"`
	UID int `json:"u_id" gorm:"column:u_id;type:int(10) unsigned; not null; index:idx_u_id; comment:'用户ID'"`
	CreatedAt time.Time `json:"created_at" gorm:"type:datetime; not null; comment:'创建时间'"`
	ExpiredAt time.Time `json:"expired_at" gorm:"type:datetime; not null; comment:'过期时间'"`
}
