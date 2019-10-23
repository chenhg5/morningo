package redis

import (
	"github.com/go-redis/redis"
	"morningo/config"
	"time"
)

type ClientType struct {
	Conn *redis.Client
}

var Client ClientType

func init() {
	Client.Conn = redis.NewClient(&redis.Options{
		Addr:     config.GetEnv().RedisIp + ":" + config.GetEnv().RedisPort,
		Password: config.GetEnv().RedisPassword,
		DB:       config.GetEnv().RedisDb,
	})
}

func (client *ClientType) Set(key string, value interface{}, expiration time.Duration) *redis.Client {
	err := client.Conn.Set(key, value, expiration).Err()
	if err != nil {
		panic(err)
	}
	return (*client).Conn
}

func (client *ClientType) Incr(key string) int64 {
	res, err := client.Conn.Incr(key).Result()
	if err != nil {
		panic(err)
	}
	return res
}

func (client *ClientType) Pipeline() redis.Pipeliner {
	return client.Conn.Pipeline()
}

func (client *ClientType) Decr(key string) int64 {
	res, err := client.Conn.Decr(key).Result()
	if err != nil {
		panic(err)
	}
	return res
}

func (client *ClientType) DecrBy(key string, de int64) int64 {
	res, err := client.Conn.DecrBy(key, de).Result()
	if err != nil {
		panic(err)
	}
	return res
}

func (client *ClientType) Expire(key string, expiration time.Duration) bool {
	res, err := client.Conn.Expire(key, expiration).Result()
	if err != nil {
		panic(err)
	}
	return res
}

func (client *ClientType) Get(key string) (string, error) {
	val, err := (*client).Conn.Get(key).Result()

	if err == redis.Nil {
		return "", nil
	}

	if err != nil {
		return "", err
	}

	return val, nil
}

func (client *ClientType) IsExist(key string) bool {
	_, err := (*client).Conn.Get(key).Result()
	return err != redis.Nil
}

func (client *ClientType) Lpop(key string) (string, error) {
	val, err := (*client).Conn.LPop(key).Result()

	if err == redis.Nil {
		return "", nil
	}

	if err != nil {
		return "", err
	}

	return val, nil
}

func (client *ClientType) Lpush(key string, values ...interface{}) bool {
	_, err := (*client).Conn.LPush(key, values...).Result()

	return err == nil
}

func (client *ClientType) Lrange(key string, start, stop int64) ([]string, error) {
	res, err := (*client).Conn.LRange(key, start, stop).Result()

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (client *ClientType) Del(key ...string) *redis.Client {
	_, err := client.Conn.Del(key...).Result()
	if err != nil {
		panic(err)
	}
	return (*client).Conn
}

func (client *ClientType) SetIfNotExist(key string, value interface{}, expiration time.Duration) bool {
	result, err := (*client).Conn.SetNX(key, value, expiration).Result()
	if err != nil {
		return false
	}
	return result
}

func (client *ClientType) PSubscribe(channels ...string) *redis.PubSub {
	ps := (*client).Conn.PSubscribe(channels...)
	return ps
}
