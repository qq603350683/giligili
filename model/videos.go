package model

import (
	"encoding/json"
	"giligili/cache"
	"giligili/util"
	"github.com/jinzhu/gorm"
	"time"
)

type Video struct {
	VId       uint    `json:"v_id" gorm:"type:int(10) unsigned auto_increment;primary_key;"`
	Title     string `json:"title" gorm:"type:varchar(30);not null"`
	Info      string `json:"info" gorm:"type:varchar(30);not null"`
	Browse    int `json:"browse" gorm:"type:int(10) unsigned;not null"`
	Love      int `json:"love" gorm:"type:int(10) unsigned;not null"`
	CreatedAt time.Time `json:"created_at" gorm:"type:ToDatetime;not null"`
	UpdatedAt time.Time `json:"update_at" gorm:"type:ToDatetime;not null"`
	DelAt time.Time `json:"del_at" gorm:"type:ToDatetime;not null;default:'1000-01-01 00:00:00'"`
}

func GetInfoById(v_id uint) *Video {
	// 获取缓存的数据

	video := Video {
		VId: v_id,
		DelAt: time.Now(),
	}

	err := DB.Where("v_id = ?", v_id).First(&video).Error
	video.BuildInfoCache()
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return &video
		}
		return &video
	}

	if util.ToDatetime(video.DelAt) == DelAtDefault {
		return &video
	}

	return &video
}

func (video *Video) BuildInfoCache() bool {
	client, err := cache.RedisCache.Get()
	if err != nil {
		panic("Redis: 连接池获取Redis失败")
	}

	defer cache.RedisCache.Put(client)

	str := util.GetEmptyJsonByte()

	if video.VId > 0 {
		str, err = json.Marshal(video)
		if err != nil {
			panic("Redis: 转化Video结构体失败")
		}
	}

	err = client.Set(cache.VideoInfoKey(video.VId), str, 0).Err()
	if err != nil {
		panic("Redis: 设置" + cache.VideoBrowseKey(video.VId) + "失败")
	}

	return true
}

// 删除Redis缓存
func (video *Video) DelInfoCache() bool {
	client, err := cache.RedisCache.Get()
	if err != nil {
		panic("Redis: 连接池获取Redis失败")
	}

	defer cache.RedisCache.Put(client)

	err = client.Del(cache.VideoInfoKey(video.VId)).Err()
	if err != nil {
		panic("Redis: 删除" + cache.VideoBrowseKey(video.VId) + "失败")
	}

	return true
}

