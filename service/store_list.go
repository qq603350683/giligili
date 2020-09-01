package service

import "giligili/model"

func GetStoreList() []model.Store {
	stores := model.GetStoreList()

	return stores
}
