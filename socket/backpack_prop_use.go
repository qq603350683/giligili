package socket

import (
	"giligili/model"
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
	//p_id  := 0
	//up_id := 0
	//b_id  := 0
	//s_id  := 0
	//
	//if _, ok := params["p_id"]; ok {
	//	SendMessage(model.UserInfo.UID, serializer.JsonByte(http.StatusInternalServerError, "参数异常", nil, "p_id 为 0"))
	//	return
	//}
	//
	//if _, ok := params["up_id"]; ok {
	//	up_id := util.StringToInt(params["up_id"])
	//}
	//
	//if _, ok := params["b_id"]; ok {
	//	b_id := util.StringToInt(params["b_id"])
	//}
	//
	//if _, ok := params["s_id"]; ok {
	//	s_id := util.StringToInt(params["s_id"])
	//}
	//
	//boolean := false
	//quantity := 0
	//enhancer_result := false
	//
	//result := NewPropUseResult()
	//result.PID = p_id
	//result.EnhancerlResult = ""
	//
	//// 道具详情
	//prop := model.GetPropInfo(p_id)
	//if prop == nil {
	//	SendMessage(model.UserInfo.UID, serializer.JsonByte(http.StatusInternalServerError, "系统繁忙", nil, "prop is nil"))
	//	return
	//}
	//
	//// 获取背包的道具
	//backpack := model.GetMyBackpackInfo(prop.PID)
	//if backpack == nil {
	//	SendMessage(model.UserInfo.UID, serializer.JsonByte(http.StatusInternalServerError, "系统繁忙", nil, "prop is nil"))
	//	return
	//}
	//
	//backpack.PropDetail = prop
	//
	//db := model.DB.Begin()
	//
	//switch prop.Type {
	//case constbase.PROP_TYPE_GOLD:
	//	quantity, boolean = backpack.UseDB(db).OpenGoldPack()
	//	result.Quantity = quantity
	//case constbase.PROP_TYPE_DIAMOND:
	//	quantity, boolean = backpack.UseDB(db).OpenDiamondPack()
	//	result.Quantity = quantity
	//case constbase.PROP_TYPE_BULLET_ENHANCER:
	//	enhancer_result, boolean = backpack.UseBulletEnhancer(up_id, b_id)
	//case constbase.PROP_TYPE_BULLET_SPEED_ENHANCER:
	//	enhancer_result, boolean = backpack.UseBulletSpeedEnhancer(up_id, b_id)
	//case constbase.PROP_TYPE_BULLET_RATE_ENHANCER:
	//	enhancer_result, boolean = backpack.UseBulletRateEnhancer(up_id, b_id)
	//case constbase.PROP_TYPE_SKILL_ENHANCER:
	//	enhancer_result, boolean = backpack.UseSkillEnhancer(up_id, s_id)
	//case constbase.PROP_TYPE_SKILL_SPEED_ENHANCER:
	//	enhancer_result, boolean = backpack.UseSkillSpeedEnhancer(up_id, s_id)
	//case constbase.PROP_TYPE_SKILL_RATE_ENHANCER:
	//	enhancer_result, boolean = backpack.UseSkillRateEnhancer(up_id, s_id)
	//}
	//
	//// SendMessage(model.UserInfo.UID, serializer.JsonByte(constbase.ENHANCER_RESULT, "success", result, ""))
	//
	//result.Type = prop.Type
}
