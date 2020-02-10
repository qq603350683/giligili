package controller

import (
	"giligili/model"
	"giligili/serializer"
	"giligili/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 获取投稿视频列表
func GetListVideo(c *gin.Context) {
	offset, _ := util.ToInt(c.Param("offset"))
	limit, _ := util.ToInt(c.Param("limit"))

	res := model.GetListVideo(offset, limit)

	serializer.Response(c, res)
}

// 投稿视频
func CreateVideo(c *gin.Context) {
	m := model.CreateVideoSerivce{}
	if err := c.ShouldBind(&m); err != nil {
		serializer.Exit(c, http.StatusOK, "保存视频失败", err.Error())
		return
	}

	res := m.CreateVideo()

	serializer.Response(c, res)
}

// 修改视频
func UpdateVideo(c *gin.Context) {
	m := model.UpdateVideoService{}
	if err := c.ShouldBind(&m); err != nil {
		serializer.Exit(c, http.StatusOK, "修改视频失败", err.Error())
		return
	}

	v_id, err := util.ToInt(c.Param("v_id"))
	if err != nil {
		serializer.Exit(c, http.StatusOK, "参数错误", err.Error())
		return
	}


	res := m.UpdateVideo(v_id)

	serializer.Response(c, res)
}

// 删除视频
func DeleteVideo(c *gin.Context) {
	v_id, err := util.ToInt(c.Param("v_id"))
	if err != nil {
		serializer.Exit(c, http.StatusOK, "参数错误", err.Error())
		return
	}

	res := model.DeleteVideo(v_id)

	serializer.Response(c, res)
}

func GetVideoInfo(c *gin.Context) {
	v_id, err := util.ToInt(c.Param("v_id"))
	if err != nil {
		serializer.Exit(c, http.StatusOK, "参数错误", err.Error())
		return
	}

	res := model.GetVideoInfo(v_id)

	serializer.Response(c, res)
}