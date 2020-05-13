package test

import (
	"giligili/cache"
	"giligili/model"
	"github.com/joho/godotenv"
	"os"
	"reflect"
)

var EmptyListType interface{}

func Init() {
	// 读取本地环境变量
	err := godotenv.Load("../.env")
	if err != nil {
		panic(err)
	}
	// 读取GORM中文错误提示
	//LoadLocales

	EmptyListType = reflect.TypeOf(map[string]interface{} {})

	// 链接数据库
	model.Database(os.Getenv("MYSQL_DSN"))

	// 初始化数据库连接池
	cache.Init()

	client, err := cache.RedisCache.Get()
	if err != nil {
		panic(err)
	}

	client.Do("FlushAll")
}
