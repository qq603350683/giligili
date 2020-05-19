package service

import (
	"encoding/json"
	"fmt"
	"giligili/model"
	"giligili/test"
	"net/http"
	"reflect"
	"testing"
)

var VID uint
var CreateTitle string
var CreateInfo string

var UpdateTitle string
var UpdateInfo string


func testGetEmptyList(t *testing.T) {
	resp := GetListVideo(0, 20)
	if resp.Status != http.StatusOK {
		t.Errorf("testGetEmptyList fail: %v %v", resp.Status, resp.Message)
	}

	marshal, _ := json.Marshal(resp.Data)
	if string(marshal) != "[]" {
		t.Errorf("testGetEmptyInfo fail: %v", "数据不是空列表")
	}
}

func testGetList(t *testing.T) {
	resp := GetListVideo(0, 20)
	if resp.Status != http.StatusOK {
		t.Errorf("testGetList fail: %v %v", resp.Status, resp.Message)
	}

	marshal, _ := json.Marshal(resp.Data)
	if string(marshal) == "[]" {
		t.Errorf("testGetList fail: %v", "列表数据为空")
	}
}

func testGetEmptyInfo(t *testing.T) {
	resp := GetVideoInfo(VID)
	if resp.Status != http.StatusNotFound {
		t.Errorf("testGetEmptyInfo fail: %v %v", resp.Status, resp.Message)
	}

	marshal, err := json.Marshal(resp.Data)

	if string(marshal) != "{}" {
		t.Errorf("testGetEmptyInfo fail: %v", err.Error())
	}
}

func testUpdateEmptyInfo(t *testing.T) {
	m := UpdateVideoService{
		Title: UpdateTitle,
		Info: UpdateInfo,
	}

	resp := m.UpdateVideo(VID)
	if resp.Status != http.StatusNotFound {
		t.Errorf("testUpdateEmptyInfo fail: %v %v", resp.Status, resp.Message)
	}

	if reflect.TypeOf(resp.Data) != test.EmptyListType {
		t.Errorf("testUpdateEmptyInfo fail: %v %v", resp.Status, "数据不为空")
	}
}

func testGetInfo(t *testing.T) {
	resp := GetVideoInfo(VID)
	if resp.Status == http.StatusNotFound {
		t.Errorf("testGetInfo fail: %v %v", resp.Status, resp.Message)
	}
}

func testCreate(t *testing.T) {
	m := CreateVideoSerivce {
		Title: CreateTitle,
		Info: CreateInfo,
	}

	resp := m.CreateVideo()

	if resp.Status != http.StatusOK {
		t.Errorf("testCreate fail: %v %v", resp.Status, resp.Message)
	}
}

func testUpdate(t *testing.T) {
	m := UpdateVideoService{
		Title: UpdateTitle,
		Info: UpdateInfo,
	}

	resp := m.UpdateVideo(VID)

	if resp.Status != http.StatusOK {
		t.Errorf("testUpdate fail: %v %v", resp.Status, resp.Message)
	}

	marshal, _ := json.Marshal(resp.Data)
	if string(marshal) == "{}" {
		t.Errorf("testUpdate fail: %v %v", resp.Status, "更新失败")
	}

	video := model.NewVideo(VID)

	err := video.GetInfoById()
	if err != nil {
		t.Errorf("testUpdate fail: %v %v", resp.Status, "查找数据失败")
	}

	if video.Title != UpdateTitle {
		t.Errorf("testUpdate fail: %v", "标题更新失败")
	}

	if video.Info != UpdateInfo {
		t.Errorf("testUpdate fail: %v", "描述更新失败")
	}
}

func TestAll(t *testing.T) {
	// 顺序不能乱
	t.Run("testGetEmptyList", testGetEmptyList)
	t.Run("testGetEmptyInfo", testGetEmptyInfo)
	t.Run("testUpdateEmptyInfo", testUpdateEmptyInfo)
	t.Run("testCreate", testCreate)
	t.Run("testGetInfo", testGetInfo)
	t.Run("testUpdate", testUpdate)
	t.Run("testGetList", testGetList)
}

func TestMain(m *testing.M) {
	fmt.Println("begin testing...")

	test.Init()

	tables := []string{"videos"}

	model.ClearTables(tables)

	VID = 1

	CreateTitle = "我是一个兵~"
	CreateInfo = "爱国爱人民"

	UpdateTitle = "修改 - 我是一个兵~"
	UpdateInfo = "修改 - 爱国爱人民~"

	m.Run()
}
