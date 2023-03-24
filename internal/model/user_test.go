package model

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestIfIdIsBlank(t *testing.T) {
	user := &User{}

	assert.Error(t, user.Validate(), "User id is required")
}

func TestIfNameIsBlank(t *testing.T) {
	user := &User{}

	user.ID = "ID"

	assert.Error(t, user.Validate(), "User name is required")
}

func TestIfEmailIsBlank(t *testing.T) {
	user := &User{}

	user.ID = "ID"
	user.Name = "Name"

	assert.Error(t, user.Validate(), "User email is required")
}

func TestIfUniversityIsBlank(t *testing.T) {
	user := &User{}

	user.ID = "ID"
	user.Name = "Name"
	user.Email = "email@user.com"

	assert.Error(t, user.Validate(), "User university is required")
}

func TestUserWithValidParams(t *testing.T) {
	user := &User{
		ID:         "ID",
		Name:       "Name",
		Email:      "email@user.com",
		University: "MIT",
		CreatedAt:  time.Now(),
	}

	assert.Equal(t, "email@user.com", user.Email)
	assert.Equal(t, "MIT", user.University)
}
