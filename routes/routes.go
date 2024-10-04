package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/ARMeeru/time-capsule/controllers"
	"github.com/ARMeeru/time-capsule/middlewares"
)

func InitRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		// Public routes
		api.POST("/register", controllers.Register)
		api.POST("/login", controllers.Login)

		// Protected routes
		api.Use(middlewares.AuthMiddleware())
		{
			api.POST("/capsules", controllers.CreateCapsule)
		}
	}
}
