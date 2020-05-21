package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"../../models"
)

// CreateNewCase : Create New Case
func (m Mongo) CreateNewCase(newCase models.Case) (id interface{}, err error) {

	// Write the new Case onto the DB
	insertResult, err := m.DB.Collection("case").InsertOne(context.Background(), newCase)

	if err != nil || insertResult.InsertedID == nil {
		return  0, nil
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
