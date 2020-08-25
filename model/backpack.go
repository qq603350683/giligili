package model

import (
	"giligili/constbase"
	"time"
)

// 背包道具
type Backpack struct {
	BID int `json:"-" gorm:"column:b_id;type:int(10) unsigned auto_increment;not null;primary_key;comment:'背包物品主键ID'"`
	UID int `json:"u_id" gorm:"column:u_id;type:int(10) unsigned; not null;default:0;index:idx_u_id;comment:'用户ID 来自 users 表的 u_id'"`
	PID int `json:"p_id" gorm:"column:p_id;type:int(10) unsigned;not null;default:0;comment:'道具ID 来自 prop 表的 p_id'"`
	PropDetail *Prop `json:"prop_detail" comment:"道具详情"`
	Quantity int `json:"quantity" gorm:"-" comment:"数量"`
	//Type string `json:"type" gorm:"column:type;type:enum('bullet_enhancer','bullet_speed_enhancer', 'skill_enhancer', 'skill_speed_enhancer');not null;comment:'道具分类'"`
	//Title string `json:"title" gorm:"column:title;type:varchar(50);not null;comment:'标题'"`
	//Remark string `json:"-" gorm:"column:remark;type:varchar(50);not null;comment:'备注说明、领取途径等'"`
	IsUse int8 `json:"-" gorm:"column:is_use;type:int(1);not null;default:0;comment:'是否已使用 0 - 未使用 1 - 已使用'"`
	UseAt time.Time `json:"-" gorm:"type:datetime;not null;default:'1000-01-01 00:00:00'; comment:'使用时间'"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at; type:datetime; not null; comment:'创建时间'"`
}

func NewBackpack() *Backpack {
	return &Backpack{
		BID:        0,
		UID:        0,
		PID:        0,
		PropDetail: nil,
		Quantity:   0,
		IsUse:      0,
		UseAt:      time.Time{},
		CreatedAt:  time.Time{},
	}
}

func (backpack *Backpack) Use() bool {
	if backpack.IsUse == constbase.YES {

	}
}