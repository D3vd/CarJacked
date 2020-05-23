package mongo

import (
	"context"
	"errors"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/R3l3ntl3ss/CarJacked/models"
)

// GetCaseByOfficerID : Get a case by ID
func (m Mongo) GetCaseByOfficerID(id string, caseDoc *models.Case) (error int) {
	// Convert case ID string to primitive Object
	officerObjID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return 400
	}

	// Create filter for search
	filter := bson.M{
		"officer": officerObjID,
	}

	//	Query DB for case
	err = m.DB.Collection("case").FindOne(context.TODO(), filter).Decode(&caseDoc)

	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			return 400
		}
		return 500
	}

	return 0
}

// UpdateCaseWithOfficerID : Update An unassigned case with new officer
func (m Mongo) UpdateCaseWithOfficerID(caseID primitive.ObjectID, officerID primitive.ObjectID) error {

	// Filter for case
	filter := bson.M{
		"_id": caseID,
	}

	// Update the officer ID and change assigned
	update := bson.D{
		{"$set", bson.D{
			{"officer", officerID},
			{"assigned", true},
		}},
	}

	updateResult, err := m.DB.Collection("case").UpdateOne(context.Background(), filter, update)

	if err != nil {
		return err
	}

	if updateResult.ModifiedCount == 0 {
		return errors.New("no document was updated")
	}

	return nil
}

// GetAllUnassignedCases : Get all active cases
func (m Mongo) GetAllUnassignedCases() (cases []models.Case, err error) {

	cur, err := m.DB.Collection("case").Find(context.Background(), bson.M{"assigned": false})

	if err != nil {
		log.Println(err)
		return nil, err
	}

	cases, err = CaseCursor(cur)

	if err != nil {
		return nil, err
	}

	if len(cases) == 0 {
		return nil, errors.New("no unassigned cases were found")
	}

	return cases, nil
}

// GetAllActiveCases : Get all active cases
func (m Mongo) GetAllActiveCases() (cases []models.Case, err error) {

	cur, err := m.DB.Collection("case").Find(context.Background(), bson.M{"active": true})

	if err != nil {
		log.Println(err)
		return nil, err
	}

	cases, err = CaseCursor(cur)

	return cases, nil
}

// CreateNewCase : Create New Case
func (m Mongo) CreateNewCase(newCase models.Case) (id interface{}, err error) {

	// Write the new Case onto the DB
	insertResult, err := m.DB.Collection("case").InsertOne(context.Background(), newCase)

	if err != nil || insertResult.InsertedID == nil {
		return 0, err
	}

	return insertResult.InsertedID, nil
}

// GetCaseByID : Get a case by ID
func (m Mongo) GetCaseByID(id string, caseDoc *models.Case) (error int) {
	// Convert case ID string to primitive Object
	caseObjID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return 400
	}

	// Create filter for search
	filter := bson.M{
		"_id": caseObjID,
	}

	//	Query DB for case
	err = m.DB.Collection("case").FindOne(context.TODO(), filter).Decode(&caseDoc)

	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			return 400
		}
		return 500
	}

	return 0
}
