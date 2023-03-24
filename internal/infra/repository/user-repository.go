package repository

import (
	"github.com/mauFade/study-hub/internal/model"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewRepository(database *gorm.DB) *UserRepository {
	return &UserRepository{DB: database}
}

func (repo *UserRepository) Create(user *model.User) *gorm.DB {
	result := repo.DB.Create(&user)

	return result
}
