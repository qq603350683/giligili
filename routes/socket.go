package routes

import (
	"encoding/json"
	"fmt"
	"giligili/model"
	"giligili/serializer"
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
		//l, _ := model.GetLevelList()

		return serializer.JsonByte(http.StatusOK, "success", l, "")
	}

	return serializer.JsonByte(http.StatusOK, "success", nil, "")
}
