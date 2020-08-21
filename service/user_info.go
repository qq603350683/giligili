package service

import (
	"errors"
	"giligili/model"
)

func GetUserInfo(u_id int) (*model.User, error) {
	if u_id == 0 {
		return nil, errors.New("用户ID不能为0")
	}

	user := &model.User{}

	err := model.DB.Where("u_id = ?", u_id).First(user).Error
	if err != nil {
		return nil, err
	}

	if model.IsDel(user.DelAt) {
		return nil, errors.New("当前用户已经删除")
	}

	if user.UpID > 0 {
		plan, err := GetUserPlanInfo(user.UpID)
		if err != nil {
			return nil, err
		}

		user.Plan = plan
	}

	return user, nil
}
