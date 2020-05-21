package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Case : Details of the Case
type Case struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Officer  primitive.ObjectID `json:"officer,omitempty" bson:"officer,omitempty"`
	UserInfo     UserInfo               `json:"user"`
	Car      Car                `json:"car"`
	Active   bool               `json:"active"`
	Assigned bool               `json:"assigned"`
}

// User : Contact details of the User
type UserInfo struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

// Car : Details of the missing Car
type Car struct {
	Color string `json:"color"`
	RegNo string `json:"regNo"`
	// LastSeen time.Time `json:"lastSeen"`
	// Location Location `json:"location"`
}

// Location : Last known location of the missing Car
type Location struct {
	Lat float64 `json:"lat,omitempty"`
	Lng float64 `json:"lng,omitempty"`
}
