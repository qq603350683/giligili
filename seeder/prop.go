package seeder

import (
	"giligili/model"
	"log"
	"time"
)

func PropRun() {
	props := []model.Prop{}

	err := model.DB.Find(&props).Error
	if err != nil {
		log.Println(err.Error())
		return
	}

	if len(props) > 0 {
		return
	}

	props = append(props, model.Prop{
		PID:       1,
		Type:      "gold",
		Image:     "prop/gold.png",
		Title:     "金币大礼包",
		Remark:    "金币大礼包",
		CreatedAt: time.Now(),
	})

	props = append(props, model.Prop{
		PID:       2,
		Type:      "diamond",
		Image:     "prop/diamond.png",
		Title:     "钻石大礼包",
		Remark:    "钻石大礼包",
		CreatedAt: time.Now(),
	})

	props = append(props, model.Prop{
		PID:       3,
		Type:      "stone_enhancer_material",
		Image:     "prop/stone_enhancer_material.png",
		Title:     "石岩碳",
		Remark:    "强化子弹、技能攻击力时所需要的石头",
		CreatedAt: time.Now(),
	})

	props = append(props, model.Prop{
		PID:       4,
		Type:      "stone_speed_enhancer_material",
		Image:     "prop/stone_speed_enhancer_material.png",
		Title:     "竹碳",
		Remark:    "强化子弹、技能速度时所需要的石头",
		CreatedAt: time.Now(),
	})

	props = append(props, model.Prop{
		PID:       1001,
		Type:      "bullet_enhancer",
		Image:     "prop/bullet_enhancer.png",
		Title:     "B攻击强化器",
		Remark:    "B攻击强化器",
		CreatedAt: time.Now(),
	})

	props = append(props, model.Prop{
		PID:       1002,
		Type:      "bullet_speed_enhancer",
		Image:     "prop/bullet_speed_enhancer.png",
		Title:     "B速度强化器",
		Remark:    "B速度强化器",
		CreatedAt: time.Now(),
	})

	props = append(props, model.Prop{
		PID:       1003,
		Type:      "skill_enhancer",
		Image:     "prop/skill_enhancer.png",
		Title:     "S攻击强化器",
		Remark:    "S攻击强化器",
		CreatedAt: time.Now(),
	})

	props = append(props, model.Prop{
		PID:       1004,
		Type:      "skill_speed_enhancer",
		Image:     "prop/skill_speed_enhancer.png",
		Title:     "S速度强化器",
		Remark:    "S速度强化器",
		CreatedAt: time.Now(),
	})

	for _, prop := range(props) {
		err = model.DB.Create(prop).Error
		if err != nil {
			log.Println(err.Error())
		}
	}
}