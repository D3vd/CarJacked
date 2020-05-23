package mongo

import (
	"context"
	"errors"
	"github.com/R3l3ntl3ss/CarJacked/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
)

// MakeOfficerUnassigned : Change Office to unassigned
func (m Mongo) MakeOfficerUnassigned(officerID string) error {

	// Convert case ID string to primitive Object
	officerObjID, err := primitive.ObjectIDFromHex(officerID)

	if err != nil {
		return err
	}

	// Filter for case
	filter := bson.M{
		"_id": officerObjID,
	}

	// Update the officer ID and change assigned
	update := bson.D{
		{"$set", bson.D{
			{"assigned", false},
		}},
	}

	updateResult, err := m.DB.Collection("officer").UpdateOne(context.Background(), filter, update)

	log.Println(updateResult, err)

	if err != nil {
		return err
	}

	if updateResult.ModifiedCount == 0 {
		return errors.New("no document was updated")
	}

	return nil
}

// MakeOfficerAssigned : Change Officer to assigned
func (m Mongo) MakeOfficerAssigned(officerID primitive.ObjectID) error {

	//	Filter officer by ID
	filter := bson.M{
		"_id": officerID,
	}

	//	Update Assigned Status
	update := bson.D{
		{"$set", bson.D{
			{"assigned", true},
		}},
	}

	updateResult, err := m.DB.Collection("officer").UpdateOne(context.Background(), filter, update)

	log.Println(updateResult, err)

	if err != nil {
		return err
	}

	if updateResult.ModifiedCount == 0 {
		return errors.New("no document was updated")
	}

	return nil
}

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

	if len(officers) == 0 {
		return nil, errors.New("no officers found")
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
