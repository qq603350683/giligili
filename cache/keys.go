package cache

import (
	"fmt"
	"strconv"
)

// 视频详情
func VideoInfoKey(v_id uint) string {
	if v_id == 0 {
		return ""
	}

	return fmt.Sprintf("video:info:%s", strconv.Itoa(int(v_id)))
}

// 视频浏览数Zset
func VideoBrowseListKey() string {
	return "video:browse:list"
}
