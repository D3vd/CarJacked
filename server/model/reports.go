package reports

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Reports struct {
	ID   primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Test int                `json:"test,omitempty"`
}
