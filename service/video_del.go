package service

import (
	"giligili/model"
	"giligili/serializer"
	"github.com/jinzhu/gorm"
	"net/http"
	"time"
)

// 删除视频
func DelVideo(v_id uint) serializer.JsonResponse {
	var video model.Video

	err := model.DB.Select("v_id").Where("v_id = ?", v_id).Where("del_at = ?", model.DelAtDefault).First(&video).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return serializer.Json(http.StatusNotFound, "找不到当前记录", nil, err.Error())
		}
		return serializer.Json(http.StatusInternalServerError, "视频删除失败~", nil, err.Error())
	}

	video.DelAt = time.Now()

	err = model.DB.Save(&video).Error
	if err != nil {
		return serializer.Json(http.StatusInternalServerError, "视频删除失败~~", nil, err.Error())
	}

	return serializer.Json(http.StatusOK, "视频删除成功~", nil, "")
}
