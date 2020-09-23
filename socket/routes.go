package socket

import (
	"encoding/json"
	"giligili/model"
	"giligili/serializer"
	"giligili/util"
	"net/http"
)

var Routes map[string]Router

type Router struct {
	f HandlerFunc
}

type GetParams map[string]interface{}

type Params map[string]string

type HandlerFunc func(params Params)

type GetMessage struct {
	Case string `json:"case"`
	Params GetParams `json:"content"`
}

func HandleGetMessage(msg []byte) {
	// 解析json格式 {"case": "sign_in/count", "content": "{\"l_id\": 1}"}
	message := GetMessage{}
	err := json.Unmarshal(msg, &message)
	if err != nil {
		SendMessage(model.UserInfo.UID, serializer.JsonByte(http.StatusInternalServerError, "系统繁忙", nil, "JSON解析失败"))
		return
	}

	if _, ok := Routes[message.Case]; !ok {
		SendMessage(model.UserInfo.UID, serializer.JsonByte(http.StatusInternalServerError, "系统繁忙", nil, "不存在此路由"))
		return
	}

	params := make(Params)

	for index, value := range(message.Params) {
		params[index] = util.InterfaceToString(value)
	}

	model.UserInfo = model.GetUserInfo(model.UserInfo.UID)

	Routes[message.Case].f(params)
}

func AddRoute(url string, f HandlerFunc) {
	if len(Routes) == 0 {
		Routes = make(map[string]Router)
	}

	Routes[url] = Router{
		f:      f,
	}
}
