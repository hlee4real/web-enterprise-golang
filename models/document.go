package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type DocumentsModel struct {
	ID        primitive.ObjectID `bson:"_id" json:"_id"`
	Filename  string             `bson:"file_name" json:"file_name"`
	CreatedBy string             `bson:"created_by" json:"created_by"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
	UserId    string             `bson:"user_id" json:"user_id"`
	IdeaId    string             `bson:"idea_id" json:"idea_id"`
}
