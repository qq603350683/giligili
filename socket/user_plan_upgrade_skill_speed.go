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

func UserPlanUpgradeSkillSpeed(params Params) {
	id := 0
	up_id := 0
	index := -1
	skill := model.Skill{}

	if _, ok := params["up_id"]; !ok {
		SendMessage(model.UserInfo.UID, serializer.JsonByte(http.StatusInternalServerError, "参数异常", nil, "up_id 为 0"))
		return
	}

	if _, ok := params["id"]; !ok {
		SendMessage(model.UserInfo.UID, serializer.JsonByte(http.StatusInternalServerError, "参数异常", nil, "id 为 0"))
		return
	}

	id = util.StringToInt(params["id"])
	up_id = util.StringToInt(params["up_id"])

	user_plan := model.GetUserPlanInfo(up_id)
	if user_plan == nil {
		SendMessage(model.UserInfo.UID, serializer.JsonByte(http.StatusInternalServerError, "升级失败", nil, fmt.Sprintf("当前up_id(%d)不存在", up_id)))
		return
	}

	for i, s := range(user_plan.Detail.Skills) {
		if s.SID == id {
			index = i
			skill = s
			break
		}
	}

	if skill.SID == 0 {
		log.Println("请选择需要强化的 skill: s_id 不能为0")
		SendMessage(model.UserInfo.UID, serializer.JsonByte(http.StatusInternalServerError, "升级失败", nil, "skill: s_id 不能为 0"))
		return
	}

	if index == -1 {
		log.Println("请选择需要强化的 skill: index 不能为 -1")
		SendMessage(model.UserInfo.UID, serializer.JsonByte(http.StatusInternalServerError, "升级失败", nil, "skill: index 不能为 -1"))
		return
	}

	// 已达到最高强化级别
	if user_plan.Detail.Skills[index].MaxSpeed <= user_plan.Detail.Skills[index].Speed {
		SendMessage(model.UserInfo.UID, serializer.JsonByte(http.StatusInternalServerError, "已达最大强化级别", nil, "skill: index 不能为 -1"))
		return
	}

	// 强化成功
	skill.Speed += 1

	// 消耗的金币
	consume_gold := skill.Speed * 10

	if model.UserInfo.Gold < consume_gold {
		SendMessage(model.UserInfo.UID, serializer.JsonByte(http.StatusInternalServerError, fmt.Sprintf("金币不足 %d", consume_gold), nil, "金币不足"))
		return
	}

	db := model.DBBegin()

	user_plan.Detail.Skills[index] = skill

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
