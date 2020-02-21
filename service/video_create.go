package service

import (
	"giligili/model"
	"giligili/serializer"
	"giligili/util"
	"net/http"
)

type CreateVideoSerivce struct {
	Title string `form:"title" json:"title" binding:"required,min=2,max=30"`
	Info string `form:"info" json:"info" binding:"max=30"`
}

// 投稿视频
func (service *CreateVideoSerivce) CreateVideo() serializer.JsonResponse {
	video := model.Video {
		Title: util.XssFilter(service.Title),
		Info: util.XssFilter(service.Info),
	}

	err := model.DB.Create(&video).Error
	if err != nil {
		return serializer.Json(http.StatusInternalServerError, "视频保存失败", nil, err.Error())
	}

	info := serializer.BuildVideo(&video)
	// 写入缓存
	go video.BuildInfoCache()

	return serializer.Json(http.StatusOK, "视频保存成功", info, "")
}