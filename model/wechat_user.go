package model

import (
	"errors"
	"github.com/jinzhu/gorm"
	"log"
)

type WechatUser struct {
	WuID int `json:"wu_id" gorm:"column:wu_id; type:int(10) unsigned auto_increment; not null; primary_key"`
	Openid string `json:"openid" gorm:"column:openid;not null;default:'';comment:'openid'"`
	Unionid string `json:"unionid" gorm:"column:unionid;not null;default:'';comment:'unionid'"`
	MiniAppOpenid string `json:"mini_app_openid" gorm:"column:mini_app_openid;not null;default:'';comment:'小程序openid'"`
	UID int `json:"u_id" gorm:"column:u_id; type:int(10) unsigned; not null; default:0; index:idx_u_id; comment:'用户ID 来自 users 表的 u_id'"`
}

func NewWechatUser() *WechatUser {
	return &WechatUser{}
}

func GetWechatUserInfoByOpenid(mini_app_openid string) (*WechatUser, error) {
	if mini_app_openid == "" {
		return nil, errors.New("Openid不能为空")
	}

	wechat_user := &WechatUser{}

	err := DB.Where("mini_app_openid = ?", mini_app_openid).First(wechat_user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		} else {
			log.Println(err.Error())
		}
		return nil, err
	}

	return wechat_user, nil
}