package redis

import (
	"github.com/go-redis/redis"
	"morningo/config"
	"time"
)

type ClientType struct {
	RedisCon *redis.Client
}

var Client ClientType

func init() {
	Client.RedisCon = redis.NewClient(&redis.Options{
		Addr:     config.GetEnv().REDIS_IP + ":" + config.GetEnv().REDIS_PORT,
		Password: config.GetEnv().REDIS_PASSWORD, // no password set
		DB:       config.GetEnv().REDIS_DB,       // use default DB
	})
}

func (client *ClientType) Set(key string, value interface{}, expiration time.Duration) *redis.Client {
	err := client.RedisCon.Set(key, value, expiration).Err()
	if err != nil {
		panic(err)
	}
	return (*client).RedisCon
}

func (client *ClientType) Get(key string) (string, *redis.Client) {
	val, err := (*client).RedisCon.Get(key).Result()

	if err == redis.Nil {
		return "", (*client).RedisCon
	}

	if err != nil {
		panic(err)
	}

	return val, (*client).RedisCon
}

func (client *ClientType) Del(key string) *redis.Client {
	_, err := client.RedisCon.Del(key).Result()
	if err != nil {
		panic(err)
	}
	return (*client).RedisCon
}
