package serializer

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
)

type JsonResponse struct {
	Status int `json:"status"`
	Message string `json:"message"`
	Data interface{} `json:"data"`
	Error string `json:"error"`
}

// 在控制器终止用户错误程序
func Exit(c *gin.Context, status int, message string, err string) {
	c.JSON(http.StatusOK, Json(status, message, nil, err))
}

// 在控制器直接返回信息
func Response(c *gin.Context, data JsonResponse) {
	c.JSON(http.StatusOK, data)
}

func JsonByte(status int, message string, data interface{}, err string) []byte {
	//if data == nil {
	//	data = map[string]interface{} {}
	//}

	jr := JsonResponse{
		Status:  status,
		Message: message,
		Data:    data,
		Error:   err,
	}

	b, err2 := json.Marshal(jr)
	if err2 != nil {
		panic(err)
	}

	return b
}

// 返回数据使用
func Json(status int, message string, data interface{}, err string) JsonResponse {
	switch data.(type) {
	case map[string]interface{}:
		//if reflect.ValueOf(data).IsValid() {
		//	data = map[string]interface{} {}
		//}
	case []map[string]interface{}:
		if reflect.ValueOf(data).IsNil() {
			data = make([]string, 0)
		}
	default:
		data = map[string]interface{} {}
	}

	return JsonResponse{
		Status:  status,
		Message: message,
		Data:    data,
		Error:   err,
	}
}