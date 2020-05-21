package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Auth struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Username string `json:"username"`
	Password string `json:"password"`
}
