package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/mauFade/study-hub/internal/infra/repository"
	"github.com/mauFade/study-hub/internal/loaders"
	"github.com/mauFade/study-hub/internal/usecase"
)

type CreateUserInputDTO struct {
	Name       string `json:"name"`
	Email      string `json:"email"`
	University string `json:"university"`
}

func CreateUserController(context *gin.Context) {
	var body CreateUserInputDTO

	context.Bind(&body)

	fmt.Print(body)

	repo := repository.NewRepository(loaders.DB)
	user := usecase.NewCreateUserUseCase(*repo)

	output, err := user.Execute(usecase.CreateUserInputDTO{Name: body.Name, Email: body.Email, University: body.University})

	if err != nil {
		// context.Status(400)
		fmt.Print(err)
		context.JSON(400, err)
		return
	}

	context.JSON(201, output)
}
