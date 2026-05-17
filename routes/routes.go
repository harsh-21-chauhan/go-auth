package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/harsh-21-chauhan/go-auth/controllers"
	"github.com/harsh-21-chauhan/go-auth/middleware"
)

func SetupRoutes(router *gin.Engine){
	router.POST("/signup", controllers.Signup)
	router.POST("/login", controllers.Login)

	protected := router.Group("/")

	protected.Use(middleware.Authenciated())

	{
		protected.GET("/users",controllers.GetAllUsers)
		protected.GET("/profile", controllers.GetProfile)
	}

}