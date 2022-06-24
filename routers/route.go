package routers

import "github.com/gin-gonic/gin"

func RoutesList() *gin.Engine {
	r := gin.Default()
	userRoutes := r.Group("/user")
	{
		// get all users data
		userRoutes.GET("/")
		// get specific user data
		userRoutes.GET("/:id")
	}
	categoryRoutes := r.Group("/categories")
	{
		categoryRoutes.GET("/")
		categoryRoutes.POST("/")
		categoryRoutes.PATCH("/:id")
		categoryRoutes.DELETE("/:id")
	}
	taskRoutes := r.Group("/tasks")
	{
		taskRoutes.GET("/")
		taskRoutes.POST("/")
		taskRoutes.PUT("/:id")
		taskRoutes.PATCH("/:id")
		taskRoutes.PATCH("/update-category/:id")
		taskRoutes.DELETE("/:id")
	}
	return nil
}
