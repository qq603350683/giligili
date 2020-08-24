package service

import (
	"errors"
	"giligili/model"
	"giligili/util"
	"github.com/jinzhu/gorm"
	"time"
)

// 今天签到
func CreateSignIn(u_id int) (bool, error) {
	if u_id == 0 {
		return false, errors.New("用户ID不能为0")
	}

	sign_in := &model.SignIn{}

	today := time.Now().Format(util.DATE)

	err := model.DB.Where("u_id = ? AND created_at >= ?", model.UserInfo.UID, today).First(sign_in).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if sign_in.SiID > 0 {
		return false, errors.New("您今天已经签到了哦~")
	}

	sign_in.UID = u_id
	sign_in.CreatedAt = time.Now()

	err = model.DB.Create(sign_in).Error
	if err != nil {
		return false, err
	}

	return true, nil
}