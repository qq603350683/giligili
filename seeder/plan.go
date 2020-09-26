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

	//plans = append(plans, model.Plan{
	//	PID:        1,
	//	DetailJson: `{"name":"初代战神","w":129,"h":120,"texture":"hero/hero2.png","bullets":[{"id":1,"title":"A导弹","w":25,"h":52,"p":3,"a":0,"level":1,"max_level":3,"rate":20,"max_rate":18,"speed":11,"max_speed":15,"texture":"bullet/10.png"}],"skills":[{"id":1,"title":"光速射线","w":50,"h":50,"p":1,"a":0,"level":2,"max_level":4,"rate":180,"max_rate":175,"speed":50,"max_speed":60,"height":9999999,"texture":"bullet/skill1.png"}]}`,
	//	DelAt:      model.DelAtDefaultTime,
	//	CreatedAt:  time.Now(),
	//})

	plans = append(plans, model.Plan{
		PID:        1,
		DetailJson: `{"name":"2020飞行机","w":129,"h":120,"texture":"hero/hero2.png","bullets":[{"id":1,"title":"2020导弹","w":25,"h":52,"p":1,"a":0,"level":1,"max_level":3,"rate":30,"max_rate":25,"speed":10,"max_speed":15,"texture":"bullet/10.png"}, {"id":2,"title":"2020导弹","w":25,"h":52,"p":3,"a":0,"level":1,"max_level":3,"rate":30,"max_rate":25,"speed":10,"max_speed":15,"texture":"bullet/10.png"}, {"id":3,"title":"2020导弹","w":25,"h":52,"p":5,"a":0,"level":1,"max_level":3,"rate":30,"max_rate":25,"speed":10,"max_speed":15,"texture":"bullet/10.png"}],"skills":[{"id":1,"title":"2020射线","w":50,"h":50,"p":1,"a":0,"level":2,"max_level":4,"rate":180,"max_rate":175,"speed":50,"max_speed":60,"height":9999999,"texture":"bullet/skill1.png"}]}`,
		DelAt:      model.DelAtDefaultTime,
		CreatedAt:  time.Now(),
	})


//
//	plans = append(plans, model.Plan{
//		PID:        3,
//		DetailJson: `{
//	"name": "FG-003",
//	"w":129,
//	"h":120,
//	"texture":"hero/hero2.png",
//	"bullets":[
//		{"id":1,"w":20,"h":20,"p":3,"a":0,"level":2,"max_level":8,"rate":50,"max_rate":40,"speed":15,"max_speed":25,"texture":"bullet/10.png"}
//	],
//	"skills":[
//		{"id":1,"w":50,"h":50,"p":3,"a":0,"level":3,"max_level":20,"rate":180,"max_rate":170,"speed":50,"max_speed":50,"height":9999999,"texture":"bullet/skill1.png"}
//	]
//}`,
//		DelAt:      model.DelAtDefaultTime,
//		CreatedAt:  time.Now(),
//	})

	for _, plan := range(plans) {
		plan.DetailJson = strings.Replace(plan.DetailJson, " ", "", -1)
		plan.DetailJson = strings.Replace(plan.DetailJson, "\n", "", -1)

		err = model.DB.Create(plan).Error
		if err != nil {
			log.Println(err.Error())
		}
	}
}
