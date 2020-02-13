package service

import (
	"giligili/message"
	"giligili/model"
	"giligili/serializer"
	"net/http"
)

// 获取投稿视频列表
func GetListVideo(offset uint, limit uint) serializer.JsonResponse {
	if limit == 0 {
		return serializer.Json(http.StatusOK, message.OK, nil, "")
	}

	var videos []model.Video

	err := model.DB.Select("v_id, title, info, created_at").Where("del_at = ?", model.DelAtDefault).Order("v_id desc").Offset(offset).Limit(limit).Find(&videos).Error
	if err != nil {
		return serializer.Json(http.StatusInternalServerError, "数据查询失败~", nil, err.Error())
	}

	list := serializer.BuildVideoList(videos)

	return serializer.Json(http.StatusOK, message.OK, list, "")
}
