package socket

import (
	"giligili/constbase"
	"giligili/model"
	"giligili/serializer"
)

func GetUserPlans(params Params) {
	plans := model.GetUserPlans(model.UserInfo.UID)

	SendMessage(model.UserInfo.UID, serializer.JsonByte(constbase.OPEN_MY_PLAN_LIST, "success", plans, ""))
}