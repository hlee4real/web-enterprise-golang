package models

import "time"

type DepartmentsModel struct {
	Name      string    `bson:"name" json:"name"`
	Column    int       `bson:"column" json:"column"`
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
}
