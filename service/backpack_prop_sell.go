package service

import (
	"giligili/constbase"
	"giligili/model"
)

type PropSell struct {
	PID int `json:"p_id"`
}

type PropSellResult struct {
	Type string `json:"type"`
	Quantity int `json:"quantity"`
}

func NewPropSell() *PropSell {
	return &PropSell{}
}

func NewPropSellResult() *PropSellResult {
	return &PropSellResult{}
}

// 出售道具
func BackpackPropSell(p_id int) *PropSellResult {
	PropSellResult := NewPropSellResult()

	backpack := model.GetMyBackpackInfo(p_id)
	if backpack == nil {
		return nil
	}

	prop := model.GetPropInfo(backpack.PID)
	if prop == nil {
		return nil
	}

	db := model.DB.Begin()

	b := backpack.Sell()
	if b == false {
		db.Rollback()
		return nil
	}

	if prop.GoldValue > 0 && prop.DiamondValue == 0 {
		PropSellResult.Type = constbase.PROP_TYPE_GOLD
		PropSellResult.Quantity = prop.GoldValue

		b := model.UserInfo.GoldAndDiamondIncr(prop.GoldValue, 0)
		if b == false {
			db.Rollback()
			return nil
		}
	} else if prop.DiamondValue > 0 && prop.GoldValue == 0 {
		PropSellResult.Type = constbase.PROP_TYPE_DIAMOND
		PropSellResult.Quantity = prop.DiamondValue

		b := model.UserInfo.GoldAndDiamondIncr(0, prop.DiamondValue)
		if b == false {
			db.Rollback()
			return nil
		}
	}

	db.Commit()

	return PropSellResult
}
