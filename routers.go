package main

import (
	"morningo/config"
	"morningo/controllers"
	"github.com/gin-contrib/cache"
	"github.com/gin-contrib/cache/persistence"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"morningo/filters/auth"
	"morningo/filters/auth/drivers"
	"morningo/routes"
	"time"
	// proxy "github.com/chenhg5/gin-reverseproxy"
)

func initRouter() *gin.Engine {
	router := gin.New()

	if config.GetEnv().DEBUG {
		router.Use(gin.Logger()) // 开发模式下使用，console打印请求记录
		pprof.Register(router)   // 性能分析工具
	}

	router.Use(handleErrors()) // 错误处理

	// 全局session
	store, _ := sessions.NewRedisStore(10, "tcp", config.GetEnv().REDIS_IP+":"+config.GetEnv().REDIS_PORT, config.GetEnv().REDIS_PASSWORD, []byte("secret"))
	router.Use(sessions.Sessions("mysession", store))

	// 全局cache
	var cacheStore persistence.CacheStore
	cacheStore = persistence.NewRedisCache(config.GetEnv().REDIS_IP+":"+config.GetEnv().REDIS_PORT, config.GetEnv().REDIS_PASSWORD, time.Minute)
	router.Use(cache.Cache(&cacheStore))

	// 全局auth cookie
	var authCookieDriver auth.Auth
	authCookieDriver = drivers.NewCookieAuthDriver()
	router.Use(auth.AuthSetMiddleware(&authCookieDriver, "web_auth"))

	// 全局auth jwt
	var authJwtDriver auth.Auth
	authJwtDriver = drivers.NewJwtAuthDriver()
	router.Use(auth.AuthSetMiddleware(&authJwtDriver, "jwt_auth"))

	router.LoadHTMLGlob(config.GetEnv().TEMPLATE_PATH + "/*") // html模板

	// ReverseProxy
	// router.Use(proxy.ReverseProxy(map[string] string {
	// 	"localhost:4000" : "localhost:9090",
	// }))

	api := router.Group("/api")
	api.GET("/index", controllers.IndexApi)
	api.GET("/cookie/set/:userid", controllers.CookieSetExample)

	// cookie auth middleware
	api.Use(auth.AuthMiddleware(&authCookieDriver))
	{
		api.GET("/orm", controllers.OrmExample)
		api.GET("/store", controllers.StoreExample)
		api.GET("/db", controllers.DBexample)
		api.GET("/cookie/get", controllers.CookieGetExample)
	}

	jwtApi := router.Group("/api")
	jwtApi.GET("/jwt/set/:userid", controllers.JwtSetExample)

	// jwt auth middleware
	jwtApi.Use(auth.AuthMiddleware(&authJwtDriver))
	{
		jwtApi.GET("/jwt/get", controllers.JwtGetExample)
	}

	routes.RegisterApiRouter(router)

	return router
}
