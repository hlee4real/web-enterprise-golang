package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type UsersModel struct {
	ID           primitive.ObjectID `bson:"_id" json:"_id"`
	Username     *string            `bson:"username" json:"username"`
	FirstName    *string            `bson:"first_name" json:"first_name"`
	LastName     *string            `bson:"last_name" json:"last_name"`
	Password     *string            `bson:"password" json:"password"`
	Token        *string            `bson:"token" json:"token"`
	RefreshToken *string            `bson:"refresh_token" json:"refresh_token"`
	Mobile       string             `bson:"mobile" json:"mobile"`
	Role         string             `bson:"role" json:"role"`
	Image        string             `bson:"image" json:"image"`
	Department   string             `bson:"department" json:"department"`
	CreatedAt    time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt    time.Time          `bson:"updated_at" json:"updated_at"`
	UserID       string             `bson:"user_id" json:"user_id"`
}
