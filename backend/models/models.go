package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Post struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title       string             `json:"title,omitempty" bson:"title,omitempty"`
	Description string             `json:"description,omitempty" bson:"description,omitempty"`
	Text        string             `json:"text,omitempty" bson:"text,omitempty"`
	CreatedAt   time.Time          `json:"created_at" bson:"created_at,omitempty"`
}

type PostAPI struct {
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	Text        string `json:"text,omitempty"`
	CreatedAt   string `json:"created_at"`
}

func (p PostAPI) GeneratePostApi(post Post) PostAPI {
	return PostAPI{
		Title:       post.Title,
		Description: post.Description,
		Text:        post.Text,
		CreatedAt:   post.CreatedAt.Format("02.01.06 15:04"),
	}
}
