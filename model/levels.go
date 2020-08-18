package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Level struct {
	LID int `json:"l_id" gorm:"column:l_id; type: int(10); unsigned auto_increment; primary_key"`
	Title string `json:"title" gorm:"column:title; type: char(10); not null; default:''; comment:'标题'"`
	Detail string `json:"detail" gorm:"column:detail; type: text; not null; comment:'关卡详情'"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at; type:datetime; not null; comment:'创建时间'"`
	DelAt time.Time `json:"-" gorm:"type:datetime;not null;default:'1000-01-01 00:00:00'; comment:'删除时间'"`
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

func GetLevelByID(l_id int) (*Level, error) {
	level := NewLevel()

	if l_id == 0 {
		return level, nil
	}

	err := DB.Where("l_id = ?", l_id, ).First(level).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return level, err
		}
		return level, err
	}

	if IsDel(level.DelAt) {
		return level, nil
	}

	return level, nil
}

func GetLevelList() (*[]Level, error) {
	levels := &[]Level{}

	err := DB.Find(levels).Error
	if err != nil {
		return levels, err
	}

	return levels, nil
}