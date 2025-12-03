package routes

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"gin-sass-salon/app/http/controllers"
	"gin-sass-salon/app/http/middleware"
)

// SetupRoutes mendefinisikan semua route API
func SetupRoutes(r *gin.Engine) {
	// Swagger route
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := r.Group("/api")
	{
		// Public routes (tidak perlu authentication)
		auth := api.Group("/auth")
		{
			auth.POST("/register", controllers.Register)
			auth.POST("/login", controllers.Login)
		}

		// Protected routes (perlu authentication dengan JWT)
		protected := api.Group("")
		protected.Use(middleware.AuthMiddleware())
		{
			// User CRUD routes
			protected.GET("/users", controllers.GetUsers)
			protected.GET("/users/:id", controllers.GetUser)
			protected.POST("/users", controllers.CreateUser)
			protected.PUT("/users/:id", controllers.UpdateUser)
			protected.DELETE("/users/:id", controllers.DeleteUser)
		}
	}
}