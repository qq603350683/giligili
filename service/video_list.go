package service

import (
	"encoding/json"
	"fmt"
	"giligili/cache"
	"giligili/constbase"
	"giligili/model"
	"giligili/serializer"
	"giligili/util"
	"github.com/go-redis/redis"
	"net/http"
)

// 获取投稿视频列表
func GetListVideo(offset uint, limit uint) serializer.JsonResponse {
	// 没有缓存到的视频数组ID
	var videos []model.Video
	var v_ids []uint

	if limit == 0 || offset > 5000 {
		return serializer.Json(http.StatusOK, constbase.OK, nil, "")
	}

	client, err := cache.RedisCache.Get()
	if err != nil {
		panic(err)
	}

	result, err := client.ZRevRangeWithScores(cache.VideoListKey(), int64(offset), int64(offset + limit - 1)).Result()
	if err != nil && err != cache.RedisNil {
		panic(err)
	}

	if len(result) > 0 {
		pipe := client.Pipeline()
		for _, v := range(result) {
			i, err := util.ToUint(v.Member.(string))
			if err != nil {
				panic(err)
			}

			v_ids = append(v_ids, i)
			pipe.Get(cache.VideoInfoKey(i))
		}
		cmders, err := pipe.Exec()
		if err != nil && err != cache.RedisNil {
			panic(err)
		}
		fmt.Println(cmders)
		for _, cmder := range(cmders) {
			i := v_ids[0]
			v_ids = v_ids[1:]

			res, err := cmder.(*redis.StringCmd).Result()
			if err != nil && err != cache.RedisNil {
				panic(err)
			}

			video := *model.NewVideo(i)

			if err == cache.RedisNil {
				video.GetInfoById()
			} else {
				err = json.Unmarshal([]byte(res), &video)
				if err != nil {
					panic(err)
				}
			}

			videos = append(videos, video)
		}

		list := serializer.BuildVideoList(videos)

		return serializer.Json(http.StatusOK, constbase.OK, list, "")
	} else {
		err = model.DB.Where("del_at = ?", model.DelAtDefault).Order("v_id desc").Limit(5000).Find(&videos).Error

		var rediszs []redis.Z

		if err != nil {
			client.ZAdd(cache.VideoListKey(), rediszs...)

			return serializer.Json(http.StatusInternalServerError, "数据查询失败~", nil, err.Error())
		}

		count := len(videos)

		// 空数据处理
		if count == 0 {
			return serializer.Json(http.StatusOK, constbase.EmptyList, model.EmptyList, "")
		}

		for _, video := range(videos) {
			video.BuildInfoCache()

			z := redis.Z {
				Score:  float64(video.VId),
				Member: video.VId,
			}

			rediszs = append(rediszs, z)
		}

		err = client.ZAdd(cache.VideoListKey(), rediszs...).Err()
		if (err != nil) {
			panic(err)
		}

		l := uint(count)
		if l <= limit {
			limit = l
		}

		videos := videos[0:limit]

		list := serializer.BuildVideoList(videos)

		return serializer.Json(http.StatusOK, constbase.OK, list, "")
	}
}
