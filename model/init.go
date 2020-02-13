package model

import (
	"giligili/util"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

var DB *gorm.DB

var DelAtDefault string
var DelAtDefaultTime time.Time

// Database 在中间件仲初始化Mysql链接
func Database(connString string) {
	// connString 用户名:密码@(主机地址:端口)/数据库名称?charset=utf8&parseTime=True&loc=Local
	db, err := gorm.Open("mysql", connString)
	if err != nil {
		panic(err)
	}

	// 设置连接池
	// 空闲
	db.DB().SetMaxIdleConns(20)

	// 打开
	db.DB().SetMaxOpenConns(20)

	// 超时
	db.DB().SetConnMaxLifetime(time.Second * 30)

	DB = db

	// 软删除默认值设置
	DelAtDefault = "1000-01-01 00:00:00"

	DelAtDefaultTime = util.ToTime(DelAtDefault)

	migration()
}
