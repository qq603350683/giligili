package socket

import (
	"giligili/constbase"
	"giligili/model"
	"giligili/serializer"
	"github.com/shopspring/decimal"
	"net/http"
	"strconv"
)

func Withdraw(params Params) {
	price := 0.00
	alipay_name := ""
	alipay_account := ""

	if _, ok := params["price"]; !ok {
		SendMessage(model.UserInfo.UID, serializer.JsonByte(http.StatusInternalServerError, "请选择提现的金额", nil, ""))
		return
	}

	if _, ok := params["alipay_name"]; !ok || len(params["alipay_name"]) == 0 {
		SendMessage(model.UserInfo.UID, serializer.JsonByte(http.StatusInternalServerError, "请输入支付宝姓名", nil, ""))
		return
	} else {
		alipay_name = params["alipay_name"]
	}

	if _, ok := params["alipay_account"]; !ok || len(params["alipay_account"]) == 0 {
		SendMessage(model.UserInfo.UID, serializer.JsonByte(http.StatusInternalServerError, "请输入支付宝账户", nil, ""))
		return
	} else {
		alipay_account = params["alipay_account"]
	}

	_price, err := strconv.Atoi(params["price"])
	if err != nil {
		SendMessage(model.UserInfo.UID, serializer.JsonByte(http.StatusInternalServerError, "仅允许提现整数金额哦~", nil, err.Error()))
		return
	}

	price = float64(_price)

	if model.UserInfo.Balance < price {
		SendMessage(model.UserInfo.UID, serializer.JsonByte(http.StatusInternalServerError, "提现余额不足~", nil, ""))
		return
	}

	db := model.DBBegin()

	res := decimal.NewFromFloat(model.UserInfo.Balance).Sub(decimal.NewFromFloat(price))
	balance, _ := res.Float64()

	row := db.Model(model.UserInfo).Where("version = ?", model.UserInfo.Version).UpdateColumns(map[string]interface{} {
		"version": model.UserInfo.Version + 1,
		"balance": balance,
		"alipay_name": alipay_name,
		"alipay_account": alipay_account,
	})
	if row.RowsAffected == 0 {
		db.Rollback()
		SendMessage(model.UserInfo.UID, serializer.JsonByte(http.StatusInternalServerError, "提现失败", nil, "扣除账户余额失败"))
		return
	}

	bill := model.NewBill()
	bill.UID = model.UserInfo.UID
	bill.TypesOf = constbase.EXPENDITURE
	bill.Price = price
	bill.Status = constbase.STATUS_WITHDRAW_NOT_ARRIVED
	bill.Remark = "提现"
	bill.AlipayName = alipay_name
	bill.AlipayAccount = alipay_account

	if err = db.Create(bill).Error; err != nil {
		db.Rollback()
		SendMessage(model.UserInfo.UID, serializer.JsonByte(http.StatusInternalServerError, "提现失败", nil, err.Error()))
		return
	}

	db.Commit()

	//log.Println(price)
	//log.Println(alipay_name)
	//log.Println(alipay_account)

	SendMessage(model.UserInfo.UID, serializer.JsonByte(http.StatusOK, "提交申请成功，将在7个工作日内完成提现", nil, ""))
}
