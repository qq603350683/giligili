package service

import (
	"giligili/model"
	"github.com/jinzhu/gorm"
	"log"
	"time"
)

/**
 * 获取累计的奖品道具
 * @param total 累计天数
 */
func GetGrandTotalPrize(total int) *model.SignInPrize {
	if total == 0 {
		return nil
	}

	_time := time.Now().Format("200601")

	// 检测是否重复领取
	user_sign_in_prize := &model.UserSignInPrize{}
	err := model.DB.Where("u_id = ? AND time = ? AND grand_total = ?", model.UserInfo.UID, _time, total).First(user_sign_in_prize).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Println(err.Error())
		return nil
	}

	if user_sign_in_prize.UsipID > 0 {
		return nil
	}

	sign_in_prize := &model.SignInPrize{}

	err = model.DB.Where("time = ? AND grand_total = ?", _time, total).First(sign_in_prize).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil
		}

		log.Println(err.Error())
		return nil
	}

	sign_in_prize.PorpDetail = model.GetPropInfo(sign_in_prize.PID)

	model.DB.Begin()

	// 添加到背包
	b := sign_in_prize.PorpDetail.AddToBackpack()
	if b == false {
		model.DB.Rollback()
		return nil
	}

	// 添加用户兑奖记录
	user_sign_in_prize.UID = model.UserInfo.UID
	user_sign_in_prize.PID = sign_in_prize.PID
	user_sign_in_prize.Time = sign_in_prize.Time
	user_sign_in_prize.GrandTotal = sign_in_prize.GrandTotal
	user_sign_in_prize.CreatedAt = time.Now()

	err = model.DB.Create(user_sign_in_prize).Error
	if err != nil {
		model.DB.Rollback()
		return nil
	}

	model.DB.CommonDB()

	return sign_in_prize
}
