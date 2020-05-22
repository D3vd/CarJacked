package mongo

import (
	"context"
	"github.com/R3l3ntl3ss/CarJacked/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
)

// GetUnassignedOfficer : Get a list of unassigned officer
func (m Mongo) GetUnassignedOfficers() ([]models.Officer, error) {

	//	Create Filter
	filter := bson.M{"assigned": false}

	cur, err := m.DB.Collection("officer").Find(context.Background(), filter)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	officers, err := OfficerCursor(cur)

	if err != nil {
		return nil, err
	}

	return officers, nil
}

// CreateOfficer : Create a new Officer based on userID
func (m Mongo) CreateOfficer(officer models.Officer, userID string) error {

	// Convert user ID string to primitive Object
	userIDObj, err := primitive.ObjectIDFromHex(userID)

	if err != nil {
		return err
	}

	// Assign UserID to Officer
	officer.UserID = userIDObj

	// Insert Officer in to Officer Collection
	_, err = m.DB.Collection("officer").InsertOne(context.Background(), officer)

	if err != nil {
		return err
	}

	return nil
}
