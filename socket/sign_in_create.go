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

	db := model.DBBegin()

	defer model.CancelDB()

	if err := db.Create(sign_in).Error; err != nil {
		log.Println(err.Error())
		db.Rollback()
		SendMessage(model.UserInfo.UID, serializer.JsonByte(constbase.SIGN_IN_FAIL, "签到失败", nil, ""))
		return
	}

	sign_in_prize := model.GetSignInPrize(model.UserInfo.UID, month_count + 1, "")
	if sign_in_prize != nil {
		// 添加到背包
		switch sign_in_prize.PorpDetail.Type {
		case constbase.PROP_TYPE_GOLD:
			// 这里是签到奖励金币
			if boolean = sign_in_prize.PorpDetail.AddToUserGold(sign_in_prize.Quantity); boolean == false {
				db.Rollback()
				SendMessage(model.UserInfo.UID, serializer.JsonByte(constbase.SIGN_IN_FAIL, "签到失败", nil, ""))
				return
			}
		case constbase.PROP_TYPE_DIAMOND:
			// 这里是签到奖励钻石
			if boolean = sign_in_prize.PorpDetail.AddToUserDiamond(sign_in_prize.Quantity); boolean == false {
				db.Rollback()
				SendMessage(model.UserInfo.UID, serializer.JsonByte(constbase.SIGN_IN_FAIL, "签到失败", nil, ""))
				return
			}
		default:
			// 其他
			if boolean = sign_in_prize.PorpDetail.AddToBackpack(); boolean == false {
				db.Rollback()
				SendMessage(model.UserInfo.UID, serializer.JsonByte(constbase.SIGN_IN_FAIL, "签到失败", nil, ""))
				return
			}
		}

	} else {
		log.Println("没有奖品...")
	}

	db.Commit()

	msg := fmt.Sprintf("本月成功签到%d次", month_count + 1)

	SendMessage(model.UserInfo.UID, serializer.JsonByte(constbase.SIGN_IN_SUCCESS, msg, sign_in_prize, ""))
	return
}
