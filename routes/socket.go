package routes

import (
	"encoding/json"
	"giligili/constbase"
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
	case "sign_in/prize":
		// 签到奖励
		count, err := service.GetSignInMonthCount(model.UserInfo.UID)
		if err != nil {
			return serializer.JsonByte(http.StatusInternalServerError, err.Error(), nil, err.Error())
		}

		prop := service.GetGrandTotalPrize(count)

		return  serializer.JsonByte(http.StatusOK, "success", prop, "")
	case "sign_in/count":
		// 当前用户本月总签到次数
		count, err := service.GetSignInMonthCount(model.UserInfo.UID)
		if err != nil {
			return serializer.JsonByte(http.StatusInternalServerError, err.Error(), nil, err.Error())
		}

		return  serializer.JsonByte(http.StatusOK, "success", count, "")
	case "forward/create":
		// 转发
		prop, err := service.CreateForward(model.UserInfo.UID)
		if err != nil {
			return serializer.JsonByte(http.StatusInternalServerError, err.Error(), nil, err.Error())
		}

		return  serializer.JsonByte(http.StatusOK, "success", prop, "")
	case "backpack/list":
		backpacks := service.GetBackpackList(model.UserInfo.UID)

		return  serializer.JsonByte(http.StatusOK, "success", backpacks, "")
	case "backpack/prop/use":
		params := model.NewPropUse()
		err = json.Unmarshal([]byte(m.Content), params)
		if err != nil {
			return serializer.JsonByte(http.StatusInternalServerError, "数据解析失败", nil, "")
		}

		result := service.BackpackPropUse(params)

		return  serializer.JsonByte(constbase.ENHANCER_RESULT, "success", result, "")
	case "store/list":
		stores := service.GetStoreList()

		return  serializer.JsonByte(http.StatusOK, "success", stores, "")
	case "store/change":
		params := service.NewStoreChange()

		err = json.Unmarshal([]byte(m.Content), params)
		if err != nil {
			return serializer.JsonByte(http.StatusInternalServerError, "数据解析失败", nil, "")
		}

		b := service.StoreChange(params.SID)
		if b == false {
			return serializer.JsonByte(http.StatusInternalServerError, "兑换失败", nil, "")
		}

		return  serializer.JsonByte(http.StatusOK, "success", nil, "")
	}

	return serializer.JsonByte(http.StatusOK, "success", nil, "")
}
