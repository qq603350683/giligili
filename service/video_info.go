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
	//err := model.DB.Select("v_id, title, info, browse, love, created_at").Where("v_id = ?", v_id).Where("del_at = ?", model.DelAtDefault).First(&video).Error
	//if err != nil {
	//	if err == gorm.ErrRecordNotFound {
	//
	//		return serializer.Json(http.StatusNotFound, message.NotFound, nil, err.Error())
	//	}
	//	return serializer.Json(http.StatusInternalServerError, "查找数据失败~", nil, err.Error())
	//}
	//
	//// 阅读数 +1
	//err = model.DB.Model(&video).UpdateColumn("browse", video.Browse + 1).Error
	//if err != nil {
	//	video.Browse += 1
	//}
	//
	//info := serializer.BuildVideo(&video)
	//
	//// 写入缓存
	//
	//return serializer.Json(http.StatusOK, message.OK, info, "")
}
