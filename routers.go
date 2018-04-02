package main

import (
	"./config"
	"./controllers"
	"github.com/gin-gonic/gin"
)

func initRouter() *gin.Engine {
	router := gin.New()

	if config.GetEnv().DEBUG {
		router.Use(gin.Logger()) // 开发模式下使用，console打印请求记录
	}

	router.Use(handleErrors()) // 错误处理

	api := router.Group("/api")
	{
		api.POST("/", controllers.IndexApi)
		api.GET("/", controllers.IndexApi)
	}

	return router
}
