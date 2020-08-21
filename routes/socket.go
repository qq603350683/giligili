package routes

import (
	"encoding/json"
	"fmt"
	"giligili/model"
	"giligili/serializer"
	"giligili/service"
	"net/http"
)

type MsgDecode struct {
	Case string
	Content string
}

func Socket(msg []byte) []byte {
	fmt.Println(string(msg))

	// 解析json格式
	m := &MsgDecode{}
	err := json.Unmarshal(msg, m)
	if err != nil {
		return serializer.JsonByte(http.StatusOK, "数据解析错误", nil, err.Error())
	}

	switch m.Case {
	case "level/get":
		level_get := model.NewLevelGet()
		err = json.Unmarshal([]byte(m.Content), level_get)
		l, _ := model.GetLevelByID(level_get.LID)

		return serializer.JsonByte(http.StatusOK, "success", l, "")
	case "user":
		// 我的详情
		user, err := service.GetUserInfo(model.UID)
		if err != nil {
			return serializer.JsonByte(http.StatusInternalServerError, err.Error(), nil, err.Error())
		}

		return  serializer.JsonByte(http.StatusOK, "succes", user, "")
	}

	return serializer.JsonByte(http.StatusOK, "success", nil, "")
}
