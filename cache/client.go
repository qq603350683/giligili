package cache

import (
	"errors"
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

// Redis 集群
type RedisPool struct {
	Num int
	Pool chan *redis.Client
}

// Redis 配置
type RedisConfig struct {
	Addr string
	Password string
	DB int
	PoolNum int
}

// 创建Redis实例
func CreateClient(config RedisConfig) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr: config.Addr,
		Password: config.Password,
		DB: config.DB,
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
	client, ok := <- r.Pool

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
func NewRedisPool(config RedisConfig) (*RedisPool, error) {
	if config.PoolNum == 0 {
		return &RedisPool{}, errors.New("请输入连接池个数")
	}

	pool := &RedisPool{
		Num:  config.PoolNum,
		Pool: make(chan *redis.Client, config.PoolNum),
	}

	go func () {
		for i := 0;i < config.PoolNum;i++ {
			go func(i int) {
				pool.Pool <- CreateClient(config)
			}(i)
		}
	}()

	return pool, nil
}