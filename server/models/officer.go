package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Officer struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	UserID   primitive.ObjectID `json:"userID,omitempty" bson:"userID,omitempty"`
	Name     string             `json:"name"`
	Assigned bool               `json:"assigned"`
}
