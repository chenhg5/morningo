package config

import "github.com/go-sql-driver/mysql"

// 本文件建议在代码协同工具(git/svn等)中忽略

var env = Env{
	Debug: true,

	ServerPort: "4000",

	Database: mysql.Config{
		User:                 "root",
		Passwd:               "root",
		Addr:                 "127.0.0.1:3306",
		DBName:               "gin-template",
		Collation:            "utf8mb4_unicode_ci",
		Net:                  "tcp",
		AllowNativePasswords: true,
	},
	MaxIdleConns: 50,
	MaxOpenConns: 100,

	RedisIp:       "127.0.0.1",
	RedisPort:     "6379",
	RedisPassword: "",
	RedisDb:       0,

	RedisSessionDb: 1,
	RedisCacheDb:   2,

	AccessLog:     true,
	AccessLogPath: "storage/logs/access.log",

	ErrorLog:     true,
	ErrorLogPath: "storage/logs/error.log",

	InfoLog:     true,
	InfoLogPath: "storage/logs/info.log",

	TemplatePath: "frontend/templates",

	//APP_SECRET: "YbskZqLNT6TEVLUA9HWdnHmZErypNJpL",
	AppSecret: "something-very-secret",
}
