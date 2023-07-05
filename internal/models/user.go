package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	Username  string             `json:"username" bson:"username"`
	Email     string             `json:"email" bson:"email"`
	Age       int16              `json:"age" bson:"age"`
	FullName  string             `json:"fullName" bson:"fullName"`
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt *time.Time         `json:"updatedAt" bson:"updatedAt,omitempty"`
	DeletedAt *time.Time         `json:"deletedAt" bson:"deletedAt,omitempty"`
}

func (u User) UpdateFields() bson.D {
	values := bson.D{}

	if u.Username != "" {
		values = append(values, bson.E{
			Key:   "username",
			Value: u.Username,
		})
	}

	if u.Email != "" {
		values = append(values, bson.E{
			Key:   "email",
			Value: u.Email,
		})
	}

	if u.Age != 0 {
		values = append(values, bson.E{
			Key:   "age",
			Value: u.Age,
		})
	}

	if u.FullName != "" {
		values = append(values, bson.E{
			Key:   "fullName",
			Value: u.FullName,
		})
	}

	now := time.Now()
	if u.UpdatedAt == nil {
		u.UpdatedAt = &now
	}

	values = append(values, bson.E{
		Key:   "updatedAt",
		Value: u.UpdatedAt,
	})

	return values
}
