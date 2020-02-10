package model

import (
	"giligili/cache"
	"giligili/message"
	"giligili/serializer"
	"giligili/util"
	"github.com/jinzhu/gorm"
	"net/http"
	"time"
)

type Video struct {
	VId       int    `json:"v_id" gorm:"type:int(10) unsigned auto_increment;primary_key;"`
	Title     string `json:"title" gorm:"type:varchar(30);not null"`
	Info      string `json:"info" gorm:"type:varchar(30);not null"`
	Browse    int `json:"browse" gorm:"type:int(10) unsigned;not null"`
	Love      int `json:"love" gorm:"type:int(10) unsigned;not null"`
	CreatedAt time.Time `json:"created_at" gorm:"type:datetime;not null"`
	UpdatedAt time.Time `json:"update_at" gorm:"type:datetime;not null"`
	DelAt time.Time `json:"-" gorm:"type:datetime;not null;default:'1000-01-01 00:00:00'"`
}

type CreateVideoSerivce struct {
	Title string `form:"title" json:"title" binding:"required,min=2,max=30"`
	Info string `form:"info" json:"info" binding:"max=30"`
}

type UpdateVideoService struct {
	Title string `form:"title" json:"title" binding:"required,min=2,max=30"`
	Info string `form:"info" json:"info" binding:"max=30"`
}

// 获取投稿视频列表
func GetListVideo(offset int, limit int) serializer.JsonResponse {
	if limit == 0 {
		return serializer.Json(http.StatusOK, message.OK, nil, "")
	}

	var videos []Video

	err := DB.Select("v_id, title, info, created_at").Where("del_at = ?", DelAtDefault).Order("v_id desc").Offset(offset).Limit(limit).Find(&videos).Error
	if err != nil {
		return serializer.Json(http.StatusInternalServerError, "数据查询失败~", nil, err.Error())
	}

	list := BuildListVideo(videos)

	return serializer.Json(http.StatusOK, message.OK, list, "")
}

// 投稿视频
func (service *CreateVideoSerivce) CreateVideo() serializer.JsonResponse {
	video := Video {
		Title: util.XssFilter(service.Title),
		Info: util.XssFilter(service.Info),
	}

	err := DB.Create(&video).Error
	if err != nil {
		return serializer.Json(http.StatusInternalServerError, "视频保存失败", nil, err.Error())
	}

	info := BuildVideo(&video)

	return serializer.Json(http.StatusOK, "视频保存成功", info, "")
}

// 修改视频
func (service *UpdateVideoService) UpdateVideo(v_id int) serializer.JsonResponse {
	var video Video

	err := DB.Select("v_id").Where("v_id = ?", v_id).Where("del_at = ?", DelAtDefault).First(&video).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return serializer.Json(http.StatusNotFound, "找不到当前记录", nil, err.Error())
		}
		return serializer.Json(http.StatusInternalServerError, "视频修改失败~", nil, err.Error())
	}

	video.Title = util.XssFilter(service.Title)
	video.Info = util.XssFilter(service.Info)

	err = DB.Save(&video).Error
	if err != nil {
		return serializer.Json(http.StatusInternalServerError, "视频修改失败~", nil, err.Error())
	}

	info := BuildVideo(&video)

	return serializer.Json(http.StatusOK, "视频修改成功", info, "")
}

// 删除视频
func DeleteVideo(v_id int) serializer.JsonResponse {
	var video Video

	err := DB.Select("v_id").Where("v_id = ?", v_id).Where("del_at = ?", DelAtDefault).First(&video).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return serializer.Json(http.StatusNotFound, "找不到当前记录", nil, err.Error())
		}
		return serializer.Json(http.StatusInternalServerError, "视频删除失败~", nil, err.Error())
	}

	video.DelAt = time.Now()

	err = DB.Save(&video).Error
	if err != nil {
		return serializer.Json(http.StatusInternalServerError, "视频删除失败~~", nil, err.Error())
	}

	return serializer.Json(http.StatusOK, "视频删除成功~", nil, "")
}

// 获取视频
func GetVideoInfo(v_id int) serializer.JsonResponse {
	re, _ := cache.GetRedis()
	err2 := re.Set("hello", "Go", 0).Err()
	if err2 != nil {
		panic(err2)
	}

	var video Video

	err := DB.Select("v_id, info, browse, love, created_at").Where("v_id = ?", v_id).Where("del_at = ?", DelAtDefault).First(&video).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return serializer.Json(http.StatusNotFound, message.NotFound, nil, err.Error())
		}
		return serializer.Json(http.StatusInternalServerError, "查找数据失败~", nil, err.Error())
	}

	// 阅读数 +1
	err = DB.Model(&video).UpdateColumn("browse", video.Browse+1).Error
	if err != nil {
		video.Browse += 1
	}
	
	info := BuildVideo(&video)

	return serializer.Json(http.StatusOK, message.OK, info, "")
}

// 获取视频详情
func BuildVideo(video *Video) map[string]interface{} {
	return map[string]interface{} {
		"v_id": video.VId,
		"title": video.Title,
		"info": video.Info,
		"browse" : video.Browse,
		"love": video.Love,
		"created_at": util.Datetime(video.CreatedAt),
	}
}

// 获取视频详情列表
func BuildListVideo(videos []Video) []map[string]interface{} {
	count := len(videos)
	if count == 0 {
		return nil
	}

	list := make([]map[string]interface{}, 0, count - 1)

	for _, video := range videos {
		list = append(list, map[string]interface{} {
			"v_id": video.VId,
			"title": video.Title,
			"info": video.Info,
			"created_at": util.Datetime(video.CreatedAt),
		})
	}

	return list
}

