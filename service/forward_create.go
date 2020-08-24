package service

import (
	"errors"
	"giligili/constbase"
	"giligili/model"
)

// 转发
func CreateForward(u_id int) (*model.PropShow, error) {
	if u_id == 0 {
		return nil, errors.New("用户ID不能为0")
	}

	prop := model.GetForwardPrize()

	if prop == nil {
		return nil, nil
	}

	// 获取今天是否已转发
	b := model.UserInfo.TodayIsForward()
	if b == true {
		return nil, nil
	}

	b = model.UserInfo.TodayForward()
	if b == false {
		return nil, errors.New(constbase.SystemBusy)
	}

	// 把礼品加入到背包
	prop.AddToBackpack()

	prop_show := &model.PropShow{
		Prop:     prop,
		Quantity: 1,
	}

	return prop_show, nil
}