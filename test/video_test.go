package test

import (
	"giligili/model"
	"testing"
)

func TestA(t *testing.T) {
	var v_id uint

	v_id = 1

	video := model.NewVideo(v_id)

	err := video.GetInfoById()
	if err != nil {
		t.Errorf("video get info by id fail: %v", err.Error())
	}

	if video.VId != v_id {
		t.Errorf("video.VID is not equal v_id")
	}

	
}
