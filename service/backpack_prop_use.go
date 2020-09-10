package service

import (
	"giligili/constbase"
	"giligili/model"
	"log"
	"time"
)

func BackpackPropUse(params *model.PropUse) *model.PropUseResult {
	b := false
	quantity := 0
	enhancer_result := false

	result := model.NewPropUseResult()
	result.PID = params.PID
	result.EnhancerlResult = constbase.ENHANCER_FAIL

	if params.PID == 0 {
		return result
	}

	// 道具详情
	prop := model.GetPropInfo(params.PID)
	if prop == nil {
		return result
	}

	result.Type = prop.Type

	// 获取背包的道具
	backpack := model.GetMyBackpackInfo(prop.PID)
	if backpack == nil {
		return result
	}

	backpack.PropDetail = prop

	db := model.DB.Begin()

	switch prop.Type {
	case constbase.PROP_TYPE_GOLD:
		quantity, b = backpack.OpenGoldPack()
		result.Quantity = quantity
	case constbase.PROP_TYPE_DIAMOND:
		quantity, b = backpack.OpenDiamondPack()
		result.Quantity = quantity
	case constbase.PROP_TYPE_BULLET_ENHANCER:
		enhancer_result, b = backpack.UseBulletEnhancer(params.UpID, params.BID)
	case constbase.PROP_TYPE_BULLET_SPEED_ENHANCER:
		enhancer_result, b = backpack.UseBulletSpeedEnhancer(params.UpID, params.BID)
	case constbase.PROP_TYPE_SKILL_ENHANCER:
		enhancer_result, b = backpack.UseSkillEnhancer(params.UpID, params.SID)
	case constbase.PROP_TYPE_SKILL_SPEED_ENHANCER:
		enhancer_result, b = backpack.UseSkillSpeedEnhancer(params.UpID, params.SID)
	}

	if b == false {
		db.Rollback()
		return result
	}

	if enhancer_result == true {
		result.EnhancerlResult = constbase.ENHANCER_SUCCESS
	}

	backpack.IsUse = constbase.YES
	backpack.UseAt = time.Now()

	err := model.DB.Save(backpack).Error
	if err != nil {
		log.Println(err.Error())
		db.Rollback()
		return result
	}

	db.Commit()

	return result
}
