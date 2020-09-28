package socket

import (
	"giligili/constbase"
	"giligili/model"
	"giligili/serializer"
	"log"
)

func ShareCreate(params Params) {
	boolean := false
	position := constbase.SHARE_POSITION_MENU   // 先设置默认为首页菜单分享

	if _, ok := params["position"]; ok {
		position = params["position"]
	}

	boolean = model.InUserSharePrizePositions(position)
	if boolean == false {
		log.Printf("分享位置 %s 不存在", position)

		// 这里重新再复制一次是为了防止前端传递过来的值覆盖了13行代码position初始值
		position = constbase.SHARE_POSITION_MENU
	}

	boolean = model.UserInfo.TodayIsShare(position)
	if boolean == true {
		SendMessage(model.UserInfo.UID, serializer.JsonByte(constbase.IGNORE, "今天分享奖励已经领取", nil, ""))
		return
	}

	// 奖品
	p_id := 2
	quantity := 200

	db := model.DBBegin()

	defer model.CancelDB()

	user_share_prize := model.NewUserSharePrize()

	user_share_prize.UID = model.UserInfo.UID
	user_share_prize.PID = p_id
	user_share_prize.Quantity = quantity
	user_share_prize.Position = position

	if err := db.Create(user_share_prize).Error; err != nil {
		db.Rollback()
		SendMessage(model.UserInfo.UID, serializer.JsonByte(constbase.IGNORE, "分享失败", nil, "分享失败001"))
		return
	}

	user_share_prize.PropDetail = model.GetPropInfo(user_share_prize.PID)

	if boolean = user_share_prize.PropDetail.AddToUserDiamond(user_share_prize.Quantity); boolean == false {
		db.Rollback()
		SendMessage(model.UserInfo.UID, serializer.JsonByte(constbase.IGNORE, "分享失败", nil, "分享失败002"))
		return
	}

	db.Commit()

	//SendMessage(model.UserInfo.UID, serializer.JsonByte(, "分享失败", nil, "分享失败002"))
}
