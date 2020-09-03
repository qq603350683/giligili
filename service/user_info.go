package service

import (
	"giligili/model"
)

func GetUserInfo(u_id int) *model.User {
	user := model.GetUserInfo(u_id)

	if model.IsDel(user.DelAt) {
		return nil
	}

	return user
}
