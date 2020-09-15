package model

import (
	"github.com/jinzhu/gorm"
	"log"
	"time"
)

// 全局 token
var Token string

type UserToken struct {
	UtID int `json:"ut_id" gorm:"column:ut_id;type:int(10) unsigned auto_increment; primary_key;"`
	UID int `json:"u_id" gorm:"column:u_id;type:int(10) unsigned; not null; index:idx_u_id; comment:'用户ID'"`
	Token string `json:"token" gorm:"column:token; type: varchar(200); not null; unique_index:uni_token; comment:'token值'"`
	CreatedAt time.Time `json:"created_at" gorm:"type:datetime; not null; comment:'创建时间'"`
	ExpiredAt time.Time `json:"expired_at" gorm:"type:datetime; not null; comment:'过期时间'"`
}

func NewUserToken() *UserToken {
	user_token := &UserToken{}
	user_token.Token = `zzzzz`
	user_token.CreatedAt = time.Now()
	user_token.ExpiredAt = time.Now().AddDate(0, 0, 30)

	return user_token
}

// 根据 token 获取详情
func GetInfoByToken(token string) *UserToken {
	ut := &UserToken{}

	err := DB.Where("token = ?", token).First(ut).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Printf("找不到当前token值: %s", token)
			return nil
		}

		log.Printf("DB: %s", err.Error())

		return nil
	}

	if ut.IsExpired(ut.ExpiredAt) {
		log.Printf("Token 已经过期 %s", token)
		return nil
	}

	return ut
}

// 根据 token 获取 u_id
func GetUIDByToken(token string) int {
	ut := GetInfoByToken(token)
	if ut == nil {
		return 0
	}

	return ut.UID
}

// 判断 token 是否已经过期
func (token *UserToken) IsExpired (expired_at time.Time) bool {
	if expired_at.Unix() < time.Now().Unix() {
		return true
	} else {
		return false
	}
 }