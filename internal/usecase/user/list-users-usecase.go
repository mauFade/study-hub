package usecase

import (
	"time"

	"github.com/mauFade/study-hub/internal/infra/repository"
	"github.com/mauFade/study-hub/internal/model"
)

type FindUserOutputDTO struct {
	ID         string
	Name       string
	Email      string
	University string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type FindUserUseCase struct {
	UserRepository repository.UserRepository
}

func NewListUserUseCase(repo repository.UserRepository) *FindUserUseCase {
	return &FindUserUseCase{
		UserRepository: repo,
	}
}

func (findUseCase *FindUserUseCase) Execute() []model.User {
	data := findUseCase.UserRepository.Find()

	return data
}
