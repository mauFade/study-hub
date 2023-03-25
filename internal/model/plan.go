package model

import "github.com/google/uuid"

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
