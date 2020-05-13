package service

import (
	"fmt"
	"giligili/model"
	"giligili/test"
	"net/http"
	"reflect"
	"testing"
)

var CreateTitle string
var CreateInfo string

// 测试视频流程
func testFlow() {
	CreateTitle = "我是一个兵~"
	CreateInfo = "爱国爱人民"
}

func testGetEmptyList(t *testing.T) {
	resp := GetListVideo(0, 20)
	if resp.Status != http.StatusOK {
		t.Errorf("testGetEmptyList fail: %v %v", resp.Status, resp.Message)
	}

	if reflect.TypeOf(resp.Data) != test.EmptyListType {
		t.Errorf("testGetEmptyList fail: %v %v", resp.Status, "列表数据不为空")
	}
}

func testGetList(t *testing.T) {
	resp := GetListVideo(0, 20)
	if resp.Status != http.StatusOK {
		t.Errorf("testGetList fail: %v %v", resp.Status, resp.Message)
	}

	if reflect.TypeOf(resp.Data) == test.EmptyListType {
		t.Errorf("testGetList fail: %v %v", resp.Status, "列表数据为空")
	}

	fmt.Println(resp.Data)
}

func testGetInfo(t *testing.T) {

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

func TestAll(t *testing.T) {
	t.Run("testGetEmptyList", testGetEmptyList)
	t.Run("testGetInfo", testGetInfo)
	t.Run("testCreate", testCreate)
	t.Run("testGetList", testGetList)
}

func TestMain(m *testing.M) {
	fmt.Println("begin testing...")

	test.Init()

	tables := []string{"videos"}

	model.ClearTables(tables)

	m.Run()
}
