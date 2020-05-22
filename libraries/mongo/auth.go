package mongo

import (
	"context"
	"github.com/R3l3ntl3ss/CarJacked/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CreateNewUser : Create a admin user and return the UserID
func (m Mongo) CreateNewUser(username string, password string) (userID string, error int) {

	//	Create Auth Model
	newUser := models.User{
		Password: password,
		Username: username,
	}

	// Check if the username is already exists
	var userPresent *models.User

	err := m.DB.Collection("user").FindOne(context.Background(), bson.M{"username": username}).Decode(&userPresent)

	// Check for errors
	if err != nil {
		if err.Error() != "mongo: no documents in result" {
			return "", 500
		}
	}

	// Make sure userPresent is nil
	if userPresent != nil {
		return "", 400
	}

	// Insert new user into DB
	insertResult, err := m.DB.Collection("user").InsertOne(context.Background(), newUser)

	if err != nil {
		return "", 500
	}

	return insertResult.InsertedID.(primitive.ObjectID).Hex(), 0
}

// GetUserPasswordAndID : Queries DB for username and returns password and userID
func (m Mongo) GetUserPasswordAndID(username string) (password string, id string, error int) {

	// Create Empty Auth model
	var auth *models.User

	//	Create filter by username
	filter := bson.M{
		"username": username,
	}

	// Query DB for username
	err := m.DB.Collection("user").FindOne(context.Background(), filter).Decode(&auth)

	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			return "", "", 400
		}
		return "", "", 500
	}

	return auth.Password, auth.ID.Hex(), 0
}
