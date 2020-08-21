package service

import (
	"giligili/constbase"
	"giligili/model"
	"giligili/serializer"
	"giligili/util"
	"github.com/jinzhu/gorm"
	"net/http"
	"time"
)

type UpdateVideoService struct {
	Title string `form:"title" json:"title" binding:"required,min=2,max=30"`
	Info string `form:"info" json:"info" binding:"max=30"`
}

// 修改视频
func (service *UpdateVideoService) UpdateVideo(v_id uint) serializer.JsonResponse {
	video := &model.Video {
		VId: v_id,
		DelAt: time.Now(),
	}

	err := video.GetInfoById()
	if err != nil {
		return serializer.Json(http.StatusInternalServerError, "修改失败~", nil, err.Error())
	}

	if model.IsDel(video.DelAt) {
		return serializer.Json(http.StatusNotFound, constbase.NotFound, nil, "")
	}

	err = model.DB.Select("v_id").Where("v_id = ?", v_id).Where("del_at = ?", model.DelAtDefault).First(&video).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			video.BuildInfoCache()
			return serializer.Json(http.StatusNotFound, constbase.NotFound, nil, err.Error())
		}
		return serializer.Json(http.StatusInternalServerError, "视频修改失败~", nil, err.Error())
	}

	video.Title = util.XssFilter(service.Title)
	video.Info = util.XssFilter(service.Info)

	err = model.DB.Save(&video).Error
	if err != nil {
		return serializer.Json(http.StatusInternalServerError, "视频修改失败~", nil, err.Error())
	}

	go video.BuildInfoCache()

	info := serializer.BuildVideo(video)

	return serializer.Json(http.StatusOK, "视频修改成功", info, "")
}
