package main

import (
	"moringo/config"
	"moringo/controllers"
	//"moringo/filters"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/pprof"
)

func initRouter() *gin.Engine {
	router := gin.New()

	if config.GetEnv().DEBUG {
		router.Use(gin.Logger()) // 开发模式下使用，console打印请求记录
		pprof.Register(router) // 性能分析工具
	}

	router.Use(handleErrors()) // 错误处理

	// router.Use(filters.AuthMiddleware()) // 中间件使用

	api := router.Group("/api")
	{
		api.POST("/index", controllers.IndexApi)
		api.GET("/index", controllers.IndexApi)
	}

	return router
}
