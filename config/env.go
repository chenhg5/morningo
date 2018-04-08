package config

// 环境配置文件
// 可配置多个环境配置，进行切换

type Env struct {
	DEBUG             bool
	DATABASE_IP       string
	DATABASE_PORT     string
	DATABASE_USERNAME string
	DATABASE_PASSWORD string
	DATABASE_NAME     string
	SERVER_PORT       string
	REDIS_IP          string
	REDIS_PORT        string
	REDIS_PASSWORD    string
	REDIS_DB          int
	REDIS_SESSION_DB  int
	REDIS_CACHE_DB    int
	ACCESS_LOG        bool
}

var env = Env{
	DEBUG: true,

	SERVER_PORT:       "4000",
	DATABASE_IP:       "127.0.0.1",
	DATABASE_PORT:     "3306",
	DATABASE_USERNAME: "root",
	DATABASE_PASSWORD: "root",
	DATABASE_NAME:     "gin-template",

	REDIS_IP:       "127.0.0.1",
	REDIS_PORT:     "6379",
	REDIS_PASSWORD: "",
	REDIS_DB:       0,

	REDIS_SESSION_DB: 1,
	REDIS_CACHE_DB:   2,

	ACCESS_LOG: true,
}

func GetEnv() *Env {
	return &env
}
