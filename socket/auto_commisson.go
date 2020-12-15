package socket

import (
	"fmt"
	"giligili/constbase"
	"giligili/model"
	"giligili/serializer"
	"giligili/util"
	"github.com/shopspring/decimal"
	"log"
	"math"
	"net/http"
	"strconv"
	"time"
)

var RAKE = 30

type CommissionList struct {
	UID int `json:"u_id" gorm:"column:u_id;" comment:"用户ID"`
	Num int `json:"num" gorm:"column:num;" comment:"分成比例"`
}

func AutoCommission(params Params) {
	// 总金额
	var total_income float64
	// 总支出
	var total_expenditure float64
	// 抽成
	var rake_price float64

	// 先获取昨天时间为默认时间
	target_date := time.Now().AddDate(0, 0, -1).Format(util.DATE)

	if td, ok := params["target_date"]; ok && len(td) > 0 {
		target_date = td
	}

	if _, ok := params["total_income"]; ok {
		tmp_total_income, err := strconv.ParseFloat(params["total_income"], 64)
		if err != nil {
			SendMessage(model.UserInfo.UID, serializer.JsonByte(http.StatusInternalServerError, "请输入正确的收入金额", nil, err.Error()))
			return
		}

		if tmp_total_income > 0 {
			total_income = tmp_total_income
			res := decimal.NewFromFloat(total_income).Mul(decimal.NewFromFloat(float64(RAKE) / float64(100)))
			rake_price, _ = res.Float64()
			res = decimal.NewFromFloat(total_income).Sub(decimal.NewFromFloat(rake_price))
			total_income, _ = res.Float64()
		} else {
			rake_price = 0
			total_income = 0
		}
	} else {
		rake_price = 0
		total_income = 0
	}

	// 查看目标日期是否已经分佣完毕
	is_commission := model.IsCommission(target_date)
	if is_commission == true {
		SendMessage(model.UserInfo.UID, serializer.JsonByte(http.StatusOK, target_date + "已成功分佣", nil, ""))
		return
	}

	var commission_list []CommissionList
	if err := model.DB.Raw("SELECT u_id, COUNT(u_id) AS num FROM user_adv_records WHERE created_at >= ? AND created_at <= ? GROUP BY u_id", target_date + " 00:00:00", target_date + " 23:59:59").Find(&commission_list).Error; err != nil {
		log.Println(err.Error())
		SendMessage(model.UserInfo.UID, serializer.JsonByte(http.StatusInternalServerError, "系统错误，请稍后再试", nil, err.Error()))
		return
	}

	uv := len(commission_list)
	sql := ""
	total_num := 0

	db := model.DBBegin()

	defer model.CancelDB()

	if uv > 0 {
		for _, com := range(commission_list) {
			total_num += com.Num
		}

		sql = `INSERT INTO bills (u_id, types_of, price, status, remark, created_at) VALUES `

		default_price := 0.01

		length := uv - 1

		for index, com := range(commission_list) {
			price := default_price

			if total_income > 0 {
				res := decimal.NewFromFloat(float64(com.Num) / float64(total_num) * total_income).Sub(decimal.NewFromFloat(0.005))
				price, _ = res.Float64()
				if price <= default_price {
					// 小于 default_price 0.01 则按0.01算
					price = default_price
				} else {
					price = float64(math.Ceil(price * 100) / 100)
				}
			}

			total_expenditure += price

			str := ""

			if length == index {
				str = `(%d, %d, %f, %d, "%s", "%s") `
			} else {
				str = `(%d, %d, %f, %d, "%s", "%s"), `
			}

			sql += fmt.Sprintf(str, com.UID, constbase.INCOME, price, constbase.STATUS_COMMISSION_NOT_ARRIVED, "广告观看收入", time.Now().Format(util.DATETIME))
		}

		err := db.Exec(sql).Error
		if err != nil {
			db.Rollback()
			log.Println(err.Error())
			SendMessage(model.UserInfo.UID, serializer.JsonByte(http.StatusInternalServerError, "插入账单失败", nil, ""))
			return
		}
	}

	res := decimal.NewFromFloat(total_income).Sub(decimal.NewFromFloat(total_expenditure))

	commission := model.NewCommission()
	commission.UV = uv
	commission.AdvNum = total_num
	commission.IncomePrice, _ = res.Float64()
	commission.IncomePrice += rake_price  // 这里 +rake_price 是系统抽取的拥佣金
	commission.PayPrice = total_expenditure
	commission.Remark = target_date + "佣金分成"

	if err := db.Create(commission).Error; err != nil {
		db.Rollback()
		log.Println(err.Error())
		SendMessage(model.UserInfo.UID, serializer.JsonByte(http.StatusInternalServerError, "插入总账单失败", nil, ""))
		return
	}

	db.Commit()

	SendMessage(model.UserInfo.UID, serializer.JsonByte(http.StatusOK, "success", sql, ""))
}
