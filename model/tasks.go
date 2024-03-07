package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type task struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Title       string             `bson:"title,omitempty" json:"title,omitempty"`
	Name        string             `bson:"name" json:"name"`
	Description string             `bson:"description" json:"description"`
	DueDate     string             `bson:"due_date" json:"due_date"`
	SubmittedAt time.Time          `bson:"submitted_at" json:"submitted_at"`
}
