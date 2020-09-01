package service

import (
	"giligili/constbase"
	"giligili/model"
	"log"
	"time"
)

func BackpackPropUse(params *model.PropUse) *model.PropUseResult {
	b := false

	result := model.NewPropUseResult()
	result.EnhancerlResult = constbase.ENHANCER_FAIL

	if params.PID == 0 {
		return result
	}

	// 道具详情
	prop := model.GetPropInfo(params.PID)
	if prop == nil {
		return result
	}

	// 获取背包的道具
	backpack := model.GetMyBackpackInfo(prop.PID)
	if backpack == nil {
		return result
	}

	backpack.PropDetail = prop

	db := model.DB.Begin()

	switch prop.Type {
	case constbase.PROP_TYPE_BULLET_ENHANCER:
		b = backpack.UseBulletEnhancer(params.UpID, params.ID)
	case constbase.PROP_TYPE_BULLET_SPEED_ENHANCER:
		b = backpack.UseBulletSpeedEnhancer(params.UpID, params.ID)
	case constbase.PROP_TYPE_SKILL_ENHANCER:
	case constbase.PROP_TYPE_SKILL_SPEED_ENHANCER:
	}

	if b == false {
		db.Rollback()
		return result
	} else {
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
