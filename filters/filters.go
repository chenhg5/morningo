package filters

import (
	"github.com/gin-contrib/cache"
	"github.com/gin-contrib/cache/persistence"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"morningo/config"
	"time"
)

func RegisterSession() gin.HandlerFunc {
	store, _ := sessions.NewRedisStore(
		10,
		"tcp",
		config.GetEnv().RedisIp+":"+config.GetEnv().RedisPort,
		config.GetEnv().RedisPassword,
		[]byte(config.GetEnv().SessionSecret))
	return sessions.Sessions(config.GetEnv().SessionKey, store)
}

func RegisterCache() gin.HandlerFunc {
	var cacheStore persistence.CacheStore
	cacheStore = persistence.NewRedisCache(config.GetEnv().RedisIp+":"+config.GetEnv().RedisPort, config.GetEnv().RedisPassword, time.Minute)
	return cache.Cache(&cacheStore)
}
