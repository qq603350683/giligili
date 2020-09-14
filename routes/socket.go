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

	// 更新model.UserInfo信息
	model.UserInfo = model.GetUserInfo(model.UserInfo.UID)

	switch m.Case {
	case "level/get":
		level_get := model.NewLevelGet()
		err = json.Unmarshal([]byte(m.Content), level_get)
		level := model.GetLevelInfo(level_get.LID)

		return serializer.JsonByte(constbase.LEVEL_INFO, "success", level, "")
	case "level/pass":
		params := service.NewLevelPassParams()
		err = json.Unmarshal([]byte(m.Content), params)
		if err != nil {
			log.Printf("json 错误: %s", m.Content)
			return serializer.JsonByte(http.StatusInternalServerError, "参数错误", nil, "")
		}

		user_pass_level := service.UserPassLevelCreate(params.LID, params.IsSuccess)

		return serializer.JsonByte(constbase.LEVEL_PASS_PRIZE, "success", user_pass_level, "")
	case "user":
		// 我的详情
		user := service.GetUserInfo(model.UserInfo.UID)
		if user == nil {
			return serializer.JsonByte(http.StatusInternalServerError, "用户信息不存在", nil, "")
		}

		return  serializer.JsonByte(constbase.LOGIN_USER_INFO, "success", user, "")
	case "user/plan/change":
		params := service.NewUserPlanChangeParams()
		err = json.Unmarshal([]byte(m.Content), params)
		if err != nil {
			log.Printf("json 错误: %s", m.Content)
			return serializer.JsonByte(http.StatusInternalServerError, "参数错误", nil, "")
		}

		bool := service.UserPlanChange(params.UpID)
		if bool == false {
			return serializer.JsonByte(constbase.USER_PLAN_CHANGE_FAIL, "更换失败", nil, "")
		}

		return serializer.JsonByte(constbase.USER_PLAN_CHANGE_SUCCESS, "更换成功", nil, "")
	case "sign_in/create":
		// 今天签到
		bool, err := service.CreateSignIn(model.UserInfo.UID)
		if bool == false {
			return serializer.JsonByte(constbase.SIGN_IN_FAIL, err.Error(), nil, err.Error())
		}

		return  serializer.JsonByte(constbase.SIGN_IN_SUCCESS, "success", nil, "")
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

		return serializer.JsonByte(constbase.BACKPACK_LIST, "success", backpacks, "")
	case "backpack/prop/use":
		params := model.NewPropUse()
		err = json.Unmarshal([]byte(m.Content), params)
		if err != nil {
			return serializer.JsonByte(http.StatusInternalServerError, "数据解析失败", nil, "")
		}

		result := service.BackpackPropUse(params)

		return  serializer.JsonByte(constbase.ENHANCER_RESULT, "success", result, "")
	case "backpack/prop/sell":
		params := service.NewPropSell()
		err = json.Unmarshal([]byte(m.Content), params)
		if err != nil {
			return serializer.JsonByte(http.StatusInternalServerError, "数据解析失败", nil, "")
		}

		result := service.BackpackPropSell(params.PID)

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
