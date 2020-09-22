package service

import (
	"giligili/constbase"
	"giligili/model"
	"giligili/serializer"
	"giligili/util"
	"github.com/jinzhu/gorm"
	"log"
	"time"
)

// 今天签到
func CreateSignIn(u_id int) (int, []byte) {
	if u_id == 0 {
		return u_id, serializer.JsonByte(constbase.SIGN_IN_FAIL, "签到失败", nil, "")
	}

	boolean := false
	sign_in := &model.SignIn{}
	today := time.Now().Format(util.DATE)

	if err := model.DB.Where("u_id = ? AND created_at >= ?", u_id, today).First(sign_in).Error; err != nil && err != gorm.ErrRecordNotFound {
		log.Println(err.Error())
		return u_id, serializer.JsonByte(constbase.SIGN_IN_FAIL, "签到失败", nil, "")
	}

	if sign_in.SiID > 0 {
		return u_id, serializer.JsonByte(constbase.SIGN_IN_FAIL, "您今天已经签到了哦", nil, "")
	}

	sign_in.UID = u_id
	sign_in.CreatedAt = time.Now()

	month_count := model.GetSignInMonthCount(u_id, "")

	tx := model.DB.Begin()

	if err := tx.Create(sign_in).Error; err != nil {
		log.Println(err.Error())
		tx.Rollback()
		return u_id, serializer.JsonByte(constbase.SIGN_IN_FAIL, "签到失败", nil, "")
	}

	sign_in_prize := model.GetSignInPrize(u_id, month_count + 1, "")
	if sign_in_prize != nil {
		if boolean = sign_in_prize.PorpDetail.AddToBackpack(); boolean == false {
			tx.Rollback()
			return u_id, serializer.JsonByte(constbase.SIGN_IN_FAIL, "签到失败", nil, "")
		}

		//return u_id, serializer.JsonByte(constbase.SIGN_IN_FAIL, "签到失败", nil, "")

		// 添加到背包
		//switch sign_in_prize.PorpDetail.Type {
		//case constbase.PROP_TYPE_GOLD:
		//	b = sign_in_prize.PorpDetail.AddToUserGold(sign_in_prize.Quantity)
		//case constbase.PROP_TYPE_DIAMOND:
		//	b = sign_in_prize.PorpDetail.AddToUserDiamond(sign_in_prize.Quantity)
		//default:
		//	b = sign_in_prize.PorpDetail.AddToBackpack()
		//}

	} else {
		log.Println("没有奖品...")
	}

	tx.Commit()

	return u_id, serializer.JsonByte(constbase.SIGN_IN_SUCCESS, "签到成功", sign_in_prize, "")
}