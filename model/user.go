package model

import (
	"github.com/jinzhu/gorm"
	"log"
	"time"
)

type User struct {
	UID int `json:"u_id" gorm:"column:u_id;type:int(10) unsigned auto_increment;primary_key;"`
	UpID int `json:"up_id" gorm:"column:up_id;type:int(10) unsigned; not null; default:0; comment:'飞机ID'"`
	Gold int `json:"gold" gorm:"column:gold;type:int(10) unsigned; not null; default:0; comment:'金币'"`
	Diamond int `json:"diamond" gorm:"column:diamond;type:int(10) unsigned; not null; default:0; comment:'钻石'"`
	Balance float64 `json:"balance" gorm:"column:balance;type:decimal(10, 2) unsigned;not null;default:0.01; comment:'余额'"`
	AlipayName string `json:"alipay_name" gorm:"column:alipay_name;type:varchar(4);not null;default:'';comment:'支付宝姓名'"`
	AlipayAccount string `json:"alipay_account" gorm:"column:alipay_account;type:varchar(50);not null;default:'';comment:'支付宝账户'"`
	LID int `json:"l_id" gorm:"column:l_id;type:int(10) unsigned; not null; default:0; comment:'最高通关级别'"`
	Version int `json:"-" gorm:"column:version;type:int(10) unsigned; not null; default:0; comment:'版本号'"`
	Plan *UserPlan `json:"plan" comment:"飞机详情"`
	LoginedAt time.Time `json:"logined_at" gorm:"type:datetime;not null;default:'1000-01-01 00:00:00'; comment:'最近一次时间'"`
	CreatedAt time.Time `json:"created_at" gorm:"type:datetime;not null; comment:'创建时间'"`
	UpdatedAt time.Time `json:"-" gorm:"type:datetime;not null; comment:'更新时间'"`
	DelAt time.Time `json:"-" gorm:"type:datetime;not null;default:'1000-01-01 00:00:00'; comment:'删除时间'"`
}

func NewUser() *User {
	user := new(User)
	user.CreatedAt = time.Now()

	return user
}

// 更新最后一次登录时间
func (user *User) UpdateLoginedAt() {
	DB.Model(user).Update("logined_at", time.Now())
}

// 获取用户详情
func GetUserInfo(u_id int) *User {
	if u_id == 0 {
		return nil
	}

	user := NewUser()

	err := DB.Where("u_id = ?", u_id).First(user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Printf("用户ID(%d)找不到记录: ", u_id)
		} else {
			log.Println(err.Error())
		}
		return nil
	}

	if IsDel(user.DelAt) {
		log.Printf("用户ID(%d)已删除: ", u_id)
		return nil
	}

	user.Plan = GetUserPlanInfo(user.UpID)

	return user
}

// 判断今天是否已经转发， 每个用户记录一次
//func (user *User) TodayIsForward() bool {
//	key := cache.UserTodayForwardListKey()
//
//	client, err := cache.RedisCache.Get()
//	if err != nil {
//		log.Printf("Redis 获取失败")
//		return false
//	}
//
//	res, err := client.SIsMember(key, user.UID).Result()
//	if err != nil {
//		log.Printf(err.Error())
//		return false
//	}
//
//	if res == true {
//		return true
//	}
//
//	return false
//}
//
//// 转发
//func (user *User) TodayForward() bool {
//	key := cache.UserTodayForwardListKey()
//
//	client, err := cache.RedisCache.Get()
//	if err != nil {
//		log.Printf("Redis 获取失败")
//		return false
//	}
//
//	i, err := client.SAdd(key, user.UID).Result()
//	if err != nil {
//		log.Printf("Redis: %s", err.Error())
//		return false
//	}
//
//	if i != 1 {
//		log.Printf("Redis: i 的值非 1")
//		return false
//	}
//
//	return true
//}

// 金币和钻石都增加
func (user *User) GoldAndDiamondUpdate(gold, diamond int) bool {
	if gold + diamond == 0 {
		return false
	}

	res := DB.Model(UserInfo).Where("gold = ? AND diamond = ?", UserInfo.Gold, UserInfo.Diamond).Updates(map[string]int{
		"gold": UserInfo.Gold + gold,
		"diamond": UserInfo.Diamond + diamond,
	})

	if res.RowsAffected == 0 {
		log.Println("更新数据失败")
		return false
	}

	return true
}

// 领取通关奖励
func (user *User) GetPassLevelPrize(l_id, gold, diamond int) bool {
	if UserInfo.LID <= l_id && gold + diamond == 0 {
		return false
	}

	new_l_id := UserInfo.LID

	if (l_id > UserInfo.LID) {
		new_l_id = l_id
	}

	db := GetDB()

	res := db.Model(UserInfo).Where("gold = ? AND diamond = ? AND l_id = ?", UserInfo.Gold, UserInfo.Diamond, UserInfo.LID).Updates(map[string]int{
		"l_id": new_l_id,
		"gold": UserInfo.Gold + gold,
		"diamond": UserInfo.Diamond + diamond,
	})

	if res.RowsAffected == 0 {
		log.Println("更新数据失败")
		return false
	}

	return true
}

func (user *User) GetRegisterCount(start string, end string) int {
	u := NewUser()

	query := DB.Model(u)
	count := 0

	if start != "" && end != "" {
		query = query.Where("created_at >= ? and created_at <= ?", start, end)
	} else if start == "" && end != "" {
		query = query.Where("created_at <= ?", end)
	} else if start != "" && end == "" {
		query = query.Where("created_at >= ?", start)
	}

	err := query.Count(&count).Error
	if err != nil {
		log.Println(err)
		return 0
	}

	return count
}

func (user *User) GetPassLevelCount(start string, end string) int {
	u := NewUser()

	query := DB.Model(u).Where("l_id >= 1")
	count := 0

	if start != "" && end != "" {
		query = query.Where("created_at >= ? and created_at <= ?", start, end)
	} else if start == "" && end != "" {
		query = query.Where("created_at <= ?", end)
	} else if start != "" && end == "" {
		query = query.Where("created_at >= ?", start)
	}

	err := query.Count(&count).Error
	if err != nil {
		log.Println(err)
		return 0
	}

	return count
}

// 更换飞机
//func (user *User) ChangePlan(up_id int) bool {
//	if UserInfo.UpID == up_id {
//		return true
//	}
//
//	current_user_plan := GetUserPlanInfo(UserInfo.UpID)
//	if current_user_plan == nil {
//		return false
//	}
//
//	if current_user_plan.UID != UserInfo.UID {
//		return false
//	}
//
//	db := DB.Begin()
//
//	user_plan := GetUserPlanInfo(up_id)
//	if user_plan == nil {
//		db.Rollback()
//		return false
//	}
//
//	if user_plan.UID != UserInfo.UID {
//		db.Rollback()
//		return false
//	}
//
//	res := db.Model(user_plan).Update("is_put_on", constbase.YES)
//	if res.RowsAffected == 0 {
//		db.Rollback()
//		return false
//	}
//
//	res = db.Model(current_user_plan).Update("is_put_on", constbase.NO)
//	if res.RowsAffected == 0 {
//		db.Rollback()
//		return false
//	}
//
//	res = db.Model(UserInfo).Update("up_id", up_id)
//	if res.RowsAffected == 0 {
//		db.Rollback()
//		return false
//	}
//
//	db.Commit()
//
//	UserInfo.Plan = user_plan
//
//	return true
//}