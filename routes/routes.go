package routes

import (
	"github.com/ARMeeru/time-capsule/controllers"
	"github.com/ARMeeru/time-capsule/middlewares"
	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.Engine) {
	// Public routes
	router.GET("/register", controllers.ShowRegisterPage)
	router.POST("/register", controllers.Register)
	router.GET("/login", controllers.ShowLoginPage)
	router.POST("/login", controllers.Login)

	// Protected routes
	authorized := router.Group("/")
	authorized.Use(middlewares.AuthMiddleware())

	authorized.GET("/capsules/new", controllers.ShowCreateCapsulePage)
	authorized.POST("/capsules", controllers.CreateCapsule)
	authorized.GET("/logout", controllers.Logout)
}
