package service

import (
	"giligili/model"
	"log"
)

func GetBackpackList(u_id int) []model.Backpack {
	backpacks := []model.Backpack{}

	err := model.DB.Raw("SELECT *, COUNT(*) as quantity FROM backpacks WHERE u_id = ? AND is_use = 0 GROUP BY p_id", u_id).Find(&backpacks).Error
	if err != nil {
		log.Println(err.Error())
		return nil
	}

	if len(backpacks) == 0 {
		return nil
	}

	for index, backpack := range(backpacks) {
		backpack.PropDetail = model.GetPropInfo(backpack.PID)

		backpacks[index] = backpack
	}

	return backpacks
}
