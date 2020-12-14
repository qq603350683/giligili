package socket

import "giligili/model"

func WatchAdv(params Params) {
	model.CreateUserAdvRecord(model.UserInfo.UID, "观看广告完毕")
}
