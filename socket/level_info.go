package socket

import (
	"giligili/model"
	"giligili/serializer"
	"net/http"
)

func GetLevelInfo(params GetParams) {
	if _, ok := params["l_id"]; !ok {
		SendMessage(model.UserInfo.UID, serializer.JsonByte(http.StatusInternalServerError, "数据异常", nil, "level.l_id 为 0"))
		return
	}

	//l_id := util.ToInt(params["l_id"].(float64))
	//if l_id == 0 {
	//	SendMessage(model.UserInfo.UID, serializer.JsonByte(http.StatusInternalServerError, "数据异常", nil, "level.l_id 为 0"))
	//	return
	//}
	//
	//level := model.GetLevelInfo(l_id)
	//
	//SendMessage(model.UserInfo.UID, serializer.JsonByte(http.StatusInternalServerError, "success", level, ""))
}
