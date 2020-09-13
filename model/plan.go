package model

import "time"

type Plan struct {
	PID int `json:"p_id" gorm:"column:p_id; type:int(10) unsigned auto_increment; not null; primary_key"`
	DetailJson string `json:"-" gorm:"column:detail; type: text; not null; comment:'飞机详情json格式'"`
	Detail *PlanDetail `json:"detail" gorm:"-" comment:"飞机详情json格式"`
	DelAt time.Time `json:"-" gorm:"type:datetime;not null;default:'1000-01-01 00:00:00'; comment:'删除时间'"`
	CreatedAt time.Time `json:"-" gorm:"column:created_at; type:datetime; not null; comment:'创建时间'"`
}

// 飞机详情 包括普通子弹 技能等信息
type PlanDetail struct {
	Name string `json:"-" comment:"名字"`
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
	MaxLevel int `json:"max_level" comment:"最大攻击力"`
	Rate int `json:"rate" comment:"射频"`
	MaxRate int `json:"max_rate" comment:"最大射频"`
	Speed int `json:"speed" comment:"速度"`
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
	MaxLevel int `json:"max_level" comment:"最大攻击力"`
	Rate int `json:"rate" comment:"频率"`
	MaxRate int `json:"max_rate" comment:"最大射频"`
	Speed int `json:"speed" comment:"速度"`
	MaxSpeed int `json:"max_speed" comment:"最大速度"`
	MaxHeight int `json:"height" comment:"技能最长长度"`
	Texture string `json:"texture" comment:"技能图片"`
}
