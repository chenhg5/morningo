package main

import (
	"morningo/config"
	"morningo/controllers"
	//"morningo/filters"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

func initRouter() *gin.Engine {
	router := gin.New()

	if config.GetEnv().DEBUG {
		router.Use(gin.Logger()) // 开发模式下使用，console打印请求记录
		pprof.Register(router)   // 性能分析工具
	}

	router.Use(handleErrors()) // 错误处理

	router.LoadHTMLGlob("frontend/templates/*") // html模板

	// router.Use(filters.AuthMiddleware()) // 中间件使用

	api := router.Group("/api")
	{
		api.GET("/orm", controllers.OrmExample)
		api.GET("/store", controllers.StoreExample)
		api.GET("/db", controllers.DBexample)
		api.GET("/index", controllers.IndexApi)
	}

	return router
}
