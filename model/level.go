package model

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
	"log"
	"time"
)

type Level struct {
	LID int `json:"l_id" gorm:"column:l_id;type:int(10) unsigned auto_increment;primary_key"`
	Level int `json:"level" gorm:"column:level; type: int(10) unsigned; not null; default:0; comment:'关卡'"`
	Title string `json:"title" gorm:"column:title; type: char(10); not null; default:''; comment:'标题'"`
	Background string `json:"background" gorm:"column:background; type: text; not null; comment:'背景详情 image - 背景图 speed - 速度'"`
	VirusJson string `json:"-" gorm:"column:detail; type: text; not null; comment:'病毒位置'"`
	Virus []Virus `json:"virus" gorm:"-"`
	CreatedAt time.Time `json:"-" gorm:"column:created_at; type:datetime; not null; comment:'创建时间'"`
	DelAt time.Time `json:"-" gorm:"type:datetime;not null;default:'1000-01-01 00:00:00'; comment:'删除时间'"`
}

// 病毒
type Virus struct {
	Time int `json:"time" comment:"病毒出现时间"`
	HP int `json:"hp" comment:"病毒HP"`
	Speed float64 `json:"speed" comment:"病毒移动速度"`
	X int `json:"x" comment:"病毒出现的X位置"`
	Y int `json:"y" comment:"病毒出现的Y位置"`
	Width int `json:"w" comment:"病毒宽度"`
	Height int `json:"h" comment:"高度"`
	Texture string `json:"texture" comment:"背景图"`
}

// 传递参数解析
type LevelGet struct {
	LID int `json:"l_id"`
}

func NewLevel() *Level {
	level := &Level{}

	return level
}

func NewLevelGet() *LevelGet {
	level_get := &LevelGet{}

	return level_get
}

func GetLevelInfo(l_id int) *Level {
	level := NewLevel()

	if l_id == 0 {
		return nil
	}

	err := DB.Where("l_id = ?", l_id,).First(level).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil
		}
		return nil
	}

	if IsDel(level.DelAt) {
		return nil
	}

	if level.VirusJson != "" {
		err = json.Unmarshal([]byte(level.VirusJson), &level.Virus)
		if err != nil {
			log.Printf("Level.VirusJson json 解析失败 l_id: %d, 失败详情: %s", l_id, err.Error())
			return nil
		}
	}

	return level
}

func GetLevelList() (*[]Level, error) {
	levels := &[]Level{}

	err := DB.Find(levels).Error
	if err != nil {
		return levels, err
	}

	return levels, nil
}