package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/mauFade/study-hub/internal/infra/repository"
	"github.com/mauFade/study-hub/internal/loaders"
	usecase "github.com/mauFade/study-hub/internal/usecase/user"
)

func ListUsersController(context *gin.Context) {
	repo := repository.NewRepository(loaders.DB)
	users := usecase.NewListUserUseCase(*repo)

	output := users.Execute()

	context.JSON(200, output)
}
