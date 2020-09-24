package seeder

import (
	"giligili/model"
	"log"
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

	sql := `INSERT INTO props (p_id, type, image, title, remark, min_quantity, max_quantity, gold_value, diamond_value, created_at) VALUES 
(1, "gold", "prop/gold.png", "金币", "金币", 0, 0, 20, 0, "2020-09-10 09:12:18"),
(2, "diamond", "prop/diamond.png", "钻石", "钻石", 0, 0, 0, 1, "2020-09-10 09:12:18"),
(1001, "gold_pack", "prop/gold.png", "金币大礼包", "金币大礼包(100~500)", 100, 500, 20, 0, "2020-09-10 09:12:18"),
(1002, "diamond_pack", "prop/diamond.png", "钻石大礼包", "钻石大礼包(20~50)", 20, 50, 0, 1, "2020-09-10 09:12:18"),
(5001, "stone_enhancer_material", "prop/stone_enhancer_material.png", "石岩碳", "强化子弹、技能攻击力时所需要的石头", 0, 0, 10, 0, "2020-09-10 09:12:18"),
(5002, "stone_speed_enhancer_material", "prop/stone_speed_enhancer_material.png", "竹碳", "强化子弹、技能速度时所需要的石头", 0, 0, 10, 0, "2020-09-10 09:12:18"),
(10001, "bullet_enhancer", "prop/bullet_enhancer.png", "B攻击强化器", "B攻击强化器", 0, 0, 300, 0, "2020-09-10 09:12:18"),
(10002, "bullet_speed_enhancer", "prop/bullet_speed_enhancer.png", "B速度强化器", "B速度强化器", 0, 0, 300, 0, "2020-09-10 09:12:18"),
(10003, "bullet_rate_enhancer", "prop/bullet_rate_enhancer.png", "B频率强化器", "B频率强化器", 0, 0, 300, 0, "2020-09-10 09:12:18"),
(10004, "skill_enhancer", "prop/skill_enhancer.png", "S攻击强化器", "S攻击强化器", 0, 0, 300, 0, "2020-09-10 09:12:18"),
(10005, "skill_speed_enhancer", "prop/skill_speed_enhancer.png", "S速度强化器", "S速度强化器", 0, 0, 300, 0, "2020-09-10 09:12:18"),
(10006, "skill_rate_enhancer", "prop/skill_rate_enhancer.png", "S频率强化器", "S频率强化器", 0, 0, 300, 0, "2020-09-10 09:12:18")
`

	err = model.DB.Exec(sql).Error
	if err != nil {
		log.Println(err.Error())
	}
}