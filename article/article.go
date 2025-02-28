package article

import "github.com/google/uuid"

type Article struct {
	ID       uuid.UUID `json:"id"`
	Title    string    `json:"title"`
	Content  string    `json:"content"`
	Image    string    `json:"image"`
	Category string    `json:"category"`
	Excerpt  string    `json:"excerpt"`
	Tags     string    `json:"tags"`
}
