package models

import (
	"github.com/google/uuid"
	"github.com/kubeblog/backend/auth"
)

type Article struct {
	ID       uuid.UUID `json:"id" gorm:"type:uuid;primary_key"`
	Title    string    `json:"title"`
	Content  string    `json:"content"`
	Image    string    `json:"image"`
	Category Category  `json:"category"`
	Excerpt  string    `json:"excerpt"`
	Tags     string    `json:"tags"`

	AuthorID uuid.UUID `json:"authorId" gorm:"type:uuid;not null"`
	Author   auth.User `json:"author" gorm:"foreignKey:AuthorID"`
	Featured bool      `json:"featured"`

	// Timestamps
	CreatedAt int64 `json:"createdAt"`
	UpdatedAt int64 `json:"updatedAt"`
}

type Category string

const (
	CategoryFrontend Category = "Frontend"
	CategoryBackend  Category = "Backend"
	CategoryDevOps   Category = "DevOps"
	CategoryCloud    Category = "Cloud"
	CategoryDatabase Category = "Database"
	CategoryGeneral  Category = "General"
	CategoryRandom   Category = "Random"
)

var CATEGORIES = []Category{
	CategoryFrontend,
	CategoryBackend,
	CategoryDevOps,
	CategoryCloud,
	CategoryDatabase,
	CategoryGeneral,
	CategoryRandom,
}
