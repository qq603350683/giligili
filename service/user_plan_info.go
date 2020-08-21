package service

import (
	"encoding/json"
	"errors"
	"giligili/model"
	"log"
)

func GetUserPlanInfo(up_id int) (*model.UserPlan, error) {
	if up_id == 0 {
		return nil, errors.New("飞机ID不能为0")
	}

	plan := &model.UserPlan{}

	err := model.DB.Where("up_id = ?", up_id).First(plan).Error
	if err != nil {
		return nil, err
	}

	if model.IsDel(plan.DelAt) {
		return nil, errors.New("当前飞机已经删除")
	}

	// 完善飞机信息
	err = json.Unmarshal([]byte(plan.DetailJson), &plan.Detail)
	if err != nil {
		log.Printf("飞机信息 json 解析字段 detail 失败 up_id: %d, 失败详情: %s", up_id, err.Error())
		return nil, errors.New("数据解析错误，请稍后再试")
	}

	return plan, nil
}