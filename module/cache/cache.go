package cache

import (
	"moringo/config"
	"moringo/connections/redis"
)

var cacheStore *redis.ClientType

func init() {
	cacheStore = &redis.Client
	cacheStore.RedisCon.Pipeline().Select(config.GetEnv().REDIS_CACHE_DB)
}

func GetStore() *redis.ClientType {
	cacheStore.RedisCon.Pipeline().Select(config.GetEnv().REDIS_CACHE_DB)
	return cacheStore
}