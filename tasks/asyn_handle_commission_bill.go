package tasks

import (
	"giligili/constbase"
	"giligili/model"
	"github.com/shopspring/decimal"
	"log"
	"time"
)

var AsynHandleCommissionBillFlag = false

var AsynHandleCommissionBillMax = 5

// 异步处理未到账账单
func AsynHandleCommissionBill() {
	if AsynHandleCommissionBillFlag == true {
		return
	}

	AsynHandleCommissionBillFlag = true

	time.Sleep(time.Second * 2)

	var bills []*model.Bill

	if err := model.DB.Select("b_id, u_id, price").Where("types_of = ? AND status = ?", constbase.INCOME, constbase.STATUS_COMMISSION_NOT_ARRIVED).Limit(300).Find(&bills).Error; err != nil {
		log.Println(err.Error())
		return
	}

	length := len(bills)

	if length == 0 {
		log.Println("not found")
		return
	} else {
		log.Printf("find %d result", length)
	}

	for _, bill := range(bills) {
		db := model.DB.Begin()

		user := model.GetUserInfo(bill.UID)
		if user == nil {
			continue
		}

		balance_res := decimal.NewFromFloat(user.Balance).Add(decimal.NewFromFloat(bill.Price))
		balance, _ := balance_res.Float64()

		res := db.Model(user).Where("version = ?", user.Version).UpdateColumns(map[string]interface{} {
			"version": user.Version + 1,
			"balance": balance,
		})
		if res.RowsAffected == 0 {
			db.Rollback()
			continue
		}

		res = db.Model(bill).Where("status = ? AND arrived_at = ?", constbase.STATUS_COMMISSION_NOT_ARRIVED, model.DelAtDefault).UpdateColumns(map[string]interface{} {
			"status": constbase.STATUS_COMMISSION_ARRIVED,
			"arrived_at": time.Now(),
		})
		if res.RowsAffected == 0 {
			db.Rollback()
			continue
		}

		db.Commit()
	}

	AsynHandleCommissionBillFlag = false


}
