package socket

import (
	"giligili/constbase"
	"giligili/model"
	"giligili/serializer"
	"giligili/util"
	"net/http"
)

type PropSellResult struct {
	Type string `json:"type"`
	Quantity int `json:"quantity"`
	UserGold int `json:"user_gold"`
	UserDiamond int `json:"user_diamond"`
}

func NewPropSellResult() *PropSellResult {
	return new(PropSellResult)
}

func BackpackPropSell(params Params) {
	p_id := 0
	boolean := false

	if _, ok := params["p_id"]; !ok {
		SendMessage(model.UserInfo.UID, serializer.JsonByte(http.StatusInternalServerError, "参数异常", nil, "BackpackPropUse p_id 参数获取失败"))
		return
	} else {
		p_id = util.StringToInt(params["p_id"])
	}

	backpack := model.GetMyBackpackInfo(p_id)
	if backpack == nil {
		SendMessage(model.UserInfo.UID, serializer.JsonByte(constbase.BACKPACK_SELL_FAIL, "当前道具已不存在", nil, "socket.BackpackPropSell 道具不存在001"))
		return
	}

	prop := model.GetPropInfo(backpack.PID)
	if prop == nil {
		SendMessage(model.UserInfo.UID, serializer.JsonByte(constbase.BACKPACK_SELL_FAIL, "当前道具已不存在", nil, "socket.BackpackPropSell 道具不存在002"))
		return
	}

	db := model.DBBegin()

	defer model.CancelDB()

	boolean = backpack.Sell()
	if boolean == false {
		db.Rollback()
		SendMessage(model.UserInfo.UID, serializer.JsonByte(constbase.BACKPACK_SELL_FAIL, "出售失败", nil, "socket.BackpackPropSell 道具出售失败"))
		return
	}

	prop_sell_result := NewPropSellResult()

	if prop.GoldValue > 0 && prop.DiamondValue == 0 {
		prop_sell_result.Type = constbase.PROP_TYPE_GOLD
		prop_sell_result.Quantity = prop.GoldValue

		boolean = model.UserInfo.GoldAndDiamondUpdate(prop.GoldValue, 0)
		if boolean == false {
			db.Rollback()
			SendMessage(model.UserInfo.UID, serializer.JsonByte(constbase.BACKPACK_SELL_FAIL, "出售失败", nil, "socket.BackpackPropSell 道具出售失败"))
			return
		}
	} else if prop.DiamondValue > 0 && prop.GoldValue == 0 {
		prop_sell_result.Type = constbase.PROP_TYPE_DIAMOND
		prop_sell_result.Quantity = prop.DiamondValue

		boolean = model.UserInfo.GoldAndDiamondUpdate(0, prop.DiamondValue)
		if boolean == false {
			db.Rollback()
			SendMessage(model.UserInfo.UID, serializer.JsonByte(constbase.BACKPACK_SELL_FAIL, "出售失败", nil, "socket.BackpackPropSell 道具出售失败"))
			return
		}
	}

	db.Commit()

	prop_sell_result.UserGold = model.UserInfo.Gold
	prop_sell_result.UserDiamond = model.UserInfo.Diamond

	SendMessage(model.UserInfo.UID, serializer.JsonByte(constbase.BACKPACK_SELL_SUCCESS, "success", prop_sell_result, ""))
}
