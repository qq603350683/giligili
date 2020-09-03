package service

import (
	"giligili/constbase"
	"giligili/model"
	"log"
)

type LevelPassParams struct {
	LID int `json:"l_id"`
	IsSuccess int8 `json:"is_success"`
}

func NewLevelPassParams() *LevelPassParams {
	return &LevelPassParams{}
}

// 创建通关记录
func UserPassLevelCreate(l_id int, is_success int8) *model.UserPassLevel {
	level := model.GetLevelInfo(l_id)
	if level == nil {
		return nil
	}

	b := false
	gold := 0
	diamond := 0

	db := model.DB.Begin()

	if is_success == constbase.YES {
		gold = 100
		diamond = 5

		b = model.UserInfo.GoldAndDiamondIncr(gold, diamond)
		if b == false {
			db.Rollback()
			return nil
		}
		log.Printf("玩家（%d）通关关卡（%d）成功", model.UserInfo.UID, l_id)
	} else {
		log.Printf("玩家（%d）通关关卡（%d）失败", model.UserInfo.UID, l_id)
	}

	user_pass_level := model.UserInfo.PassLevel(l_id, is_success, gold, diamond)
	if user_pass_level == nil {
		db.Rollback()
		return nil
	}

	db.Commit()

	return user_pass_level
}
