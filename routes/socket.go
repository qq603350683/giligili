package routes

import (
	"encoding/json"
	"giligili/model"
	"giligili/serializer"
	"giligili/service"
	"log"
	"net/http"
)

type MsgDecode struct {
	Case string
	Content string
}

func Socket(msg []byte) []byte {
	log.Printf("接收数据: %s", string(msg))

	// 解析json格式 {"case": "sign_in/count", "content": "{\"l_id\": 1}"}
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
		user, err := service.GetUserInfo(model.UserInfo.UID)
		if err != nil {
			return serializer.JsonByte(http.StatusInternalServerError, err.Error(), nil, err.Error())
		}

		return  serializer.JsonByte(http.StatusOK, "success", user, "")
	case "sign_in/create":
		// 今天签到
		bool, err := service.CreateSignIn(model.UserInfo.UID)
		if bool == false {
			return serializer.JsonByte(http.StatusInternalServerError, err.Error(), nil, err.Error())
		}

		return  serializer.JsonByte(http.StatusOK, "success", nil, "")
	case "sign_in/count":
		// 当前用户本月总签到次数
		count, err := service.GetSignInMonthCount(model.UserInfo.UID)
		if err != nil {
			return serializer.JsonByte(http.StatusInternalServerError, err.Error(), nil, err.Error())
		}

		return  serializer.JsonByte(http.StatusOK, "success", count, "")
	case "forward/create":
		prop, err := service.CreateForward(model.UserInfo.UID)
		if err != nil {
			return serializer.JsonByte(http.StatusInternalServerError, err.Error(), nil, err.Error())
		}

		return  serializer.JsonByte(http.StatusOK, "success", prop, "")
	}

	return serializer.JsonByte(http.StatusOK, "success", nil, "")
}
