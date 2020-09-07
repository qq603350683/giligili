package service

import (
	"giligili/model"
)

type BackpackList struct {
	Plans     []model.UserPlan `json:"plans"`
	Backpacks []model.Backpack `json:"backpacks"`
}

func GetBackpackList(u_id int) BackpackList {
	backpack_list := BackpackList{
		Plans:     nil,
		Backpacks: nil,
	}

	backpack_list.Backpacks = model.GetBackpacks(model.UserInfo.UID)

	backpack_list.Plans = model.GetUserPlans(model.UserInfo.UID)

	return backpack_list
}
