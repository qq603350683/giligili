package socket

import (
	"giligili/model"
	"giligili/serializer"
	"giligili/util"
	"net/http"
)

func UserPlanChange(params Params) {
	if _, ok := params["up_id"]; !ok {
		SendMessage(model.UserInfo.UID, serializer.JsonByte(http.StatusInternalServerError, "参数异常", nil, "up_id 为 0"))
		return
	}

	up_id := util.StringToInt(params["up_id"])
	if l_id == 0 {
		SendMessage(model.UserInfo.UID, serializer.JsonByte(http.StatusInternalServerError, "系统繁忙", nil, "up_id 为 0"))
		return
	}

	model.UserInfo.ChangePlan(up_id)
}
