package model

import (
	"log"
	"time"
)

type UserAdvRecord struct {
	UarID int `json:"uar_id" gorm:"column:uar_id;type:int(10) unsigned auto_increment;not null;primary_key;comment:'记录ID'"`
	UID int `json:"u_id" gorm:"column:u_id; type:int(10) unsigned; not null; default:0; index:idx_u_id; comment:'用户ID 来自 users 表的 u_id'"`
	Remark string `json:"remark" gorm:"column:remark;type:varchar(200);not null;comment:'备注'"`
	CreatedAt time.Time `json:"-" gorm:"column:created_at; type:datetime;not null;index:idx_created_at comment:'创建时间'"`
}


func NewUserAdvRecord() *UserAdvRecord {
	user_adv_record := new(UserAdvRecord)
	user_adv_record.CreatedAt = time.Now()

	return user_adv_record
}

func CreateUserAdvRecord(u_id int, remark string) bool {
	user_adv_record := NewUserAdvRecord()
	user_adv_record.UID = u_id
	user_adv_record.Remark = remark

	db := GetDB()

	if err := db.Create(user_adv_record).Error; err != nil {
		log.Printf("用户ID(%d)添加广告查看完毕记录失败", u_id)
		return false
	} else {
		return true
	}
}
