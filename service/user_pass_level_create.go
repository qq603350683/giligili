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
		count, err := model.CountTodayPass(model.UserInfo.UID, l_id)
		if err != nil {
			return nil
		}

		if count < 5 {
			gold = 100
			diamond = 1

			b = model.UserInfo.GetPassLevelPrize(l_id, gold, diamond)
			if b == false {
				db.Rollback()
				return nil
			}
		}

		log.Printf("玩家（%d）通关关卡（%d）成功", model.UserInfo.UID, l_id)
	} else {
		log.Printf("玩家（%d）通关关卡（%d）失败", model.UserInfo.UID, l_id)
	}

	user_pass_level := model.NewUserPassLevel()

	user_pass_level.UID = model.UserInfo.UID
	user_pass_level.LID = l_id
	user_pass_level.IsSucess = is_success
	user_pass_level.Gold = gold
	user_pass_level.Diamond = diamond

	err := model.DB.Create(user_pass_level).Error
	if err != nil {
		log.Println(err.Error())
		return nil
	}

	if user_pass_level == nil {
		db.Rollback()
		return nil
	}

	db.Commit()

	return user_pass_level
}
