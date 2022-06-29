package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/takadev15/kanban-api/config"
	"github.com/takadev15/kanban-api/controller"
	"github.com/takadev15/kanban-api/middleware"
)

func RoutesList() *gin.Engine {
	r := gin.Default()
  db := config.GetDB()
  handler := controller.Handlers{Connect: db}
	userRoutes := r.Group("/user")
	{
		// get all users data
		userRoutes.POST("/login", handler.UserLogin)
		userRoutes.POST("/register", handler.UserRegister)
		userRoutes.PUT("/update-account", middleware.Authentication(), handler.UpdateUser)
		userRoutes.DELETE("/delete-account", middleware.Authentication(), handler.DeleteUser)
	}
	categoryRoutes := r.Group("/categories")
  categoryRoutes.Use(middleware.Authentication(), middleware.AdminAuth())
	{
		categoryRoutes.GET("/", handler.GetAllCategory)
		categoryRoutes.POST("/", handler.CreateCategory)
		categoryRoutes.PATCH("/:id", handler.UpdateCategory)
		categoryRoutes.DELETE("/:id", handler.DeleteCtegory)
	}
	taskRoutes := r.Group("/tasks")
  taskRoutes.Use(middleware.Authentication())
	{
		taskRoutes.GET("/", handler.GetAllTask)
		taskRoutes.POST("/", handler.CreateTask)
		taskRoutes.PUT("/:id", handler.UpdateTask)
		taskRoutes.PATCH("/:id", handler.UpdateStatusTask)
    taskRoutes.PATCH("/update-category/:id", handler.UpdateCategoryTask)
		taskRoutes.DELETE("/:id", handler.DeleteTask)
	}
	return r
}
