package tasks

import (
	"giligili/constbase"
	"giligili/model"
	"github.com/shopspring/decimal"
	"log"
	"time"
)

var AsynHandleCommissionBillFlag = false

// 控制并发数
var AsynHandleCommissionBillMax = 5

// 异步处理未到账账单
func AsynHandleCommissionBill() {
	c := make(chan bool, AsynHandleCommissionBillMax)

	for i := 0;i < AsynHandleCommissionBillMax; i++ {
		c <- true
	}

	for {
		if AsynHandleCommissionBillFlag == true {
			continue
		}

		AsynHandleCommissionBillFlag = true

		time.Sleep(time.Second * 5)

		var bills []*model.Bill

		if err := model.DB.Select("b_id, u_id, price").Where("types_of = ? AND status = ?", constbase.INCOME, constbase.STATUS_COMMISSION_NOT_ARRIVED).Limit(300).Find(&bills).Error; err != nil {
			log.Println(err.Error())
			AsynHandleCommissionBillFlag = false
			continue
		}

		length := len(bills)

		if length == 0 {
			//log.Println("not found any commission bills")
			AsynHandleCommissionBillFlag = false
			continue
		} else {
			log.Printf("find %d commission bills", length)
		}

		for _, bill := range(bills) {
			if _, ok := <- c; ok {
				go func(bill *model.Bill) {
					db := model.DB.Begin()

					user := model.GetUserInfo(bill.UID)
					if user == nil {
						db.Rollback()
						c <- true
						return
					}

					balance_res := decimal.NewFromFloat(user.Balance).Add(decimal.NewFromFloat(bill.Price))
					balance, _ := balance_res.Float64()

					res := db.Model(user).Where("version = ?", user.Version).UpdateColumns(map[string]interface{} {
						"version": user.Version + 1,
						"balance": balance,
					})
					if res.RowsAffected == 0 {
						db.Rollback()
						c <- true
						return
					}

					res = db.Model(bill).Where("status = ? AND arrived_at = ?", constbase.STATUS_COMMISSION_NOT_ARRIVED, model.DelAtDefault).UpdateColumns(map[string]interface{} {
						"status": constbase.STATUS_COMMISSION_ARRIVED,
						"arrived_at": time.Now(),
					})
					if res.RowsAffected == 0 {
						db.Rollback()
						c <- true
						return
					}

					db.Commit()
					c <- true
				} (bill)
			}
		}

		AsynHandleCommissionBillFlag = false
	}
}

