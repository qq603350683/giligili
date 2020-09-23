package seeder

import (
	"giligili/model"
	"log"
	"strings"
	"time"
)

func PlanRun() {
	plans := []model.Plan{}

	err := model.DB.Find(&plans).Error
	if err != nil {
		log.Println(err.Error())
		return
	}

	if len(plans) > 0 {
		return
	}

	plans = append(plans, model.Plan{
		PID:        1,
		DetailJson: `{
	"name": "FG-001",
	"w":129,
	"h":120,
	"texture":"hero/hero2.png",
	"bullets":[
		{"id":1,"w":20,"h":20,"p":3,"a":0,"level":1,"max_level":3,"rate":180,"max_rate":175,"speed":10,"max_speed":15,"texture":"bullet/10.png"}
	],
	"skills":[
		{"id":1,"w":50,"h":50,"p":3,"a":0,"level":2,"max_level":12,"rate":300,"max_rate":290,"speed":50,"max_speed":45,"height":9999999,"texture":"bullet/skill1.png"}
	]
}`,
		DelAt:      model.DelAtDefaultTime,
		CreatedAt:  time.Now(),
	})

	plans = append(plans, model.Plan{
		PID:        2,
		DetailJson: `{
	"name": "FG-002",
	"w":129,
	"h":120,
	"texture":"hero/hero2.png",
	"bullets":[
		{"id":1,"w":20,"h":20,"p":3,"a":0,"level":2,"max_level":4,"rate":160,"max_rate":155,"speed":10,"max_speed":15,"texture":"bullet/10.png"}
	],
	"skills":[
		{"id":1,"w":50,"h":50,"p":3,"a":0,"level":2,"max_level":12,"rate":250,"max_rate":240,"speed":50,"max_speed":50,"height":9999999,"texture":"bullet/skill1.png"}
	]
}`,
		DelAt:      model.DelAtDefaultTime,
		CreatedAt:  time.Now(),
	})

	plans = append(plans, model.Plan{
		PID:        3,
		DetailJson: `{
	"name": "FG-003",
	"w":129,
	"h":120,
	"texture":"hero/hero2.png",
	"bullets":[
		{"id":1,"w":20,"h":20,"p":3,"a":0,"level":2,"max_level":8,"rate":50,"max_rate":40,"speed":15,"max_speed":25,"texture":"bullet/10.png"}
	],
	"skills":[
		{"id":1,"w":50,"h":50,"p":3,"a":0,"level":3,"max_level":20,"rate":180,"max_rate":170,"speed":50,"max_speed":50,"height":9999999,"texture":"bullet/skill1.png"}
	]
}`,
		DelAt:      model.DelAtDefaultTime,
		CreatedAt:  time.Now(),
	})

	for _, plan := range(plans) {
		plan.DetailJson = strings.Replace(plan.DetailJson, " ", "", -1)
		plan.DetailJson = strings.Replace(plan.DetailJson, "\n", "", -1)

		err = model.DB.Create(plan).Error
		if err != nil {
			log.Println(err.Error())
		}
	}
}
