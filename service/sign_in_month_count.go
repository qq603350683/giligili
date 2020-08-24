package service

import (
	"errors"
	"giligili/model"
)

// 获取用户在当月签到的次数
func GetSignInMonthCount(u_id int) (int, error) {
	if u_id == 0 {
		return 0, errors.New("用户ID不能为0")
	}

	month := "2020-08-01"

	count := 0

	err := model.DB.Model(&model.SignIn{}).Where("u_id = ? AND created_at >= ?", u_id, month).Count(&count).Error
	if err != nil {
		return 0, err
	}

	return count, nil
}
