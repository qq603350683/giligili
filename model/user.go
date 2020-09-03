package model

import (
	"giligili/cache"
	"github.com/jinzhu/gorm"
	"log"
	"time"
)

type User struct {
	UID int `json:"u_id" gorm:"column:u_id;type:int(10) unsigned auto_increment;primary_key;"`
	UpID int `json:"up_id" gorm:"column:up_id;type:int(10) unsigned; not null; default:0; comment:'飞机ID'"`
	Gold int `json:"gold" gorm:"column:gold;type:int(10) unsigned; not null; default:0; comment:'金币'"`
	Diamond int `json:"diamond" gorm:"column:diamond;type:int(10) unsigned; not null; default:0; comment:'钻石'"`
	Plan *UserPlan `json:"plan" comment:"飞机详情"`
	CreatedAt time.Time `json:"created_at" gorm:"type:datetime;not null; comment:'创建时间'"`
	UpdatedAt time.Time `json:"-" gorm:"type:datetime;not null; comment:'更新时间'"`
	DelAt time.Time `json:"-" gorm:"type:datetime;not null;default:'1000-01-01 00:00:00'; comment:'删除时间'"`
}

// 获取用户详情
func GetUserInfo(u_id int) *User {
	if u_id == 0 {
		return nil
	}

	user := &User{}

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
func (user *User) TodayIsForward() bool {
	key := cache.UserTodayForwardListKey()

	client, err := cache.RedisCache.Get()
	if err != nil {
		log.Printf("Redis 获取失败")
		return false
	}

	res, err := client.SIsMember(key, user.UID).Result()
	if err != nil {
		log.Printf(err.Error())
		return false
	}

	if res == true {
		return true
	}

	return false
}

// 转发
func (user *User) TodayForward() bool {
	key := cache.UserTodayForwardListKey()

	client, err := cache.RedisCache.Get()
	if err != nil {
		log.Printf("Redis 获取失败")
		return false
	}

	i, err := client.SAdd(key, user.UID).Result()
	if err != nil {
		log.Printf("Redis: %s", err.Error())
		return false
	}

	if i != 1 {
		log.Printf("Redis: i 的值非 1")
		return false
	}

	return true
}

// 金币和钻石都增加
func (user *User) GoldAndDiamondIncr(gold, diamond int) bool {
	if gold + diamond == 0 {
		return false
	}

	u := GetUserInfo(UserInfo.UID)

	res := DB.Model(u).Where("gold = ? AND diamond = ?", u.Gold, u.Diamond).Update(map[string]int{
		"gold": u.Gold + gold,
		"diamind": u.Diamond + diamond,
	})

	if res.RowsAffected == 0 {
		return false
	}

	return true
}