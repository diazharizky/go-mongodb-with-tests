package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	Username  string             `json:"username" bson:"username"`
	Email     string             `json:"email" bson:"email"`
	Age       int16              `json:"age" bson:"age"`
	FullName  string             `json:"fullName" bson:"fullName"`
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt *time.Time         `json:"updatedAt" bson:"updatedAt"`
	DeletedAt *time.Time         `json:"deletedAt" bson:"deletedAt"`
}
