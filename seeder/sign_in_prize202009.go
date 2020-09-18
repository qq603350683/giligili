package seeder

import (
	"fmt"
	"giligili/model"
	"log"
)

func SignInPrizeRun() {
	sign_in_prizes := []model.SignInPrize{}

	err := model.DB.Where("time = ?", 202009).Find(&sign_in_prizes).Error
	if err != nil {
		log.Println(err.Error())
		return
	}

	if len(sign_in_prizes) > 0 {
		return
	}

	sign_in_prizes = append(sign_in_prizes, model.SignInPrize{
		SipID:      0,
		PID:        2,
		Quantity:   500,
		Time:       "202009",
		GrandTotal: 1,
	})

	sign_in_prizes = append(sign_in_prizes, model.SignInPrize{
		SipID:      0,
		PID:        1,
		Quantity:   500,
		Time:       "202009",
		GrandTotal: 2,
	})

	sign_in_prizes = append(sign_in_prizes, model.SignInPrize{
		SipID:      0,
		PID:        1,
		Quantity:   550,
		Time:       "202009",
		GrandTotal: 3,
	})

	sign_in_prizes = append(sign_in_prizes, model.SignInPrize{
		SipID:      0,
		PID:        1,
		Quantity:   600,
		Time:       "202009",
		GrandTotal: 4,
	})

	sign_in_prizes = append(sign_in_prizes, model.SignInPrize{
		SipID:      0,
		PID:        1,
		Quantity:   650,
		Time:       "202009",
		GrandTotal: 5,
	})

	sign_in_prizes = append(sign_in_prizes, model.SignInPrize{
		SipID:      0,
		PID:        1,
		Quantity:   700,
		Time:       "202009",
		GrandTotal: 6,
	})

	sign_in_prizes = append(sign_in_prizes, model.SignInPrize{
		SipID:      0,
		PID:        2,
		Quantity:   100,
		Time:       "202009",
		GrandTotal: 7,
	})

	sign_in_prizes = append(sign_in_prizes, model.SignInPrize{
		SipID:      0,
		PID:        1,
		Quantity:   1000,
		Time:       "202009",
		GrandTotal: 8,
	})


	sign_in_prizes = append(sign_in_prizes, model.SignInPrize{
		SipID:      0,
		PID:        1,
		Quantity:   500,
		Time:       "202009",
		GrandTotal: 9,
	})

	sign_in_prizes = append(sign_in_prizes, model.SignInPrize{
		SipID:      0,
		PID:        1,
		Quantity:   550,
		Time:       "202009",
		GrandTotal: 10,
	})

	sign_in_prizes = append(sign_in_prizes, model.SignInPrize{
		SipID:      0,
		PID:        1,
		Quantity:   600,
		Time:       "202009",
		GrandTotal: 11,
	})

	sign_in_prizes = append(sign_in_prizes, model.SignInPrize{
		SipID:      0,
		PID:        1,
		Quantity:   650,
		Time:       "202009",
		GrandTotal: 12,
	})

	sign_in_prizes = append(sign_in_prizes, model.SignInPrize{
		SipID:      0,
		PID:        1,
		Quantity:   700,
		Time:       "202009",
		GrandTotal: 13,
	})

	sign_in_prizes = append(sign_in_prizes, model.SignInPrize{
		SipID:      0,
		PID:        2,
		Quantity:   100,
		Time:       "202009",
		GrandTotal: 14,
	})

	sign_in_prizes = append(sign_in_prizes, model.SignInPrize{
		SipID:      0,
		PID:        1,
		Quantity:   1000,
		Time:       "202009",
		GrandTotal: 15,
	})

	sign_in_prizes = append(sign_in_prizes, model.SignInPrize{
		SipID:      0,
		PID:        1,
		Quantity:   550,
		Time:       "202009",
		GrandTotal: 16,
	})

	sign_in_prizes = append(sign_in_prizes, model.SignInPrize{
		SipID:      0,
		PID:        1,
		Quantity:   600,
		Time:       "202009",
		GrandTotal: 17,
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
