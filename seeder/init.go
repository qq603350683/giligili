package seeder

import (
	"giligili/model"
	"log"
	"time"
)

// 初始化数据库
func Run() {
	//var virus = {
	//hp: this.getRandomNumber(10, 100),
	//	speed: Math.ceil(Math.random() * 20),
	//		x: Math.ceil(Math.random() * 600),
	//		y: 0,
	//		w: 150,
	//		h: 150,
	//		texture: "virus/virus.png"
	//}

	levels := []model.Level{}

	err := model.DB.Find(&levels).Error
	if err != nil {
		log.Println(err.Error())
		return
	}

	if len(levels) > 0 {
		return
	}

	levels = append(levels, model.Level{
		LID:        1,
		Level:      1,
		Title:      "第一关",
		Background: "bg/bg0.jpg",
		Virus:      `[
{"time": 10, "hp": 5, "speed": 1, "x": 100, "y": 0, "w": 150, "h": 150, "texture": "virus/virus.png"}, 
{"time": 60, "hp": 6, "speed": 1, "x": 200, "y": 0, "w": 150, "h": 150, "texture": "virus/virus.png"},
{"time": 120, "hp": 10, "speed": 1, "x": 300, "y": 0, "w": 150, "h": 150, "texture": "virus/virus.png"},
{"time": 180, "hp": 12, "speed": 1, "x": 400, "y": 0, "w": 150, "h": 150, "texture": "virus/virus.png"}
]`,
		CreatedAt:  time.Now(),
		DelAt:      model.DelAtDefaultTime,
	})

	levels = append(levels, model.Level{
		LID:        2,
		Level:      2,
		Title:      "第二关",
		Background: "bg/bg1.jpg",
		Virus:      `[
{"time": 10, "hp": 7, "speed": 1, "x": 100, "y": 0, "w": 150, "h": 150, "texture": "virus/virus.png"}, 
{"time": 60, "hp": 8, "speed": 1, "x": 200, "y": 0, "w": 150, "h": 150, "texture": "virus/virus.png"},
{"time": 120, "hp": 11, "speed": 1, "x": 300, "y": 0, "w": 150, "h": 150, "texture": "virus/virus.png"},
{"time": 180, "hp": 15, "speed": 1, "x": 400, "y": 0, "w": 150, "h": 150, "texture": "virus/virus.png"}
]`,
		CreatedAt:  time.Now(),
		DelAt:      model.DelAtDefaultTime,
	})

	for _, level := range(levels) {
		err = model.DB.Create(level).Error
		if err != nil {
			log.Println(err.Error())
		}
	}


}
