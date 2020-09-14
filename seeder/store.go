package seeder

import (
	"giligili/model"
	"log"
)

func StoreRun() {
	stores := []model.Store{}

	err := model.DB.Find(&stores).Error
	if err != nil {
		log.Println(err.Error())
		return
	}

	if len(stores) > 0 {
		return
	}

	sql := `INSERT INTO stores (s_id, title, p_id, plan_id, gold, diamond, quantity, sell, sort, is_shelf, del_at, created_at) VALUES 
(1, "金币大礼包", 1, 0, 0, 5, 1, 0, 99999999, 1, "1000-01-01 00:00:00", "2020-09-10 09:12:18"),
(2, "子弹攻击力强化器", 10001, 0, 1500, 0, 1, 0, 0, 1, "1000-01-01 00:00:00", "2020-09-10 09:12:18"),
(3, "子弹速度强化器", 10002, 0, 1500, 0, 1, 0, 0, 1, "1000-01-01 00:00:00", "2020-09-10 09:12:18"),
(4, "子弹射频强化器", 10003, 0, 1500, 0, 1, 0, 0, 1, "1000-01-01 00:00:00", "2020-09-10 09:12:18"),
(5, "技能攻击力强化器", 10004, 0, 1500, 0, 1, 0, 0, 1, "1000-01-01 00:00:00", "2020-09-10 09:12:18"),
(6, "技能速度强化器", 10005, 0, 1500, 0, 1, 0, 0, 1, "1000-01-01 00:00:00", "2020-09-10 09:12:18"),
(7, "技能射频强化器", 10006, 0, 1500, 0, 1, 0, 0, 1, "1000-01-01 00:00:00", "2020-09-10 09:12:18")
`

	err = model.DB.Exec(sql).Error
	if err != nil {
		log.Println(err.Error())
	}
}