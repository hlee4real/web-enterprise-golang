package models

import "time"

type CommentsModel struct {
	UserId    string    `bson:"user_id" json:"user_id"`
	IdeaId    string    `bson:"idea_id" json:"idea_id"`
	Comment   string    `bson:"comment" json:"comment"`
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
}
