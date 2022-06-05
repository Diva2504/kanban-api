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
  return nil
}
