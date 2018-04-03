package session

import (
	"../../config"
	"../../connections/redis"
)

var sessionStore *redis.ClientType

func init() {
	sessionStore = &redis.Client
	sessionStore.RedisCon.Pipeline().Select(config.GetEnv().REDIS_SESSION_DB)
}

func GetStore() *redis.ClientType {
	sessionStore.RedisCon.Pipeline().Select(config.GetEnv().REDIS_CACHE_DB)
	return sessionStore
}