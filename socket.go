package main

import (
	"giligili/cache"
	"giligili/model"
	"giligili/seeder"
	"giligili/socket"
	"giligili/tasks"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	//http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
	//	w.Write([]byte("hello world"))
	//})
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	// 读取本地环境变量
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	// 链接数据库
	model.Database(os.Getenv("MYSQL_DSN"))

	// 初始化数据库连接池
	cache.Init()

	go tasks.SendFakeMessage()

	seeder.Run()

	socket.Run()
}
