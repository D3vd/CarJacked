package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	OfficerID primitive.ObjectID `json:"officerID,omitempty" bson:"officerID,omitempty"`
	Username  string             `json:"username"`
	Password  string             `json:"password"`
}
