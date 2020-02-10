package cache

import (
	"fmt"
	"giligili/util"
	"github.com/go-redis/redis"
	"os"
)

// Redis 集群
var RedisPoolNum int
var RedisPool chan *redis.Client

// Redis 基本参数
var NewClientAddr string
var NewClientPassword string
var NewClientDB int

// 初始化参数
func Init() {
	var str string
	var i int

	NewClientAddr = os.Getenv("REDIS_ADDR")
	NewClientPassword = os.Getenv("REDIS_PASSWORD")

	str = os.Getenv("REDIS_DB")
	i, err := util.ToInt(str)
	if err != nil {
		panic(err)
	}
	NewClientDB = i

	str = os.Getenv("REDIS_POOL_NUM")
	i, err = util.ToInt(str)
	if err != nil {
		panic(err)
	}
	RedisPoolNum = i

	// 开启连接池
	NewRedisPool(RedisPoolNum)
}

// 创建Redis实例
func CreateClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr: NewClientAddr,
		Password: NewClientPassword,
		DB: NewClientDB,
	})

	return client
}

// 把Redis放到连接池中去
func PutRedis(client *redis.Client) {
	RedisPool <- client
}

// 取Redis实例使用
func GetRedis() (*redis.Client, error) {
	fmt.Println(RedisPool)
	client := <- RedisPool
	fmt.Println(client)
	//if !ok {
	//	// 等待一下
	//	time.Sleep(time.Second / 5)
	//	client, ok = <- RedisPool
	//	if !ok {
	//		return nil, errors.New("系统繁忙~")
	//	}
	//}

	return client, nil
}

// 创建Redis连接池
func NewRedisPool(max int) {
	//a := make(chan *redis.Client, max)
	for i := 0;i < max;i++ {
		go func() {
			client := CreateClient()

			RedisPool <- client

			fmt.Println(RedisPool)
		}()
	}
}