package routes

import (
	"github.com/gin-gonic/gin"
	"morningo/controllers"
	"morningo/filters/auth"
)

func RegisterApiRouter(router *gin.Engine) {
	apiRouter := router.Group("api")
	{
		apiRouter.GET("/test/index", controllers.IndexApi)
	}

	api := router.Group("/api")
	api.GET("/index", controllers.IndexApi)
	api.GET("/cookie/set/:userid", controllers.CookieSetExample)

	// cookie auth middleware
	api.Use(auth.Middleware("cookie"))
	{
		api.GET("/orm", controllers.OrmExample)
		api.GET("/store", controllers.StoreExample)
		api.GET("/db", controllers.DBexample)
		api.GET("/cookie/get", controllers.CookieGetExample)
	}

	jwtApi := router.Group("/api")
	jwtApi.GET("/jwt/set/:userid", controllers.JwtSetExample)

	// jwt auth middleware
	jwtApi.Use(auth.Middleware("jwt"))
	{
		jwtApi.GET("/jwt/get", controllers.JwtGetExample)
	}
}
