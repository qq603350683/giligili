package socket

import (
	"giligili/constbase"
	"giligili/model"
	"giligili/serializer"
)

type StoreListResult struct {
	Type string `json:"type"`
	Stores []model.Store `json:"stores"`
}

func NewStoreListResult() *StoreListResult {
	return new(StoreListResult)
}

func GetStoreList(params Params) {
	result := NewStoreListResult()

	result.Type  = params["type"]
	result.Stores = model.GetStoreList(params["type"])



	SendMessage(model.UserInfo.UID, serializer.JsonByte(constbase.STORE_LIST, "success", result, ""))
}
