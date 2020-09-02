package model

import (
	"log"
	"time"
)

// 闯关记录
type UserPassLevel struct {
	UplID int `json:"upl_id" gorm:"column:upl_id;type:int(10) unsigned auto_increment;primary_key"`
	UID int `json:"u_id" gorm:"column:u_id; type:int(10) unsigned; not null; default:0; index:idx_u_id; comment:'用户ID 来自 users 表的 u_id'"`
	LID int `json:"l_id" gorm:"column:l_id;type:int(10) unsigned; not null; default:0; comment:'关卡ID'"`
	IsSucess int8 `json:"is_success" gorm:"column:is_success;type:tinyint(1) unsigned;not null;default:0;comment:'是否成功闯关 0 - 下架 1 - 上架'"`
	Gold int `json:"gold" gorm:"column:gold;type:int(10) unsigned; not null; default:0; comment:'通关奖励金币'"`
	Diamond int `json:"diamond" gorm:"column:diamond;type:int(10) unsigned; not null; default:0; comment:'通关奖励钻石'"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at; type:datetime; not null; comment:'创建时间'"`
}

func NewUserPassLevel() *UserPassLevel {
	return &UserPassLevel{}
}

// 顺利通关
func (user *User) PassLevel(l_id int, is_success int8, gold int, diamond int) bool {
	user_pass_level := NewUserPassLevel()

	user_pass_level.UID = UserInfo.UID
	user_pass_level.LID = l_id
	user_pass_level.IsSucess = is_success
	user_pass_level.Gold = gold
	user_pass_level.Diamond = diamond

	err := DB.Create(user_pass_level).Error
	if err != nil {
		log.Println(err.Error())
		return false
	}

	return true
}

