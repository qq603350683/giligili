package model

import "time"

type UserSignInPrize struct {
	UsipID int `json:"usip_id" gorm:"column:usip_id;type:int(10) unsigned auto_increment;not null;primary_key;comment:'用户签到奖品ID'"`
	UID int `json:"u_id" gorm:"column:u_id;type:int(10) unsigned;not null;default:0;comment:'用户ID 来自 users 表的 u_id'"`
	PID int `json:"p_id" gorm:"column:p_id;type:int(10) unsigned;not null;default:0;comment:'道具ID 来自 prop 表的 p_id'"`
	Time string `json:"time" gorm:"column:time;type:char(6);not null;default:'';index:idx_time;comment:'年月'"`
	GrandTotal int8 `json:"grand_total" gorm:"column:grand_total;type:tinyint(3);not null;default:0;comment:'累计天数'"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at; type:datetime; not null; comment:'创建时间'"`
	PorpDetail *Prop `json:"prop" gorm:"-" comment:"道具详情"`
}

func NewUserSignInPrize() *UserSignInPrize {
	return new(UserSignInPrize)
}
