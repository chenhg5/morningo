package main

import (
	"morningo/config"
	"morningo/controllers"
	//"morningo/filters"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cache"
	"morningo/filters/auth"
	"morningo/filters/auth/drivers"
	"github.com/gin-contrib/cache/persistence"
	"time"
)

func initRouter() *gin.Engine {
	router := gin.New()

	if config.GetEnv().DEBUG {
		router.Use(gin.Logger()) // 开发模式下使用，console打印请求记录
		pprof.Register(router)   // 性能分析工具
	}

	router.Use(handleErrors()) // 错误处理

	store, _ := sessions.NewRedisStore(10, "tcp", config.GetEnv().REDIS_IP+":"+config.GetEnv().REDIS_PORT, config.GetEnv().REDIS_PASSWORD, []byte("secret"))
	router.Use(sessions.Sessions("mysession", store)) // 全局session

	var cacheStore persistence.CacheStore
	cacheStore = persistence.NewRedisCache(config.GetEnv().REDIS_IP+":"+config.GetEnv().REDIS_PORT, config.GetEnv().REDIS_PASSWORD, time.Minute)
	router.Use(cache.Cache(&cacheStore))

	//cookie := sessions.NewCookieStore([]byte("12323"))
	//router.Use(sessions.Sessions("mysession", cookie))

	var authDriver auth.Auth
	authDriver = drivers.NewCacheAuthDriver()
	router.Use(auth.AuthSetMiddleware(&authDriver, "web_auth"))

	var authJwtDriver auth.Auth
	authJwtDriver = drivers.NewJwtAuthDriver()
	router.Use(auth.AuthSetMiddleware(&authJwtDriver, "jwt_auth"))

	router.LoadHTMLGlob("frontend/templates/*") // html模板

	// router.Use(filters.AuthMiddleware()) // 中间件使用

	api := router.Group("/api")
	api.GET("/index", controllers.IndexApi)
	api.GET("/cookie/set/:userid", controllers.CookieSetExample)
	api.Use(auth.AuthMiddleware(&authDriver))
	{
		api.GET("/orm", controllers.OrmExample)
		api.GET("/store", controllers.StoreExample)
		api.GET("/db", controllers.DBexample)
		api.GET("/cookie/get", controllers.CookieGetExample)
	}
	jwtApi := router.Group("/api")
	jwtApi.GET("/jwt/set/:userid", controllers.JwtSetExample)
	jwtApi.Use(auth.AuthMiddleware(&authJwtDriver))
	{
		jwtApi.GET("/jwt/get", controllers.JwtGetExample)
	}

	return router
}
