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

type SignInResult struct {
	SignInPrize *model.SignInPrize `json:"sign_in_prize"`
	UserGold int `json:"user_gold"`
	DiamondGold int `json:"user_diamond"`
}

func NewSignInResult() *SignInResult {
	return new(SignInResult)
}

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
		switch sign_in_prize.PropDetail.Type {
		case constbase.PROP_TYPE_GOLD:
			// 这里是签到奖励金币
			if boolean = sign_in_prize.PropDetail.AddToUserGold(sign_in_prize.Quantity); boolean == false {
				db.Rollback()
				SendMessage(model.UserInfo.UID, serializer.JsonByte(constbase.SIGN_IN_FAIL, "签到失败", nil, ""))
				return
			}
		case constbase.PROP_TYPE_DIAMOND:
			// 这里是签到奖励钻石
			if boolean = sign_in_prize.PropDetail.AddToUserDiamond(sign_in_prize.Quantity); boolean == false {
				db.Rollback()
				SendMessage(model.UserInfo.UID, serializer.JsonByte(constbase.SIGN_IN_FAIL, "签到失败", nil, ""))
				return
			}
		default:
			// 其他
			if boolean = sign_in_prize.PropDetail.AddToBackpack(); boolean == false {
				db.Rollback()
				SendMessage(model.UserInfo.UID, serializer.JsonByte(constbase.SIGN_IN_FAIL, "签到失败", nil, ""))
				return
			}
		}

	} else {
		log.Println("没有奖品...")
	}

	db.Commit()

	sign_in_result := NewSignInResult()
	sign_in_result.SignInPrize = sign_in_prize
	sign_in_result.UserGold = model.UserInfo.Gold
	sign_in_result.DiamondGold = model.UserInfo.Diamond

	msg := fmt.Sprintf("本月成功签到%d次", month_count + 1)

	SendMessage(model.UserInfo.UID, serializer.JsonByte(constbase.SIGN_IN_SUCCESS, msg, sign_in_result, ""))
	return
}
