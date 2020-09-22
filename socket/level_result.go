package socket

import (
	"giligili/constbase"
	"giligili/model"
	"giligili/serializer"
	"giligili/util"
	"log"
	"net/http"
)

//type LevelResult struct {
//	LID int `json:"l_id"`
//	IsSuccess int8 `json:"is_success"`
//}
//
//func NewLevelResult() *LevelResult {
//	return new(LevelResult)
//}

// 闯关成功/失败结果
func GetLevelResult(params Params) {
	l_id := 0
	is_success := constbase.NO

	if _, ok := params["l_id"]; !ok {
		SendMessage(model.UserInfo.UID, serializer.JsonByte(http.StatusInternalServerError, "参数异常", nil, "参数 l_id 为 不存在"))
		return
	}

	if _, ok := params["is_success"]; !ok {
		SendMessage(model.UserInfo.UID, serializer.JsonByte(http.StatusInternalServerError, "参数异常", nil, "参数 is_success 不存在"))
		return
	}

	l_id = util.StringToInt(params["l_id"])
	is_success = util.StringToInt(params["is_success"])

	boolean := false
	gold := 0
	diamond := 0

	if (l_id > model.UserInfo.LID + 1) {
		SendMessage(model.UserInfo.UID, serializer.JsonByte(http.StatusInternalServerError, "通关关卡异常", nil, "通关关卡异常"))
		return
	}

	db := model.DBBegin()

	defer model.CancelDB()

	// 通关成功领取的奖励
	if is_success == constbase.YES {
		count, err := model.CountTodayPass(model.UserInfo.UID, l_id)
		if err != nil {
			SendMessage(model.UserInfo.UID, serializer.JsonByte(http.StatusInternalServerError, "系统繁忙", nil, "CountTodayPass 返回异常"))
			return
		}

		if count < 5 {
			gold = 100
			diamond = 1

			boolean = model.UserInfo.GetPassLevelPrize(l_id, gold, diamond)
			if boolean == false {
				db.Rollback()
				SendMessage(model.UserInfo.UID, serializer.JsonByte(http.StatusInternalServerError, "系统繁忙", nil, "奖励领取失败"))
				return
			}
		}

		log.Printf("玩家（%d）通关关卡（%d）成功", model.UserInfo.UID, l_id)
	} else {
		log.Printf("玩家（%d）通关关卡（%d）失败", model.UserInfo.UID, l_id)
	}

	user_pass_level := model.NewUserPassLevel()

	user_pass_level.UID = model.UserInfo.UID
	user_pass_level.LID = l_id
	user_pass_level.IsSucess = int8(is_success)
	user_pass_level.Gold = gold
	user_pass_level.Diamond = diamond

	err := db.Create(user_pass_level).Error
	if err != nil {
		log.Println(err.Error())
		db.Rollback()
		SendMessage(model.UserInfo.UID, serializer.JsonByte(http.StatusInternalServerError, "系统繁忙", nil, err.Error()))
		return
	}

	db.Commit()

	SendMessage(model.UserInfo.UID, serializer.JsonByte(constbase.LEVEL_PASS_PRIZE, "success", user_pass_level, ""))
}