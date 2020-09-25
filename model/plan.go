package model

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
	"log"
	"time"
)

type Plan struct {
	PID int `json:"p_id" gorm:"column:p_id; type:int(10) unsigned auto_increment; not null; primary_key"`
	DetailJson string `json:"-" gorm:"column:detail; type: text; not null; comment:'飞机详情json格式'"`
	Detail *PlanDetail `json:"detail" gorm:"-" comment:"飞机详情json格式"`
	DelAt time.Time `json:"-" gorm:"type:datetime;not null;default:'1000-01-01 00:00:00'; comment:'删除时间'"`
	CreatedAt time.Time `json:"-" gorm:"column:created_at; type:datetime; not null; comment:'创建时间'"`
}

// 飞机详情 包括普通子弹 技能等信息
type PlanDetail struct {
	Name string `json:"name" comment:"名字"`
	Width int `json:"w" comment:"宽度"`
	Height int `json:"h" comment:"高度"`
	Texture string `json:"texture" comment:"飞机图片"`
	Bullets []Bullet `json:"bullets" comment:"子弹集合"`
	Skills []Skill `json:"skills" comment:"技能集合"`
}

// 普通子弹
type Bullet struct {
	BID int `json:"id" comment:"暂时不要"`
	Title string `json:"title" comment:"名字"`
	Width int `json:"w" comment:"宽度"`
	Height int `json:"h" comment:"高度"`
	Position int8 `json:"p" comment:"位置"`
	Angle int8 `json:"a" comment:"角度"`
	Level int `json:"level" comment:"攻击力"`
	BaseLevel int `json:"base_level" comment:"初始攻击力"`
	MaxLevel int `json:"max_level" comment:"最大攻击力"`
	Rate int `json:"rate" comment:"射频"`
	BaseRate int `json:"rate" comment:"初始射频"`
	MaxRate int `json:"max_rate" comment:"最大射频"`
	Speed int `json:"speed" comment:"速度"`
	BaseSpeed int `json:"base_speed" comment:"初始速度"`
	MaxSpeed int `json:"max_speed" comment:"最大速度"`
	Texture string `json:"texture" comment:"子弹图片"`
}

// 被动技能
type Skill struct {
	SID int `json:"id" comment:"暂时不要"`
	Title string `json:"title" comment:"名字"`
	Width int `json:"w" comment:"宽度"`
	Height int `json:"h" comment:"高度"`
	Position int8 `json:"p" comment:"位置"`
	Angle int8 `json:"a" comment:"角度"`
	Level int `json:"level" comment:"攻击力"`
	BaseLevel int `json:"base_level" comment:"初始攻击力"`
	MaxLevel int `json:"max_level" comment:"最大攻击力"`
	Rate int `json:"rate" comment:"频率"`
	BaseRate int `json:"rate" comment:"初始射频"`
	MaxRate int `json:"max_rate" comment:"最大射频"`
	Speed int `json:"speed" comment:"速度"`
	BaseSpeed int `json:"base_speed" comment:"初始速度"`
	MaxSpeed int `json:"max_speed" comment:"最大速度"`
	MaxHeight int `json:"height" comment:"技能最长长度"`
	Texture string `json:"texture" comment:"技能图片"`
}

func NewPlan() *Plan {
	plan := new(Plan)

	return plan
}

func GetPlanInfo(p_id int) *Plan {
	if p_id == 0 {
		return nil
	}

	plan := NewPlan()

	err := DB.Where("p_id = ?", p_id).First(plan).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Printf("飞机ID(%d)找不到记录: ", p_id)
		} else {
			log.Println(err.Error())
		}
		return nil
	}

	if IsDel(plan.DelAt) {
		log.Printf("飞机ID(%d)已删除: ", p_id)
		return nil
	}

	// 完善飞机信息
	err = json.Unmarshal([]byte(plan.DetailJson), &plan.Detail)
	if err != nil {
		log.Printf("飞机信息 json 解析字段 detail 失败 p_id: %d, 失败详情: %s", p_id, err.Error())
		return nil
	}

	return plan
}

// 添加到我的飞机里面
func (plan *Plan) AddToUserPlan() bool {
	if plan == nil {
		log.Println("model.AddToUserPlan plan is nil")
		return false
	}

	var err error

	user_plan := NewUserPlan()

	//detail_json, err := json.Marshal(&)
	user_plan.UID = UserInfo.UID
	user_plan.DetailJson = plan.DetailJson
	user_plan.CreatedAt = time.Now()

	log.Println(user_plan)

	db := GetDB()

	err = db.Create(user_plan).Error
	if err != nil {
		log.Println(err.Error())
		return false
	}

	return true
}
