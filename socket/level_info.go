package socket

import (
	"giligili/constbase"
	"giligili/model"
	"giligili/serializer"
	"giligili/util"
	"net/http"
)

// 获取关卡详情
func GetLevelInfo(params Params) {
	if _, ok := params["l_id"]; !ok {
		SendMessage(model.UserInfo.UID, serializer.JsonByte(http.StatusInternalServerError, "参数异常", nil, "l_id 为 0"))
		return
	}

	l_id := util.StringToInt(params["l_id"])
	if l_id == 0 {
		SendMessage(model.UserInfo.UID, serializer.JsonByte(http.StatusInternalServerError, "系统繁忙", nil, "level.l_id 为 0"))
		return
	}

	level := model.GetLevelInfo(l_id)

	SendMessage(model.UserInfo.UID, serializer.JsonByte(constbase.LEVEL_INFO, "success", level, ""))
}
