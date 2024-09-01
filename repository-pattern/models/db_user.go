package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	UserName  string             `json:"username,omitempty"`
	FisrtName string             `json:"first_name,omitempty"`
	LastName  string             `json:"last_name,omitempty"`
	Email     string             `json:"email,omitempty"`
}
