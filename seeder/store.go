package seeder

import (
	"giligili/model"
	"log"
)

func StoreRun() {
	stores := []model.Store{}

	err := model.DB.Find(&stores).Error
	if err != nil {
		log.Println(err.Error())
		return
	}

	if len(stores) > 0 {
		return
	}

	
}