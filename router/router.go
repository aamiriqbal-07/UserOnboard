package router

import (
	"userOnboard/controller"
	"userOnboard/repository"
	"userOnboard/service"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	repo := repository.NewUserRepository()
	service := service.NewUserService(repo)
	controller := controller.NewUserController(service)

	auth := r.Group("/api")
	auth.Use(basicAuthMiddleware())
	{
		auth.POST("/users", controller.CreateUser)
		auth.GET("/users/:id", controller.GetUserByID)
		auth.GET("/users", controller.ListUsers)
	}
}

func basicAuthMiddleware() gin.HandlerFunc {
	return gin.BasicAuth(gin.Accounts{
		"admin": "password", // Hard-coded credentials for testing
	})
}
