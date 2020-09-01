package service

import (
	"giligili/model"
)

type StoreChangeParams struct {
	SID int `json:"s_id"`
}

func NewStoreChange() *StoreChangeParams {
	return &StoreChangeParams{}
}

func StoreChange(s_id int) bool {
	store := model.GetSroteInfo(s_id)
	if store == nil {
		return false
	}

	db :=model.DB.Begin()

	b := store.Change()

	if b == false {
		db.Rollback()
		//model.DB.Rollback()
		return false
	}

	db.Commit()
	//model.DB.CommonDB()

	return b
}
