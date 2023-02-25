package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type IdeasModel struct {
	ID         primitive.ObjectID `bson:"_id" json:"_id"`
	Image      string             `bson:"image" json:"image"`
	Title      string             `bson:"title" json:"title"`
	Slug       string             `bson:"slug" json:"slug"`
	Filename   string             `bson:"file_name" json:"file_name"`
	Department string             `bson:"department" json:"department"`
	Content    string             `bson:"content" json:"content"`
	Category   string             `bson:"category" json:"category"`
	Views      int                `bson:"views" json:"views"`
	UpVote     int                `bson:"up_vote" json:"up_vote"`
	DownVote   int                `bson:"down_vote" json:"down_vote"`
	CreatedAt  time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt  time.Time          `bson:"updated_at" json:"updated_at"`
	UserId     string             `bson:"user_id" json:"user_id"`
}
