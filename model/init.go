package model

import (
	"giligili/util"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"time"
)

var DB *gorm.DB
var DBTransaction *gorm.DB

var DelAtDefault string
var DelAtDefaultTime time.Time

var EmptyList []map[string]interface{}

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

	// 开启SQL打印模式
	db.LogMode(true)

	DB = db

	// 软删除默认值设置
	DelAtDefault = "1000-01-01 00:00:00"

	DelAtDefaultTime = util.ToTime(DelAtDefault)

	migration()
}

// 取消 DB 事物句柄
func CancelDB() {
	if DBTransaction != nil {
		DBTransaction = nil
		log.Println("DBTransaction 清空成功！")
	}
}

// 获取 DB 句柄
func GetDB() *gorm.DB {
	if DBTransaction != nil {
		log.Println("DBTransaction 获取成功！")
		return DBTransaction
	} else {
		log.Println("DB 获取成功！")
		return DB
	}
}

// 开启事物并获取 DB
func DBBegin() *gorm.DB {
	log.Println("DB Begin")
	DBTransaction = DB.Begin()
	return DBTransaction
}
