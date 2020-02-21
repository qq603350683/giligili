package serializer

import (
	"giligili/model"
	"giligili/util"
)

const (
	VId = "v_id"
	Title = "title"
	Info = "info"
	Browse = "browse"
	Love = "love"
	CreateAt = "create_at"
)

// 获取视频详情
func BuildVideo(video *model.Video) map[string]interface{} {
	return map[string]interface{} {
		VId: video.VId,
		Title: video.Title,
		Info: video.Info,
		Browse : model.GetVideoBrowse(video.VId),
		Love: video.Love,
		CreateAt: util.ToDatetime(video.CreatedAt),
	}
}

// 获取视频详情列表
func BuildVideoList(videos []model.Video) []map[string]interface{} {
	count := len(videos)
	if count == 0 {
		return nil
	}

	list := make([]map[string]interface{}, 0, count - 1)

	for _, video := range videos {
		list = append(list, map[string]interface{} {
			VId: video.VId,
			Title: video.Title,
			Info: video.Info,
			Browse: model.GetVideoBrowse(video.VId),
			Love: video.Love,
			CreateAt: util.ToDatetime(video.CreatedAt),
		})
	}

	return list
}