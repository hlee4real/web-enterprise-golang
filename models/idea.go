package models

import "time"

type IdeasModel struct {
	Image     string    `bson:"image" json:"image"`
	Title     string    `bson:"title" json:"title"`
	Slug      string    `bson:"slug" json:"slug"`
	Filename  string    `bson:"file_name" json:"file_name"`
	Content   string    `bson:"content" json:"content"`
	Category  string    `bson:"category" json:"category"`
	Views     int       `bson:"views" json:"views"`
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
	UserId    string    `bson:"user_id" json:"user_id"`
	Comment   string    `bson:"comment" json:"comment"`
}
