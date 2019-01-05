package config

import "github.com/go-sql-driver/mysql"

// 环境配置文件
// 可配置多个环境配置，进行切换

type Env struct {
	DEBUG            bool
	DATABASE         mysql.Config
	MAXIDLECONNS     int
	MAXOPENCONNS     int
	SERVER_PORT      string
	REDIS_IP         string
	REDIS_PORT       string
	REDIS_PASSWORD   string
	REDIS_DB         int
	REDIS_SESSION_DB int
	REDIS_CACHE_DB   int
	APP_SECRET       string

	ACCESS_LOG      bool
	ACCESS_LOG_PATH string
	ERROR_LOG       bool
	ERROR_LOG_PATH  string
	INFO_LOG        bool
	INFO_LOG_PATH   string

	SQL_LOG bool

	TEMPLATE_PATH string // 静态文件相对路径
}

func GetEnv() *Env {
	return &env
}