package socket

import (
	"giligili/constbase"
	"giligili/model"
	"giligili/serializer"
	"giligili/util"
	"net/http"
)

func StoreBuy(params Params) {
	s_id := 0

	if _, ok := params["s_id"]; !ok {
		SendMessage(model.UserInfo.UID, serializer.JsonByte(http.StatusInternalServerError, "请选择需要购买的商品", nil, "参数 s_id 为 不存在"))
		return
	}

	s_id = util.StringToInt(params["s_id"])

	store := model.GetSroteInfo(s_id)
	if store == nil {
		SendMessage(model.UserInfo.UID, serializer.JsonByte(http.StatusInternalServerError, "当前购买的商品已不存在", nil, "socket.StoreBuy 商品不存在"))
		return
	}

	db := model.DBBegin()

	defer model.CancelDB()

	boolean := store.Buy()
	if boolean == false {
		db.Rollback()
		SendMessage(model.UserInfo.UID, serializer.JsonByte(http.StatusInternalServerError, "购买失败", nil, "socket.StoreBuy 异常001"))
		return
	}

	db.Commit()

	SendMessage(model.UserInfo.UID, serializer.JsonByte(constbase.STORE_BUY_SUCCESS, "购买成功", store, ""))
}
