package cache

import (
	"giligili/util"
	"os"
)

var RedisCache *RedisPool

// 初始化缓存
func Init() {
	str := os.Getenv("REDIS_DB")
	rDB, err := util.ToInt(str)
	if err != nil {
		panic(err)
	}

	str = os.Getenv("REDIS_POOL_NUM")
	num, err := util.ToInt(str)
	if err != nil {
		panic(err)
	}

	config := RedisConfig{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       rDB,
		PoolNum:  num,
	}

	SetRedisNil()
	cache, _ := NewRedisPool(config)

	RedisCache = cache
}
