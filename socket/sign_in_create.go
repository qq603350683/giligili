package socket

import (
	"fmt"
	"giligili/constbase"
	"giligili/model"
	"giligili/serializer"
	"giligili/util"
	"github.com/jinzhu/gorm"
	"log"
	"time"
)

// 今天签到
func SignInCreate(params Params) {
	if model.UserInfo.UID == 0 {
		SendMessage(model.UserInfo.UID, serializer.JsonByte(constbase.SIGN_IN_FAIL, "签到失败", nil, ""))
		return
	}

	boolean := false
	sign_in := model.NewSignIn()
	today := time.Now().Format(util.DATE)

	if err := model.DB.Where("u_id = ? AND created_at >= ?", model.UserInfo.UID, today).First(sign_in).Error; err != nil && err != gorm.ErrRecordNotFound {
		log.Println(err.Error())
		SendMessage(model.UserInfo.UID, serializer.JsonByte(constbase.SIGN_IN_FAIL, "签到失败", nil, ""))
		return
	}

	if sign_in.SiID > 0 {
		SendMessage(model.UserInfo.UID, serializer.JsonByte(constbase.SIGN_IN_FAIL, "您今天已经签到了哦", nil, ""))
		return
	}

	sign_in.UID = model.UserInfo.UID
	sign_in.CreatedAt = time.Now()

	month_count := model.GetSignInMonthCount(model.UserInfo.UID, "")

	tx := model.DB.Begin()

	if err := tx.Create(sign_in).Error; err != nil {
		log.Println(err.Error())
		tx.Rollback()
		SendMessage(model.UserInfo.UID, serializer.JsonByte(constbase.SIGN_IN_FAIL, "签到失败", nil, ""))
		return
	}

	sign_in_prize := model.GetSignInPrize(model.UserInfo.UID, month_count + 1, "")
	if sign_in_prize != nil {
		//if boolean = sign_in_prize.PorpDetail.UseDB(tx).AddToBackpack(); boolean == false {
		//	tx.Rollback()
		//	SendMessage(model.UserInfo.UID, serializer.JsonByte(constbase.SIGN_IN_FAIL, "签到失败", nil, ""))
		//  return
		//}

		// 添加到背包
		switch sign_in_prize.PorpDetail.Type {
		case constbase.PROP_TYPE_GOLD:
			// 这里是签到奖励金币
			if boolean = sign_in_prize.PorpDetail.UseDB(tx).AddToUserGold(sign_in_prize.Quantity); boolean == false {
				tx.Rollback()
				SendMessage(model.UserInfo.UID, serializer.JsonByte(constbase.SIGN_IN_FAIL, "签到失败", nil, ""))
				return
			}
		case constbase.PROP_TYPE_DIAMOND:
			// 这里是签到奖励钻石
			if boolean = sign_in_prize.PorpDetail.UseDB(tx).AddToUserDiamond(sign_in_prize.Quantity); boolean == false {
				tx.Rollback()
				SendMessage(model.UserInfo.UID, serializer.JsonByte(constbase.SIGN_IN_FAIL, "签到失败", nil, ""))
				return
			}
		default:
			// 其他
			if boolean = sign_in_prize.PorpDetail.UseDB(tx).AddToBackpack(); boolean == false {
				tx.Rollback()
				SendMessage(model.UserInfo.UID, serializer.JsonByte(constbase.SIGN_IN_FAIL, "签到失败", nil, ""))
				return
			}
		}

	} else {
		log.Println("没有奖品...")
	}

	tx.Commit()

	msg := fmt.Sprintf("本月成功签到%d次", month_count + 1)

	SendMessage(model.UserInfo.UID, serializer.JsonByte(constbase.SIGN_IN_SUCCESS, msg, sign_in_prize, ""))
	return
}
