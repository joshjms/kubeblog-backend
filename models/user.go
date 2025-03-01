package models

import (
	"github.com/google/uuid"
)

type User struct {
	ID          uuid.UUID `json:"id" gorm:"primaryKey"`
	Username    string    `json:"username" gorm:"unique;not null"`
	Email       string    `json:"email"`
	DisplayName string    `json:"displayName"`

	Bio string `json:"bio"`

	Github  string `json:"github"`
	Twitter string `json:"twitter"`
	Website string `json:"website"`
}

func NewUser(email, displayName string) *User {
	return &User{
		ID:          uuid.New(),
		Username:    email,
		Email:       email,
		DisplayName: displayName,
	}
}
