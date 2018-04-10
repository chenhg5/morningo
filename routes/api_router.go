package routes

import (
	"github.com/gin-gonic/gin"
	"morningo/controllers"
)

func RegisterApiRouter(r *gin.Engine) {
	apiRouter := r.Group("api")
	{
		apiRouter.GET("/test/index", controllers.IndexApi)
	}
}
