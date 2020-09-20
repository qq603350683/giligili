package socket

import (
	"encoding/json"
	"giligili/model"
	"giligili/serializer"
	"net/http"
)

var Routes map[string]Router

type Router struct {
	f HandlerFunc
}

type GetParams map[string]interface{}

type HandlerFunc func(params GetParams)

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

	Routes[message.Case].f(message.Params)
}

func AddRoute(url string, f HandlerFunc) {
	if len(Routes) == 0 {
		Routes = make(map[string]Router)
	}

	Routes[url] = Router{
		f:      f,
	}
}
