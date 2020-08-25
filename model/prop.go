package model

import (
	"github.com/jinzhu/gorm"
	"log"
	"time"
)

type Prop struct {
	PID int `json:"p_id" gorm:"column:p_id;type:int(10) unsigned auto_increment;not null;primary_key;comment:'道具ID'"`
	Type string `json:"type" gorm:"column:type;type:enum('bullet_enhancer','bullet_speed_enhancer', 'skill_enhancer', 'skill_speed_enhancer');not null;comment:'道具分类'"`
	Image string `json:"image" gorm:"column:image;type:char(35);default:'';not null;comment:'图标'"`
	Title string `json:"title" gorm:"column:title;type:varchar(50);default:'';not null;comment:'标题'"`
	Remark string `json:"-" gorm:"column:remark;type:varchar(50);default:'';not null;comment:'备注说明、领取途径等'"`
	CreatedAt time.Time `json:"-" gorm:"column:created_at;type:datetime;not null;comment:'签到时间'"`
}

// 展示在前端的格式
type PropShow struct {
	*Prop `comment:"道具详情"`
	Quantity int `json:"quantity" comment:"数量"`
}

// 获取道具详情
func GetPropInfo(p_id int) *Prop {
	if p_id == 0 {
		return nil
	}

	prop := &Prop{}

	err := DB.Where("p_id = ?", p_id).First(prop).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Printf("道具ID(%d)找不到记录: ", p_id)
		} else {
			log.Println(err.Error())
		}

		return nil
	}

	return prop
}

// 道具加入到背包
func (prop *Prop) AddToBackpack() bool {
	backpack := NewBackpack()

	backpack.UID = UserInfo.UID
	backpack.PID = prop.PID

	err := DB.Create(backpack).Error
	if err != nil {
		log.Printf("用户ID(%d)道具领取失败", UserInfo.UID)
		return false
	}

	return true
}
