package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/takadev15/kanban-api/config"
	"github.com/takadev15/kanban-api/controller"
)

func RoutesList() *gin.Engine {
  r := gin.Default()
  db := config.GetDB()
  handler := controller.Handlers{Connect: db}

  userRoutes := r.Group("/user")
  {
    userRoutes.POST("/login", handler.UserLogin)
    userRoutes.POST("/register", handler.UserRegister)
    userRoutes.POST("/update-account")
    userRoutes.POST("/delete-account")
  }

  categoriesRoutes := r.Group("/categories")
  {
    categoriesRoutes.GET("/")
    categoriesRoutes.POST("/")
    categoriesRoutes.PATCH("/:id")
    categoriesRoutes.DELETE("/:id")
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
