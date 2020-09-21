package socket

import (
	"giligili/constbase"
	"giligili/model"
	"giligili/serializer"
	"giligili/service"
)

func GetUserInfo(params Params) {
	user := service.GetUserInfo(model.UserInfo.UID)

	SendMessage(model.UserInfo.UID, serializer.JsonByte(constbase.LOGIN_USER_INFO, "success", user, ""))
}
