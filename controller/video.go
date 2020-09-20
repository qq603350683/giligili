package controller

import (
	"fmt"
	"giligili/serializer"
	"giligili/service"
	"giligili/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 获取投稿视频列表
func GetListVideo(c *gin.Context) {
	offset := GetOffset(c.Param("offset"))
	limit := GetLimit(c.Param("limit"))

	res := service.GetListVideo(offset, limit)

	serializer.Response(c, res)
}

// 投稿视频
func CreateVideo(c *gin.Context) {
	m := service.CreateVideoSerivce{}
	if err := c.ShouldBind(&m); err != nil {
		serializer.Exit(c, http.StatusOK, "保存视频失败", err.Error())
		return
	}

	res := m.CreateVideo()
	fmt.Println(res)
	serializer.Response(c, res)
}

// 修改视频
func UpdateVideo(c *gin.Context) {
	m := service.UpdateVideoService{}
	if err := c.ShouldBind(&m); err != nil {
		serializer.Exit(c, http.StatusOK, "修改视频失败", err.Error())
		return
	}

	v_id := util.ToUint(c.Param("v_id"))
	//if err != nil {
	//	serializer.Exit(c, http.StatusOK, "参数错误", err.Error())
	//	return
	//}

	res := m.UpdateVideo(v_id)

	serializer.Response(c, res)
}

// 删除视频
func DelVideo(c *gin.Context) {
	v_id := util.ToUint(c.Param("v_id"))
	//if err != nil {
	//	serializer.Exit(c, http.StatusOK, "参数错误", err.Error())
	//	return
	//}

	res := service.DelVideo(v_id)

	serializer.Response(c, res)
}

// 获取详情
func GetVideoInfo(c *gin.Context) {
	v_id := util.ToUint(c.Param("v_id"))
	//if err != nil {
	//	serializer.Exit(c, http.StatusOK, "参数错误", err.Error())
	//	return
	//}

	res := service.GetVideoInfo(v_id)

	serializer.Response(c, res)
}