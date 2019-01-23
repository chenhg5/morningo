package config

import "github.com/go-sql-driver/mysql"

// 本文件建议在代码协同工具(git/svn等)中忽略

var env = Env{
	DEBUG: true,

	SERVER_PORT: "4000",

	DATABASE: mysql.Config{
		User:                 "root",
		Passwd:               "root",
		Addr:                 "127.0.0.1:3306",
		DBName:               "gin-template",
		Collation:            "utf8mb4_unicode_ci",
		Net:                  "tcp",
		AllowNativePasswords: true,
	},
	MAXIDLECONNS: 50,
	MAXOPENCONNS: 100,

	REDIS_IP:       "127.0.0.1",
	REDIS_PORT:     "6379",
	REDIS_PASSWORD: "",
	REDIS_DB:       0,

	REDIS_SESSION_DB: 1,
	REDIS_CACHE_DB:   2,

	ACCESS_LOG:      true,
	ACCESS_LOG_PATH: "storage/logs/access.log",

	ERROR_LOG:      true,
	ERROR_LOG_PATH: "storage/logs/error.log",

	INFO_LOG:      true,
	INFO_LOG_PATH: "storage/logs/info.log",

	TEMPLATE_PATH: "frontend/templates",

	//APP_SECRET: "YbskZqLNT6TEVLUA9HWdnHmZErypNJpL",
	APP_SECRET: "something-very-secret",
}