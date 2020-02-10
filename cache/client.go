package cache

import (
	"errors"
	"fmt"
	"giligili/util"
	"github.com/go-redis/redis"
	"os"
	"time"
)

// Redis 基本参数
var NewClientAddr string
var NewClientPassword string
var NewClientDB int

// Redis 集群
type RedisPool struct {
	Num int
	Pool chan *redis.Client
}


// 初始化参数
func RedisInit() *RedisPool {
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
	num, err := util.ToInt(str)
	if err != nil {
		panic(err)
	}

	pool := RedisPool{
		Num: num,
	}

	pool.NewRedisPool(pool.Num)

	return &pool
}

// 创建Redis实例
func CreateClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr: NewClientAddr,
		Password: NewClientPassword,
		DB: NewClientDB,
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	return client
}

// 把Redis放到连接池中去
func (r *RedisPool) Put(client *redis.Client) {
	 r.Pool <- client
}

// 取Redis实例使用
func (r *RedisPool) Get() (*redis.Client, error) {
	fmt.Println(r.Pool)
	client, ok := <- r.Pool
	fmt.Println(client)
	if !ok {
		// 等待一下
		time.Sleep(time.Second / 5)
		client, ok = <- r.Pool
		if !ok {
			return nil, errors.New("系统繁忙~")
		}
	}

	return client, nil
}

// 创建Redis连接池
func (r *RedisPool) NewRedisPool(max int) {
	for i := 0;i < max;i++ {
		go func() {
			r.Pool <- CreateClient()
		}()
	}

	time.Sleep(5 * time.Second)
}