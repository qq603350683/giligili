package service

import (
	"giligili/model"
	"log"
)

type BackpackList struct {
	plans     []model.UserPlan `json:"plans"`
	backpacks []model.Backpack `json:"backpacks"`
}

func GetBackpackList(u_id int) BackpackList {
	plans     := []model.UserPlan{}
	backpacks := []model.Backpack{}

	backpack_list := BackpackList{
		plans:     nil,
		backpacks: nil,
	}

	err := model.DB.Raw("SELECT *, COUNT(*) as quantity FROM backpacks WHERE u_id = ? AND is_use = 0 GROUP BY p_id", u_id).Find(&backpacks).Error
	if err != nil {
		log.Println(err.Error())
		return backpack_list
	}

	if len(backpacks) > 0 {
		for index, backpack := range(backpacks) {
			backpack.PropDetail = model.GetPropInfo(backpack.PID)

			backpacks[index] = backpack
		}

		backpack_list.backpacks = backpacks
	}

	plans = model.GetUserPlans(model.UserInfo.UID)
	backpack_list.plans = plans

	log.Println(backpack_list)

	return backpack_list
}
