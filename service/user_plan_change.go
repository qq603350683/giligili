package service

import "giligili/model"

type UserPlanChangeParams struct {
	UpID int `json:"up_id"`
}

func NewUserPlanChangeParams() *UserPlanChangeParams {
	return &UserPlanChangeParams{}
}

func UserPlanChange(up_id int) bool {
	db := model.DB.Begin()

	bool := model.UserInfo.ChangePlan(up_id)
	if bool == false {
		db.Rollback()
		return false
	}

	db.Commit()

	return true
}