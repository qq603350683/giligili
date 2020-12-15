package model

import "time"

type Bill struct {
	BID int `json:"b_id" gorm:"column:b_id;type:int(10) unsigned auto_increment;not null;primary_key;comment:'账单ID'"`
	UID int `json:"u_id" gorm:"column:u_id; type:int(10) unsigned; not null; default:0; index:idx_u_id; comment:'用户ID 来自 users 表的 u_id'"`
	TypesOf int8 `json:"types_of" gorm:"column:types_of;type:int(1);not null;default:0;comment:'类型 0 - 未知 1 - 收入 2 - 支出'"`
	Price float64 `json:"price" gorm:"column:price;type:decimal(10, 2) unsigned;not null;default:0.01; comment:'金额'"`
	AlipayName string `json:"alipay_name" gorm:"column:alipay_name;type:varchar(4);not null;default:'';comment:'支付宝姓名'"`
	AlipayAccount string `json:"alipay_account" gorm:"column:alipay_account;type:varchar(50);not null;default:'';comment:'支付宝账户'"`
	Status int8 `json:"status" gorm:"column:status;type:int(1);not null;default:0;comment:'状态 0 - 其他 1 - 佣金未到账 2 - 佣金已到账 3 - 提现未到账 4 - 提现已到账 90 - 转账失败'"`
	Remark string `json:"remark" gorm:"column:remark;type:varchar(200);not null;comment:'备注'"`
	ArrivedAt time.Time `json:"-" gorm:"column:arrived_at; type:datetime;not null;default:'1000-01-01 00:00:00';comment:'到账时间'"`
	CreatedAt time.Time `json:"-" gorm:"column:created_at; type:datetime;not null;index:idx_created_at;comment:'创建时间'"`
}

func NewBill() *Bill {
	bill := new(Bill)
	bill.CreatedAt = time.Now()

	return bill
}
