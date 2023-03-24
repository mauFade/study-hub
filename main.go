package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mauFade/study-hub/internal/loaders"
)

func init() {
	loaders.GetEnvironmentVariables()
	loaders.ConnectToDatabase()
}

func main() {
	router := gin.Default()

	router.POST("/users")

	router.Run()
}
