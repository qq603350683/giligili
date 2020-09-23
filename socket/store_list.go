package socket

import (
	"giligili/constbase"
	"giligili/model"
	"giligili/serializer"
)

func GetStoreList(params Params) {
	stores := model.GetStoreList(params["type"])

	SendMessage(model.UserInfo.UID, serializer.JsonByte(constbase.STORE_LIST, "success", stores, ""))
}
