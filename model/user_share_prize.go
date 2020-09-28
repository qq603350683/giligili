package model

import (
	"giligili/constbase"
	"giligili/util"
	"github.com/jinzhu/gorm"
	"log"
	"time"
)

type UserSharePrize struct {
	UspID int `json:"usp_id" gorm:"column:usp_id;type:int(10) unsigned auto_increment;not null;primary_key;comment:'签到奖品ID'"`
	UID int `json:"u_id" gorm:"column:u_id;type:int(10) unsigned;not null;default:0;index:idx_u_id;comment:'用户ID 关联来自 users 表的 u_id'"`
	PID int `json:"p_id" gorm:"column:p_id;type:int(10) unsigned;not null;default:0;comment:'道具ID 来自 prop 表的 p_id'"`
	Position string `json:"position" gorm:"column:type;type:enum('menu','index_button');not null;comment:'分享位置 menu - 菜单 index_button - 首页分享'"`
	Quantity int `json:"quantity" gorm:"column:quantity;type:int(10);not null;default:0;comment:'个数'"`
	PropDetail *Prop `json:"prop" gorm:"-" comment:"道具详情"`
	CreatedAt time.Time `json:"created_at" gorm:"type:datetime;not null; comment:'创建时间'"`
}

func NewUserSharePrize() *UserSharePrize {
	user_share_prize := new(UserSharePrize)
	user_share_prize.CreatedAt = time.Now()

	return user_share_prize
}

// 判断 position 是否在指定 positions 中
func InUserSharePrizePositions(position string) bool {
	positions := []string {constbase.SHARE_POSITION_MENU, constbase.SHARE_POSITION_INDEX_BUTTON}

	for _, p := range(positions) {
		if p == position {
			return true
		}
	}

	return false
}

// 检测用户今天是否已经分享了
func (user *User) TodayIsShare(position string) bool {
	date := time.Now().Format(util.DATE)

	user_share_prize := NewUserSharePrize()

	err := DB.Where("u_id = ? AND created_at >= ?", user.UID, date).First(user_share_prize).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false
		} else {
			log.Println(err.Error())
			return false
		}
	}

	return true
}
