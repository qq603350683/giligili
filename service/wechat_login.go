package service

import (
	"encoding/json"
	"fmt"
	"giligili/constbase"
	"giligili/model"
	"giligili/util"
	"log"
	"os"
)

type WechatLoginResult struct {
	Message string `json:"message"`
	Status int `json:"status"`
	Data *model.UserToken `json:"data"`
}

// 小程序登录时返回的格式
type WechatMiniappLoginResult struct {
	MiniAppOpenid string `json:"openid"`
	SessionKey string `json:"session_key"`
	Unionid string `json:"unionid"`
	Errcode int `json:"-"`
	Errmsg string `json:"-"`
}

func NewWechatLoginResult() *WechatLoginResult {
	return &WechatLoginResult{
		Message: "登录失败",
		Status:  constbase.LOGIN_FAIL,
		Data: nil,
	}
}

func NewWechatMiniappLoginResult() *WechatMiniappLoginResult {
	return &WechatMiniappLoginResult{}
}

func WechantLogin(code string) *WechatLoginResult {
	wechat_login_result := NewWechatLoginResult()

	url := fmt.Sprintf("https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code", os.Getenv("WECHAT_MINIAPP_APPID"), os.Getenv("WECHAT_MINIAPP_APPSECRET"), code)

	str := util.CURL("GET", url, nil, nil)
	if str == "" {
		return wechat_login_result
	}

	wechat_miniapp_login_result := NewWechatMiniappLoginResult()

	err := json.Unmarshal([]byte(str), &wechat_miniapp_login_result)
	if err != nil {
		log.Printf("解析微信小程序Code返回数据失败, 失败详情: %s", err.Error())
		return wechat_login_result
	}

	if wechat_miniapp_login_result.Errcode != 0 {
		log.Printf("解析微信小程序Code返回数据异常 errcode: %d errmsg: %s", wechat_miniapp_login_result.Errcode, wechat_miniapp_login_result.Errmsg)
	}

	//log.Println(wechat_miniapp_login_result)

	//wechat_miniapp_login_result.Unionid = ""
	//wechat_miniapp_login_result.MiniAppOpenid = "qqqqc"
	//wechat_miniapp_login_result.SessionKey = "xxx"

	// 判断当前Openid是否存在
	wechat_user, err := model.GetWechatUserInfoByOpenid(wechat_miniapp_login_result.MiniAppOpenid)
	if err != nil {
		log.Println(err.Error())
		return wechat_login_result
	}

	user_token := model.NewUserToken()

	if wechat_user == nil {
		db := model.DB.Begin()

		// 建立用户信息
		user := model.NewUser()
		if err = db.Create(user).Error; err != nil {
			db.Rollback()
			log.Println(err.Error())
			return wechat_login_result
		}

		// 初级飞机
		user_plan := model.NewUserPlan()
		user_plan.UID = user.UID
		user_plan.IsPutOn = constbase.YES
		user_plan.DetailJson = `{"w":129,"h":120,"texture":"hero/hero2.png","bullets":[{"id":1,"title":"A导弹","w":20,"h":20,"p":3,"a":0,"level":1,"max_level":3,"rate":300,"max_rate":90,"speed":10,"max_speed":8,"texture":"bullet/10.png"}],"skills":[{"id":1,"title":"光速射线","w":50,"h":50,"p":1,"a":0,"level":2,"max_level":4,"rate":600,"max_rate":550,"speed":50,"max_speed":60,"height":9999999,"texture":"bullet/skill1.png"}]}`
		if err = db.Create(user_plan).Error; err != nil {
			db.Rollback()
			log.Println(err.Error())
			return wechat_login_result
		}

		res := db.Model(user).UpdateColumns(map[string]int {
			"up_id": user_plan.UpID,
		})
		if res.RowsAffected == 0 {
			db.Rollback()
			log.Printf("用户(%d) 更换飞机失败", user_plan.UpID)
			return wechat_login_result
		}

		// 建立微信关联用户信息
		new_wechat_user := model.NewWechatUser()
		new_wechat_user.Unionid = wechat_miniapp_login_result.Unionid
		new_wechat_user.MiniAppOpenid = wechat_miniapp_login_result.MiniAppOpenid
		new_wechat_user.UID = user.UID
		if err = db.Create(new_wechat_user).Error; err != nil {
			db.Rollback()
			log.Println(err.Error())
			return wechat_login_result
		}

		// 返回token
		user_token.UID = user.UID
		err = db.Create(user_token).Error
		if err != nil {
			db.Rollback()
			log.Println(err.Error())
			return wechat_login_result
		}

		db.Commit()
	} else {
		// 返回token
		user_token.UID = wechat_user.UID
		err = model.DB.Create(user_token).Error
		if err != nil {
			log.Println(err.Error())
			return wechat_login_result
		}
	}

	wechat_login_result.Message = "登录成功"
	wechat_login_result.Status = constbase.LOGIN_SUCCESS
	wechat_login_result.Data = user_token

	return wechat_login_result
}
