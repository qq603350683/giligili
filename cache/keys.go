package cache

import (
	"fmt"
	"giligili/util"
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
