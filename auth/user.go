package auth

import (
	"github.com/google/uuid"
)

type User struct {
	ID          uuid.UUID `json:"id" gorm:"primaryKey"`
	Username    string    `json:"username" gorm:"unique;not null"`
	Email       string    `json:"email"`
	DisplayName string    `json:"display_name"`
}

func NewUser(email, displayName string) *User {
	return &User{
		ID:          uuid.New(),
		Username:    email,
		Email:       email,
		DisplayName: displayName,
	}
}
