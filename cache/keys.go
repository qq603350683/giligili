package cache

import (
	"fmt"
	"giligili/util"
	"time"
)

// 视频详情
func VideoInfoKey(v_id uint) string {
	if v_id == 0 {
		return ""
	}

	return fmt.Sprintf("video:info:%s", util.ToString(int(v_id)))
}

// 视频列表
func VideoListKey() string {
	// ZRevRange video:list 0 20 withscores
	return fmt.Sprintf("video:list")
}

// 视频浏览数Zset
func VideoBrowseListKey() string {
	return "video:browse:list"
}

// 记录每个用户当天是否转发
func UserTodayForwardListKey() string {
	return "user_today_forward_" + time.Now().Format("20060102")
}
