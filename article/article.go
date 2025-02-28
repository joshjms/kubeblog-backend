package article

import "github.com/google/uuid"

type Article struct {
	ID       uuid.UUID `json:"id"`
	Title    string    `json:"title"`
	Content  string    `json:"content"`
	Image    string    `json:"image"`
	Category Category  `json:"category"`
	Excerpt  string    `json:"excerpt"`
	Tags     string    `json:"tags"`

	Author   string `json:"author"`
	Featured bool   `json:"featured"`

	// Timestamps
	CreatedAt int64 `json:"created_at"`
	UpdatedAt int64 `json:"updated_at"`
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
