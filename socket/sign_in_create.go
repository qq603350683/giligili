package socket

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
func SignInCreate(params GetParams) {
	if model.UserInfo.UID == 0 {
		SendMessage(model.UserInfo.UID, serializer.JsonByte(constbase.SIGN_IN_FAIL, "签到失败", nil, ""))
	}

	boolean := false
	sign_in := model.NewSignIn()
	today := time.Now().Format(util.DATE)

	if err := model.DB.Where("model.UserInfo.UID = ? AND created_at >= ?", model.UserInfo.UID, today).First(sign_in).Error; err != nil && err != gorm.ErrRecordNotFound {
		log.Println(err.Error())
		SendMessage(model.UserInfo.UID, serializer.JsonByte(constbase.SIGN_IN_FAIL, "签到失败", nil, ""))
	}

	if sign_in.SiID > 0 {
		SendMessage(model.UserInfo.UID, serializer.JsonByte(constbase.SIGN_IN_FAIL, "您今天已经签到了哦", nil, ""))
	}

	sign_in.UID = model.UserInfo.UID
	sign_in.CreatedAt = time.Now()

	month_count := model.GetSignInMonthCount(model.UserInfo.UID, "")

	tx := model.DB.Begin()

	if err := tx.Create(sign_in).Error; err != nil {
		log.Println(err.Error())
		tx.Rollback()
		SendMessage(model.UserInfo.UID, serializer.JsonByte(constbase.SIGN_IN_FAIL, "签到失败", nil, ""))
	}

	sign_in_prize := model.GetSignInPrize(model.UserInfo.UID, month_count + 1, "")
	if sign_in_prize != nil {
		if boolean = sign_in_prize.PorpDetail.UseDB(tx).AddToBackpack(); boolean == false {
			tx.Rollback()
			SendMessage(model.UserInfo.UID, serializer.JsonByte(constbase.SIGN_IN_FAIL, "签到失败", nil, ""))
		}

		//SendMessage(model.UserInfo.UID, serializer.JsonByte(constbase.SIGN_IN_FAIL, "签到失败", nil, ""))

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

	SendMessage(model.UserInfo.UID, serializer.JsonByte(constbase.SIGN_IN_SUCCESS, "签到成功", sign_in_prize, ""))
}
