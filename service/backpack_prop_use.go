package service

import (
	"giligili/constbase"
	"giligili/model"
)

func BackpackPropUse(params *model.PropUse) *model.PropUseResult {
	result := model.NewPropUseResult()

	if params.PID == 0 {
		return result
	}

	// 道具详情
	prop := model.GetPropInfo(params.PID)
	if prop == nil {
		return result
	}

	// 获取背包的道具
	backpack := model.GetMyBackpackInfo(prop.Type)
	if backpack == nil {
		return result
	}

	backpack.PropDetail = prop

	switch prop.Type {
	case constbase.PROP_TYPE_BULLET_ENHANCER:
		backpack.UseBulletEnhancer(params.ID)
	case constbase.PROP_TYPE_BULLET_SPEED_ENHANCER:
	case constbase.PROP_TYPE_SKILL_ENHANCER:
	case constbase.PROP_TYPE_SKILL_SPEED_ENHANCER:
	}

	return result
}
