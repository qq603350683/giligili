package model

import (
	"github.com/jinzhu/gorm"
	"log"
	"time"
)

type SignInPrize struct {
	SipID int `json:"sip_id" gorm:"column:sip_id;type:int(10) unsigned auto_increment;not null;primary_key;comment:'签到奖品ID'"`
	PID int `json:"p_id" gorm:"column:p_id;type:int(10) unsigned;not null;default:0;comment:'道具ID 来自 prop 表的 p_id'"`
	Quantity int `json:"quantity" gorm:"column:quantity;type:int(10);not null;default:0;comment:'个数'"`
	Time string `json:"time" gorm:"column:time;type:char(6);not null;default:'';index:idx_time;comment:'年月'"`
	GrandTotal int8 `json:"grand_total" gorm:"column:grand_total;type:tinyint(3);not null;default:0;comment:'累计天数'"`
	PropDetail *Prop `json:"prop" gorm:"-" comment:"道具详情"`
}

func NewSignInPrize() *SignInPrize {
	return new(SignInPrize)
}

func GetSignInPrize(u_id int, total int, monthday string) *SignInPrize {
	if total == 0 {
		return nil
	}

	if monthday == "" {
		monthday = time.Now().Format("200601")
	}

	// 检测是否重复领取
	user_sign_in_prize := NewUserSignInPrize()
	err := DB.Where("u_id = ? AND time = ? AND grand_total = ?", u_id, monthday, total).First(user_sign_in_prize).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Println(err.Error())
		return nil
	}

	if user_sign_in_prize.UsipID > 0 {
		return nil
	}

	sign_in_prize := NewSignInPrize()

	err = DB.Where("time = ? AND grand_total = ?", monthday, total).First(sign_in_prize).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil
		}

		log.Println(err.Error())
		return nil
	}

	sign_in_prize.PropDetail = GetPropInfo(sign_in_prize.PID)

	return sign_in_prize
}


