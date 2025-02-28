package user

import (
	"fmt"

	"github.com/google/uuid"
)

type User struct {
	ID          uuid.UUID `json:"id" gorm:"primaryKey"`
	Username    string    `json:"username"`
	Email       string    `json:"email"`
	DisplayName string    `json:"display_name"`
}

func NewUser(email, displayName string) *User {
	return &User{
		ID:          uuid.New(),
		Username:    fmt.Sprintf("user-%s", uuid.New().String()),
		Email:       email,
		DisplayName: displayName,
	}
}
