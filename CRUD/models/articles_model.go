package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Article represents the structure of an article document in the database.
type Article struct {
	ID        primitive.ObjectID `json:"id" `
	Title     string             `json:"title" `
	Content   string             `json:"content"`
	Author    string             `json:"author"`
	Image     string             `json:"image" `
	CreatedAt time.Time          `json:"created_at" `
	UpdatedAt time.Time          `json:"updated_at" `
}

// NewArticle creates a new instance of the Article struct.
func NewArticle(title, content, author, image string) *Article {
	return &Article{
		Title:     title,
		Content:   content,
		Author:    author,
		Image:     image,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
