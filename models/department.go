package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type DepartmentsModel struct {
	ID        primitive.ObjectID `bson:"_id" json:"_id"`
	Name      string             `bson:"name" json:"name"`
	Column    int                `bson:"column" json:"column"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
}
