package model

type SignInPrize struct {
	SipID int `json:"sip_id" gorm:"column:sip_id;type:int(10) unsigned auto_increment;not null;primary_key;comment:'签到奖品ID'"`
	PID int `json:"p_id" gorm:"column:p_id;type:int(10) unsigned;not null;default:0;comment:'道具ID 来自 prop 表的 p_id'"`
	Quantity int `json:"quantity" gorm:"column:quantity;type:int(10);not null;default:0;comment:'个数'"`
	Time string `json:"time" gorm:"column:time;type:char(6);not null;default:'';index:idx_time;comment:'年月'"`
	GrandTotal int8 `json:"grand_total" gorm:"column:grand_total;type:tinyint(3);not null;default:0;comment:'累计天数'"`
	PorpDetail *Prop `json:"prop" gorm:"-" comment:"道具详情"`
}


