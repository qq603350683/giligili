package cache

import "fmt"

var RedisCache *RedisPool

// 初始化缓存
func Init() {
	cache := RedisInit()
	fmt.Println(cache.Pool)
	RedisCache = cache
}
