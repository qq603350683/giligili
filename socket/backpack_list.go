package socket

import (
	"giligili/constbase"
	"giligili/model"
	"giligili/serializer"
)

type BackpackList struct {
	Plans     []model.UserPlan `json:"plans"`
	Backpacks []model.Backpack `json:"backpacks"`
}

func GetBackpackList(params Params) {
	backpack_list := BackpackList{
		Plans:     nil,
		Backpacks: nil,
	}

	backpack_list.Backpacks = model.GetBackpacks(model.UserInfo.UID)

	backpack_list.Plans = model.GetUserPlans(model.UserInfo.UID)

	SendMessage(model.UserInfo.UID, serializer.JsonByte(constbase.BACKPACK_LIST, "success", backpack_list, ""))
}
