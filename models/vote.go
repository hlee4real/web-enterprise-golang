package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type UserVoteModels struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	UpVote    bool               `bson:"up_vote" json:"up_vote"`
	DownVote  bool               `bson:"down_vote" json:"down_vote"`
	Username  string             `bson:"username" json:"username"`
	IdeaId    string             `bson:"idea_id" json:"idea_id"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
}
