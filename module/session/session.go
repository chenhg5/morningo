package session

import (
	"morningo/config"
	"morningo/connections/redis"
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

// TODO:
// 1. 根据过期时间与用户认证信息等生成 sessionId
// 2. 判断 sessionId 是否过期
// 3. 通过 sessionId 获取用户信息
// 4. 使得 sessionId 过期
// 5. sessionId 根据 driver 存储：文件/数据库/redis

func generateSessionId() string {
	return ""
}

func isValidSessionId() bool {
	return false
}

func getInfoFromSessionId() []map[string]interface{} {
	info := make([]map[string]interface{}, 0)
	return info
}

func invalidSessionId() bool {
	return true
}

func storeSessionId() bool {
	return true
}