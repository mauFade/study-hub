package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/mauFade/study-hub/internal/infra/repository"
	"github.com/mauFade/study-hub/internal/loaders"
	"github.com/mauFade/study-hub/internal/usecase"
)

func ListUsersController(context *gin.Context) {
	repo := repository.NewRepository(loaders.DB)
	users := usecase.NewListUserUseCase(*repo)

	output := users.Execute()

	fmt.Print(output)

	context.JSON(200, output)
}
