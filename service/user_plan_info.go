package service

import (
	"giligili/model"
)

func GetUserPlanInfo(up_id int) (*model.UserPlan) {

	plan := model.GetUserPlanInfo(up_id)

	return plan
}