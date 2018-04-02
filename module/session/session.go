package session

import (
	"../../config"
	"../../connections/redis"
)

var SessionStore *redis.ClientType

func init() {
	SessionStore = &redis.Client
	SessionStore.RedisCon.Pipeline().Select(config.GetEnv().REDIS_SESSION_DB)
}
