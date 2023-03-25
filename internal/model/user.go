package model

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID         string
	Name       string
	Email      string
	University string
	Password   string
	PlanID     string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func NewUser(name string, email string, university string, password string, plan_id string) *User {
	return &User{
		ID:         uuid.NewString(),
		Name:       name,
		Email:      email,
		University: university,
		Password:   password,
		PlanID:     plan_id,
		CreatedAt:  time.Now(),
	}
}

func (user *User) Validate() error {
	if user.ID == "" {
		return errors.New("User id is required")
	}

	if user.Name == "" {
		return errors.New("User name is required")
	}

	if user.Email == "" {
		return errors.New("User email is required")
	}

	if user.University == "" {
		return errors.New("User university is required")
	}

	if user.Password == "" {
		return errors.New("User password is required")
	}

	return nil
}

type UserRepository interface {
	Create(user *User) error
	Find() []User
	FindByID(user_id string) (User, error)
}
