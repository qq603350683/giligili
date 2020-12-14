package model

import (
	"giligili/util"
	"github.com/jinzhu/gorm"
	"log"
	"time"
)

type Commission struct {
	CID int `json:"c_id" gorm:"column:c_id;type:int(10) unsigned auto_increment;not null;primary_key;comment:'分佣ID'"`
	IncomePrice float64 `json:"income_price" gorm:"column:income_price;type:decimal(10, 2);not null;default:0; comment:'收入金额'"`
	PayPrice float64 `json:"pay_price" gorm:"column:pay_price;type:decimal(10, 2) unsigned;not null;default:0; comment:'支出金额'"`
	UV int `json:"uv" gorm:"column:uv; type:int(10) unsigned;not null;default:0;comment:'受益用户数'"`
	AdvNum int `json:"adv_num" gorm:"column:adv_num; type:int(10) unsigned;not null;default:0;comment:'总观看次数'"`
	Remark string `json:"remark" gorm:"column:remark;type:varchar(200);not null;comment:'备注'"`
	CreatedAt time.Time `json:"-" gorm:"column:created_at; type:datetime;not null;index:idx_created_at;comment:'创建时间'"`
}

func NewCommission() *Commission {
	commission := new(Commission)
	commission.CreatedAt = time.Now()

	return commission
}

// 判断目标日期是否已经分佣成功
func IsCommission(date string) bool {
	b := util.IsDatetime(date)
	if (b == false) {
		return false
	}

	db := GetDB()

	commission := NewCommission()

	if err := db.Where("created_at >= ? AND created_at <= ?", date, date + " 23:59:59").First(&commission).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return false
		} else {
			log.Printf("系统错误: %s", err.Error())
			return false
		}
	} else {
		return true
	}
}
