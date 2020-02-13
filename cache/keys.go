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

// 视频浏览数
func VideoBrowseKey(v_id uint) string {
	if v_id == 0 {
		return ""
	}

	return fmt.Sprintf("video:browse:%s", strconv.Itoa(int(v_id)))
}
