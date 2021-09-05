package model

import (
	"encoding/json"
	"log"
	"time"
)

type UserPlan struct {
	UpID int `json:"up_id" gorm:"column:up_id; type:int(10) unsigned auto_increment; not null; primary_key"`
	UID int `json:"u_id" gorm:"column:u_id; type:int(10) unsigned; not null; default:0; index:idx_u_id; comment:'用户ID 来自 users 表的 u_id'"`
	IsPutOn int8 `json:"is_put_on" gorm:"column:is_put_on;type:tinyint(1) unsigned;not null;default:0;comment:'是否已装备 0 - 否 1 - 是'"`
	DetailJson string `json:"-" gorm:"column:detail; type: text; not null; comment:'飞机详情json格式'"`
	Detail *PlanDetail `json:"detail" gorm:"-" comment:"飞机详情json格式"`
	DelAt time.Time `json:"-" gorm:"type:datetime;not null;default:'1000-01-01 00:00:00'; comment:'删除时间'"`
	CreatedAt time.Time `json:"-" gorm:"column:created_at; type:datetime; not null; comment:'创建时间'"`
}

func NewUserPlan() *UserPlan {
	user_plan := new(UserPlan)
	user_plan.CreatedAt = time.Now()

	return user_plan
}

func GetUserPlanInfo(up_id int) *UserPlan {
	if up_id == 0 {
		return nil
	}

	plan := &UserPlan{}

	err := DB.Where("up_id = ?", up_id).First(plan).Error
	if err != nil {
		return nil
	}

	if IsDel(plan.DelAt) {
		return nil
	}

	// 完善飞机信息
	err = json.Unmarshal([]byte(plan.DetailJson), &plan.Detail)
	if err != nil {
		log.Printf("飞机信息 json 解析字段 detail 失败 up_id: %d, 失败详情: %s", up_id, err.Error())
		return nil
	}

	return plan
}


func GetUserPlans(u_id int) []UserPlan {
	if u_id == 0 {
		return nil
	}

	var plans []UserPlan

	err := DB.Where("u_id = ? AND del_at = ?", u_id, DelAtDefaultTime).Order("is_put_on desc, up_id desc").Find(&plans).Error
	if err != nil {
		log.Println(err.Error())
		return nil
	}

	if len(plans) == 0 {
		return nil
	}

	for index, plan := range(plans) {
		err = json.Unmarshal([]byte(plan.DetailJson), &plan.Detail)
		if err != nil {
			log.Printf("飞机信息 json 解析字段 detail 失败 up_id: %d, 失败详情: %s", plan.UpID, err.Error())
			return nil
		}

		plans[index] = plan
	}

	return plans
}