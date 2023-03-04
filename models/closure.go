package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type ClosureModels struct {
	ID           primitive.ObjectID `bson:"_id" json:"_id"`
	FirstClosure time.Time          `bson:"first_closure" json:"first_closure"`
	FinalClosure time.Time          `bson:"final_closure" json:"final_closure"`
}
