package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mauFade/study-hub/internal/infra/controllers"
	"github.com/mauFade/study-hub/internal/loaders"
)

func init() {
	loaders.GetEnvironmentVariables()
	loaders.ConnectToDatabase()
}

func main() {
	router := gin.Default()

	router.POST("/users", controllers.CreateUserController)
	router.GET("/users", controllers.ListUsersController)

	router.Run()
}
