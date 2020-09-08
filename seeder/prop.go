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
		Image:     "prop_gold.png",
		Title:     "金币大礼包",
		Remark:    "金币大礼包",
		CreatedAt: time.Now(),
	})

	props = append(props, model.Prop{
		PID:       2,
		Type:      "diamond",
		Image:     "prop_diamond.png",
		Title:     "钻石大礼包",
		Remark:    "钻石大礼包",
		CreatedAt: time.Now(),
	})

	props = append(props, model.Prop{
		PID:       3,
		Type:      "stone_enhancer_material",
		Image:     "prop_stone_enhancer_material.png",
		Title:     "石岩碳",
		Remark:    "强化子弹、技能攻击力时所需要的石头",
		CreatedAt: time.Now(),
	})

	props = append(props, model.Prop{
		PID:       4,
		Type:      "stone_speed_enhancer_material",
		Image:     "prop_stone_speed_enhancer_material.png",
		Title:     "竹碳",
		Remark:    "强化子弹、技能速度时所需要的石头",
		CreatedAt: time.Now(),
	})

	props = append(props, model.Prop{
		PID:       1001,
		Type:      "bullet_enhancer",
		Image:     "prop_bullet_enhancer.png",
		Title:     "子弹攻击力强化器",
		Remark:    "子弹攻击力强化器",
		CreatedAt: time.Now(),
	})

	props = append(props, model.Prop{
		PID:       1002,
		Type:      "bullet_speed_enhancer",
		Image:     "prop_bullet_speed_enhancer.png",
		Title:     "子弹速度强化器",
		Remark:    "子弹速度强化器",
		CreatedAt: time.Now(),
	})

	props = append(props, model.Prop{
		PID:       1003,
		Type:      "skill_enhancer",
		Image:     "prop_skill_enhancer.png",
		Title:     "技能攻击力强化器",
		Remark:    "技能攻击力强化器",
		CreatedAt: time.Now(),
	})

	props = append(props, model.Prop{
		PID:       1004,
		Type:      "skill_speed_enhancer",
		Image:     "prop_skill_speed_enhancer.png",
		Title:     "技能速度强化器",
		Remark:    "技能速度强化器",
		CreatedAt: time.Now(),
	})

	for _, prop := range(props) {
		err = model.DB.Create(prop).Error
		if err != nil {
			log.Println(err.Error())
		}
	}
}