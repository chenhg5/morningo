package main

import (
	"morningo/config"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"morningo/filters/auth"
	"morningo/filters"
	routeRegister "morningo/routes"
	// proxy "github.com/chenhg5/gin-reverseproxy"
)

func initRouter() *gin.Engine {
	router := gin.New()

	router.LoadHTMLGlob(config.GetEnv().TEMPLATE_PATH + "/*") // html模板

	if config.GetEnv().DEBUG {
		router.Use(gin.Logger()) // 开发模式下使用，console打印请求记录
		pprof.Register(router)   // 性能分析工具
	}

	router.Use(handleErrors()) // 错误处理
	router.Use(filters.RegisterSession())  // 全局session
	router.Use(filters.RegisterCache())    // 全局cache

	router.Use(auth.RegisterGlobalAuthDriver("cookie", "web_auth")) // 全局auth cookie
	router.Use(auth.RegisterGlobalAuthDriver("jwt", "jwt_auth"))    // 全局auth jwt

	routeRegister.RegisterApiRouter(router)

	// ReverseProxy
	// router.Use(proxy.ReverseProxy(map[string] string {
	// 	"localhost:4000" : "localhost:9090",
	// }))

	return router
}
