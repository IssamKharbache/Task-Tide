package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Todo struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Completed bool               `json:"completed"`
	Body      string             `json:"body"`
	UserID    string             `json:"userId"`
	DueDate   time.Time          `json:"dueDate"`
	CreatedAt time.Time          `json:"createdAt"`
	UpdatedAt time.Time          `json:"updateAt"`
}
