package socket

import (
	"encoding/json"
	"fmt"
	"giligili/constbase"
	"giligili/model"
	"giligili/serializer"
	"giligili/util"
	"log"
	"net/http"
)

func UserPlanUpgradeBulletsLevel(params Params) {
	up_id := 0
	level := 0

	if _, ok := params["up_id"]; !ok {
		SendMessage(model.UserInfo.UID, serializer.JsonByte(http.StatusInternalServerError, "参数异常", nil, "up_id 为 0"))
		return
	}

	up_id = util.StringToInt(params["up_id"])

	user_plan := model.GetUserPlanInfo(up_id)
	if user_plan == nil {
		SendMessage(model.UserInfo.UID, serializer.JsonByte(http.StatusInternalServerError, "升级失败", nil, fmt.Sprintf("当前up_id(%d)不存在", up_id)))
		return
	}

	for i, b := range(user_plan.Detail.Bullets) {
		if level == 0 {
			level = b.Level

			if b.Level >= b.MaxLevel {
				SendMessage(model.UserInfo.UID, serializer.JsonByte(http.StatusInternalServerError, "火力已达最大等级", nil, ""))
				return
			}
		}

		b.Level += 1

		user_plan.Detail.Bullets[i] = b
	}

	// 消耗的金币
	consume_gold := level * 100

	if model.UserInfo.Gold < consume_gold {
		SendMessage(model.UserInfo.UID, serializer.JsonByte(http.StatusInternalServerError, fmt.Sprintf("金币不足 %d", consume_gold), nil, "金币不足"))
		return
	}

	db := model.DBBegin()

	str, err := json.Marshal(user_plan.Detail)
	if err != nil {
		db.Rollback()
		log.Println(err.Error())
		SendMessage(model.UserInfo.UID, serializer.JsonByte(http.StatusInternalServerError, "升级失败", nil, err.Error()))
		return
	}

	user_plan.DetailJson = string(str)

	err = db.Save(user_plan).Error
	if err != nil {
		db.Rollback()
		log.Println(err.Error())
		SendMessage(model.UserInfo.UID, serializer.JsonByte(http.StatusInternalServerError, "升级失败", nil, err.Error()))
		return
	}

	model.UserInfo.GoldAndDiamondUpdate(-consume_gold, 0)

	db.Commit()

	if model.UserInfo.UpID == up_id {
		model.UserInfo.Plan = user_plan
	}

	user_plan_upgrade := NewUserPlanUpgrade()
	user_plan_upgrade.Gold = model.UserInfo.Gold
	user_plan_upgrade.Diamond = model.UserInfo.Diamond
	user_plan_upgrade.UserPlan = user_plan

	SendMessage(model.UserInfo.UID, serializer.JsonByte(constbase.USER_PLAN_UPGRADE_SUCCESS, "升级成功", user_plan_upgrade, ""))
}