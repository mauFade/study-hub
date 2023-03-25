package model

import (
	"errors"

	"github.com/google/uuid"
)

type Plan struct {
	ID         string
	Name       string
	StudentId  string
	Subject    string
	Difficulty string
	Importance string
}

func NewPlan(name string, student_id string, subject string, difficulty string, importance string) *Plan {
	return &Plan{
		ID:         uuid.NewString(),
		Name:       name,
		StudentId:  student_id,
		Subject:    subject,
		Difficulty: difficulty,
		Importance: importance,
	}
}

func (plan *Plan) Validate() error {
	if plan.ID == "" {
		return errors.New("Plan id is required")
	}

	if plan.Name == "" {
		return errors.New("Plan name is required")
	}

	if plan.StudentId == "" {
		return errors.New("Plan student is required")
	}

	if plan.Subject == "" {
		return errors.New("Plan Subject is required")
	}

	if plan.Importance == "" {
		return errors.New("Plan importance is requried")
	}

	if plan.Difficulty == "" {
		return errors.New("Plan difficulty is required")
	}

	return nil
}

type PlanRepository interface {
	Create(plan *Plan) error
}
