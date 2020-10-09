package socket

import (
	"encoding/json"
	"fmt"
	"giligili/model"
	"giligili/serializer"
	"giligili/util"
	"log"
	"net/http"
)

func UserPlanUpgradeBulletLevel(params Params) {
	id := 0
	up_id := 0
	index := -1
	bullet := model.Bullet{}

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

	for i, b := range(user_plan.Detail.Bullets) {
		if b.BID == id {
			index = i
			bullet = b
			break
		}
	}

	if bullet.BID == 0 {
		log.Println("请选择需要强化的 bullet: b_id 不能为0")
		SendMessage(model.UserInfo.UID, serializer.JsonByte(http.StatusInternalServerError, "升级失败", nil, "bullet: b_id 不能为 0"))
		return
	}

	if index == -1 {
		log.Println("请选择需要强化的 bullet: index 不能为 -1")
		SendMessage(model.UserInfo.UID, serializer.JsonByte(http.StatusInternalServerError, "升级失败", nil, "bullet: index 不能为 -1"))
		return
	}

	// 已达到最高强化级别
	if user_plan.Detail.Bullets[index].MaxLevel <= user_plan.Detail.Bullets[index].Level {
		SendMessage(model.UserInfo.UID, serializer.JsonByte(http.StatusInternalServerError, "已达最大强化级别", nil, "bullet: index 不能为 -1"))
		return
	}

	// 强化成功
	bullet.Level += 1

	// 消耗的金币
	consume_gold := bullet.Level * 100

	if model.UserInfo.Gold < consume_gold {
		SendMessage(model.UserInfo.UID, serializer.JsonByte(http.StatusInternalServerError, fmt.Sprintf("金币不足 %d", consume_gold), nil, "金币不足"))
		return
	}

	db := model.DBBegin()

	user_plan.Detail.Bullets[index] = bullet

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

	SendMessage(model.UserInfo.UID, serializer.JsonByte(http.StatusOK, "升级成功", model.UserInfo.Plan, ""))
}
