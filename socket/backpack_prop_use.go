package socket

import (
	"giligili/constbase"
	"giligili/model"
	"giligili/serializer"
	"giligili/util"
	"log"
	"net/http"
	"time"
)

type PropUseResult struct {
	PID int `json:"p_id" comment:"道具ID"`
	Type string `json:"type" comment:"道具类型"`
	Quantity int `json:"quantity" comment:"数量"`
	EnhancerlResult string `json:"enhancer_result" comment:"强化结果"`
	Plan *model.UserPlan `json:"plan" comment:"飞机详情"`
}

func NewPropUseResult() *PropUseResult {
	return new(PropUseResult)
}

func BackpackPropUse(params Params) {
	p_id  := 0
	up_id := 0
	b_id  := 0
	s_id  := 0

	if _, ok := params["p_id"]; !ok {
		SendMessage(model.UserInfo.UID, serializer.JsonByte(http.StatusInternalServerError, "参数异常", nil, "BackpackPropUse p_id 参数获取失败"))
		return
	} else {
		p_id = util.StringToInt(params["p_id"])
	}

	if _, ok := params["up_id"]; ok {
		up_id = util.StringToInt(params["up_id"])
	}

	if _, ok := params["b_id"]; ok {
		b_id = util.StringToInt(params["b_id"])
	}

	if _, ok := params["s_id"]; ok {
		s_id = util.StringToInt(params["s_id"])
	}

	if p_id == 0 {
		SendMessage(model.UserInfo.UID, serializer.JsonByte(http.StatusInternalServerError, "参数异常", nil, "BackpackPropUse p_id 不能为 0"))
		return
	}

	//log.Printf("BackpackPropUse 参数 p_id: %d, up_id: %d, b_id: %d, s_id: %d", p_id, up_id, b_id, s_id)

	boolean := false
	quantity := 0
	enhancer_result := false

	result := NewPropUseResult()
	result.PID = p_id
	result.EnhancerlResult = ""

	// 道具详情
	prop := model.GetPropInfo(p_id)
	if prop == nil {
		SendMessage(model.UserInfo.UID, serializer.JsonByte(http.StatusInternalServerError, "系统繁忙", nil, "socket.BackpackPropUse prop is nil"))
		return
	}

	result.Type = prop.Type

	// 获取背包的道具
	backpack := model.GetMyBackpackInfo(prop.PID)
	if backpack == nil {
		SendMessage(model.UserInfo.UID, serializer.JsonByte(http.StatusInternalServerError, "系统繁忙", nil, "socket.BackpackPropUse 背包没有当前道具"))
		return
	}

	backpack.PropDetail = prop

	db := model.DBBegin()

	defer model.CancelDB()

	switch prop.Type {
	case constbase.PROP_TYPE_GOLD:
		quantity, boolean = backpack.OpenGoldPack()
		result.Quantity = quantity
	case constbase.PROP_TYPE_DIAMOND:
		quantity, boolean = backpack.OpenDiamondPack()
		result.Quantity = quantity
	case constbase.PROP_TYPE_BULLET_ENHANCER:
		enhancer_result, boolean = backpack.UseBulletEnhancer(up_id, b_id)
	case constbase.PROP_TYPE_BULLET_SPEED_ENHANCER:
		enhancer_result, boolean = backpack.UseBulletSpeedEnhancer(up_id, b_id)
	case constbase.PROP_TYPE_BULLET_RATE_ENHANCER:
		enhancer_result, boolean = backpack.UseBulletRateEnhancer(up_id, b_id)
	case constbase.PROP_TYPE_SKILL_ENHANCER:
		enhancer_result, boolean = backpack.UseSkillEnhancer(up_id, s_id)
	case constbase.PROP_TYPE_SKILL_SPEED_ENHANCER:
		enhancer_result, boolean = backpack.UseSkillSpeedEnhancer(up_id, s_id)
	case constbase.PROP_TYPE_SKILL_RATE_ENHANCER:
		enhancer_result, boolean = backpack.UseSkillRateEnhancer(up_id, s_id)
	}

	if boolean == false {
		db.Rollback()
		SendMessage(model.UserInfo.UID, serializer.JsonByte(http.StatusInternalServerError, "系统繁忙", nil, "socket.BackpackPropUse 异常003"))
		return
	}

	if enhancer_result == true {
		result.EnhancerlResult = constbase.ENHANCER_SUCCESS
		result.Plan = model.GetUserPlanInfo(up_id)
	} else {
		result.EnhancerlResult = constbase.ENHANCER_FAIL
	}
	log.Println(backpack)
	res := db.Model(backpack).Update(map[string]interface{}{
		"use_at": time.Now(),
		"is_use": constbase.YES,
	})
	if res.RowsAffected == 0 {
		db.Rollback()
		SendMessage(model.UserInfo.UID, serializer.JsonByte(http.StatusInternalServerError, "系统繁忙", nil, "socket.BackpackPropUse 异常004"))
		return
	}

	db.Commit()

	SendMessage(model.UserInfo.UID, serializer.JsonByte(constbase.ENHANCER_RESULT, "success", result, ""))

	result.Type = prop.Type
}
