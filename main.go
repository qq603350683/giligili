package main

import (
	"giligili/cache"
	"giligili/model"
	"giligili/routes"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	// 读取本地环境变量
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	// 读取GORM中文错误提示
	//LoadLocales

	// 链接数据库
	model.Database(os.Getenv("MYSQL_DSN"))

	// 初始化数据库连接池
	cache.Init()

	routes.NewRouter()


}
