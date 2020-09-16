package service

import "giligili/model"

type UserPlanChangeParams struct {
	UpID int `json:"up_id"`
}

func NewUserPlanChangeParams() *UserPlanChangeParams {
	return &UserPlanChangeParams{}
}

func UserPlanChange(up_id int) bool {

	bool := model.UserInfo.ChangePlan(up_id)
	if bool == false {
		return false
	}

	return true
}