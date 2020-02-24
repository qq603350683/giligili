package service

import (
	"giligili/message"
	"giligili/model"
	"giligili/serializer"
	"net/http"
)

// 获取视频
func GetVideoInfo(v_id uint) serializer.JsonResponse {
	video := model.NewVideo(v_id)

	err := video.GetInfoById()
	if err != nil {
		panic(err)
	}

	if model.IsDel(video.DelAt) {
		return serializer.Json(http.StatusNotFound, message.OK, nil, "")
	}

	video.IncrBrowse()

	info := serializer.BuildVideo(video)

	return serializer.Json(http.StatusOK, message.OK, info, "")
}
