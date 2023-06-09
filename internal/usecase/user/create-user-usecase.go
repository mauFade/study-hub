package usecase

import (
	"time"

	"github.com/mauFade/study-hub/internal/infra/repository"
	"github.com/mauFade/study-hub/internal/model"
)

type CreateUserInputDTO struct {
	Name       string
	Email      string
	University string
	Password   string
}

type CreateUserOutputDTO struct {
	ID         string
	Name       string
	Email      string
	University string
	Password   string
	PlanId     string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type CreateUserUseCase struct {
	UserRepository repository.UserRepository
}

func NewCreateUserUseCase(repository repository.UserRepository) *CreateUserUseCase {
	return &CreateUserUseCase{
		UserRepository: repository,
	}
}

func (createUseCase *CreateUserUseCase) Execute(data CreateUserInputDTO) (*CreateUserOutputDTO, error) {
	user := model.NewUser(data.Name, data.Email, data.University, data.Password, "")

	err := user.Validate()

	if err != nil {
		return nil, err
	}

	createUseCase.UserRepository.Create(user)

	return &CreateUserOutputDTO{
		ID:         user.ID,
		Name:       user.Name,
		Email:      user.Email,
		University: user.University,
		Password:   user.Password,
		PlanId:     user.PlanID,
		CreatedAt:  user.CreatedAt,
		UpdatedAt:  user.UpdatedAt,
	}, nil
}
