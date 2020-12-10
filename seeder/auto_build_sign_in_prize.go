package seeder

import (
	"fmt"
	"giligili/model"
	"log"
	"time"
)

// 自动创建每个月的奖励
func AutoBuildSignInPrize() {
	sign_in_prizes := []model.SignInPrize{}

	t := time.Now().Format("200601")

	err := model.DB.Where("time = ?", t).Find(&sign_in_prizes).Error
	if err != nil {
		log.Println(err.Error())
		return
	}

	if len(sign_in_prizes) > 0 {
		return
	}

	sign_in_prizes = append(sign_in_prizes, model.SignInPrize{
		PID:        2,
		Quantity:   500,
		Time:       t,
		GrandTotal: 1,
	})

	sign_in_prizes = append(sign_in_prizes, model.SignInPrize{
		PID:        2,
		Quantity:   500,
		Time:       t,
		GrandTotal: 2,
	})

	sign_in_prizes = append(sign_in_prizes, model.SignInPrize{
		PID:        2,
		Quantity:   500,
		Time:       t,
		GrandTotal: 3,
	})

	sign_in_prizes = append(sign_in_prizes, model.SignInPrize{
		PID:        1,
		Quantity:   500,
		Time:       t,
		GrandTotal: 4,
	})

	sign_in_prizes = append(sign_in_prizes, model.SignInPrize{
		PID:        1,
		Quantity:   600,
		Time:       t,
		GrandTotal: 5,
	})

	sign_in_prizes = append(sign_in_prizes, model.SignInPrize{
		PID:        1,
		Quantity:   700,
		Time:       t,
		GrandTotal: 6,
	})

	sign_in_prizes = append(sign_in_prizes, model.SignInPrize{
		PID:        1,
		Quantity:   800,
		Time:       t,
		GrandTotal: 7,
	})

	sign_in_prizes = append(sign_in_prizes, model.SignInPrize{
		PID:        1,
		Quantity:   900,
		Time:       t,
		GrandTotal: 8,
	})

	sign_in_prizes = append(sign_in_prizes, model.SignInPrize{
		PID:        1,
		Quantity:   1000,
		Time:       t,
		GrandTotal: 9,
	})

	sign_in_prizes = append(sign_in_prizes, model.SignInPrize{
		PID:        2,
		Quantity:   500,
		Time:       t,
		GrandTotal: 10,
	})

	sign_in_prizes = append(sign_in_prizes, model.SignInPrize{
		PID:        1,
		Quantity:   500,
		Time:       t,
		GrandTotal: 11,
	})

	sign_in_prizes = append(sign_in_prizes, model.SignInPrize{
		PID:        1,
		Quantity:   600,
		Time:       t,
		GrandTotal: 12,
	})

	sign_in_prizes = append(sign_in_prizes, model.SignInPrize{
		PID:        1,
		Quantity:   700,
		Time:       t,
		GrandTotal: 13,
	})

	sign_in_prizes = append(sign_in_prizes, model.SignInPrize{
		PID:        1,
		Quantity:   800,
		Time:       t,
		GrandTotal: 14,
	})

	sign_in_prizes = append(sign_in_prizes, model.SignInPrize{
		PID:        1,
		Quantity:   900,
		Time:       t,
		GrandTotal: 15,
	})

	sign_in_prizes = append(sign_in_prizes, model.SignInPrize{
		PID:        1,
		Quantity:   1000,
		Time:       t,
		GrandTotal: 16,
	})

	sign_in_prizes = append(sign_in_prizes, model.SignInPrize{
		PID:        2,
		Quantity:   500,
		Time:       t,
		GrandTotal: 17,
	})

	sign_in_prizes = append(sign_in_prizes, model.SignInPrize{
		PID:        1,
		Quantity:   500,
		Time:       t,
		GrandTotal: 18,
	})

	sign_in_prizes = append(sign_in_prizes, model.SignInPrize{
		PID:        1,
		Quantity:   600,
		Time:       t,
		GrandTotal: 19,
	})

	sign_in_prizes = append(sign_in_prizes, model.SignInPrize{
		PID:        1,
		Quantity:   700,
		Time:       t,
		GrandTotal: 20,
	})

	sign_in_prizes = append(sign_in_prizes, model.SignInPrize{
		PID:        1,
		Quantity:   800,
		Time:       t,
		GrandTotal: 21,
	})

	sign_in_prizes = append(sign_in_prizes, model.SignInPrize{
		PID:        1,
		Quantity:   900,
		Time:       t,
		GrandTotal: 22,
	})

	sign_in_prizes = append(sign_in_prizes, model.SignInPrize{
		PID:        1,
		Quantity:   1000,
		Time:       t,
		GrandTotal: 23,
	})

	sign_in_prizes = append(sign_in_prizes, model.SignInPrize{
		PID:        2,
		Quantity:   500,
		Time:       t,
		GrandTotal: 24,
	})

	sign_in_prizes = append(sign_in_prizes, model.SignInPrize{
		PID:        1,
		Quantity:   500,
		Time:       t,
		GrandTotal: 25,
	})

	sign_in_prizes = append(sign_in_prizes, model.SignInPrize{
		PID:        1,
		Quantity:   600,
		Time:       t,
		GrandTotal: 26,
	})

	sign_in_prizes = append(sign_in_prizes, model.SignInPrize{
		PID:        1,
		Quantity:   700,
		Time:       t,
		GrandTotal: 27,
	})

	sign_in_prizes = append(sign_in_prizes, model.SignInPrize{
		PID:        1,
		Quantity:   800,
		Time:       t,
		GrandTotal: 28,
	})

	sign_in_prizes = append(sign_in_prizes, model.SignInPrize{
		PID:        1,
		Quantity:   900,
		Time:       t,
		GrandTotal: 29,
	})

	sign_in_prizes = append(sign_in_prizes, model.SignInPrize{
		PID:        1,
		Quantity:   1000,
		Time:       t,
		GrandTotal: 30,
	})

	sign_in_prizes = append(sign_in_prizes, model.SignInPrize{
		PID:        2,
		Quantity:   500,
		Time:       t,
		GrandTotal: 31,
	})

	sign_in_prizes = append(sign_in_prizes, model.SignInPrize{
		PID:        1,
		Quantity:   500,
		Time:       t,
		GrandTotal: 32,
	})

	sql := `INSERT INTO sign_in_prizes (p_id, quantity, time, grand_total) VALUES `

	length := len(sign_in_prizes) - 1

	for index, sign_in_prize := range(sign_in_prizes) {
		str := ""
		if length == index {
			str = `(%d, %d, "%s", %d) `
		} else {
			str = `(%d, %d, "%s", %d), `
		}

		sql += fmt.Sprintf(str, sign_in_prize.PID, sign_in_prize.Quantity, sign_in_prize.Time, sign_in_prize.GrandTotal)
	}

	err = model.DB.Exec(sql).Error
	if err != nil {
		log.Println(err.Error())
	}
}
