package routes

import (
	"github.com/gin-api-template/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	// Create controllers
	userController := controllers.NewUserController()

	// API v1 routes
	v1 := r.Group("/api/v1")
	{
		// User routes
		users := v1.Group("/users")
		{
			users.GET("/", userController.GetUsers)
			users.GET("/:id", userController.GetUser)
			users.POST("/", userController.CreateUser)
			users.PUT("/:id", userController.UpdateUser)
			users.DELETE("/:id", userController.DeleteUser)
		}
	}
} 