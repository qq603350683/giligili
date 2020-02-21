package model

import (
	"encoding/json"
	"errors"
	"fmt"
	"giligili/cache"
	"giligili/util"
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	"time"
)

type Video struct {
	VId       uint    `json:"v_id" gorm:"type:int(10) unsigned auto_increment;primary_key;"`
	Title     string `json:"title" gorm:"type:varchar(30);not null"`
	Info      string `json:"info" gorm:"type:varchar(30);not null"`
	Browse    uint `json:"browse" gorm:"type:int(10) unsigned;not null"`
	Love      uint `json:"love" gorm:"type:int(10) unsigned;not null"`
	CreatedAt time.Time `json:"created_at" gorm:"type:ToDatetime;not null"`
	UpdatedAt time.Time `json:"update_at" gorm:"type:ToDatetime;not null"`
	DelAt time.Time `json:"del_at" gorm:"type:ToDatetime;not null;default:'1000-01-01 00:00:00'"`
}

// 获取默认状态下的Video结构体
func NewVideo(v_id uint) *Video {
	video := &Video {
		VId:   v_id,
		DelAt: time.Now(),
	}

	return video
}

// 根据ID获取详情
func (video *Video) GetInfoById() error {
	// 获取缓存的数据
	err := video.GetInfoCache()
	if err != nil {
		return err
	}

	err = DB.Where("v_id = ?", video.VId).First(video).Error

	// 写入缓存
	video.BuildInfoCache()
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil
		}
		return err
	}

	if IsDel(video.DelAt) {
		return nil
	}
	// 写入浏览数缓存
	video.SetVideoBrowse()

	return nil
}

// 获取视频的浏览量
func GetVideoBrowse(v_id uint) uint {
	client, err := cache.RedisCache.Get()
	if err != nil {
		panic("Redis: 连接池获取Redis失败")
	}

	member := util.ToString(int(v_id))

	result, err := client.ZScore(cache.VideoBrowseListKey(), member).Result()
	if err != nil {
		if err != cache.RedisNil {
			return 0
		}

		// 这里执行SetVideoBrowse为了防止redis击穿
		video := NewVideo(v_id)
		return video.SetVideoBrowse()
	}

	res := uint(result)

	return res
}

// 设置浏览量
func (video *Video) SetVideoBrowse() uint {
	if IsDel(video.DelAt) {
		return 0
	}

	client, err := cache.RedisCache.Get()
	if err != nil {
		panic("Redis: 连接池获取Redis失败")
	}

	member := util.ToString(int(video.VId))

	z := redis.Z {
		Score: float64(video.Browse),
		Member: member,
	}

	err = client.ZAdd(cache.VideoBrowseListKey(), z).Err()
	if err != nil {
		panic(err)
	}

	return video.Browse
}

// 浏览量 +1
func (video *Video) IncrBrowse() error {
	if IsDel(video.DelAt) {
		return errors.New("Video: " + util.ToString(int(video.VId)) + " 已是删除数据")
	}

	video.Browse += 1
	fmt.Println("1")
	// 更新缓存
	client, err := cache.RedisCache.Get()
	if err != nil {
		panic(err)
	}
	fmt.Println("2")
	member := util.ToString(int(video.VId))
	result, err := client.ZIncrBy(cache.VideoBrowseListKey(), 1.01, member).Result()
	fmt.Println(result)
	if err != nil {
		panic(err)
	}
	fmt.Println("3")
	// 每100个播放量更新数据库, 缓存一次
	if video.Browse % 100 == 0 {
		err = DB.Model(video).UpdateColumn("browse", video.Browse).Error
		if err != nil {
			panic(err)
		}

		video.BuildInfoCache()
	}

	return nil
}

// 获取缓存详情
func (video *Video) GetInfoCache() error {
	client, err := cache.RedisCache.Get()
	if err != nil {
		panic("Redis: 连接池获取Redis失败")
	}

	result, err := client.Get(cache.VideoInfoKey(video.VId)).Result()
	if err != nil {
		if err == cache.RedisNil {
			return nil
		}
		return err
	}

	err = json.Unmarshal([]byte(result), &video)
	if err != nil {
		return err
	}

	return nil
}

// 建立详情缓存
func (video *Video) BuildInfoCache() bool {
	client, err := cache.RedisCache.Get()
	if err != nil {
		panic("Redis: 连接池获取Redis失败")
	}

	str := util.GetEmptyJsonByte()

	if video.VId > 0 {
		str, err = json.Marshal(video)
		if err != nil {
			panic("Redis: 转化Video结构体失败")
		}
	}

	err = client.Set(cache.VideoInfoKey(video.VId), str, 0).Err()
	if err != nil {
		panic("Redis: 设置" + cache.VideoInfoKey(video.VId) + "失败")
	}

	return true
}

// 删除Redis缓存
func (video *Video) DelInfoCache() bool {
	client, err := cache.RedisCache.Get()
	if err != nil {
		panic("Redis: 连接池获取Redis失败")
	}

	err = client.Del(cache.VideoInfoKey(video.VId)).Err()
	if err != nil {
		panic("Redis: 删除" + cache.VideoInfoKey(video.VId) + "失败")
	}

	return true
}

