package main

import (
	"giligili/cache"
	"giligili/model"
	"giligili/socket"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	//http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
	//	w.Write([]byte("hello world"))
	//})
	// 读取本地环境变量
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	// 链接数据库
	model.Database(os.Getenv("MYSQL_DSN"))

	// 初始化数据库连接池
	cache.Init()

	socket.Run()
}
