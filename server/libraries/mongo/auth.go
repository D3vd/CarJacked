package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"

	"../../models"
)

// GetUserPasswordAndID : Queries DB for username and returns password and userID
func (m Mongo) GetUserPasswordAndID(username string) (password string, id string, error int) {

	// Create Empty Auth model
	var auth *models.Auth

	//	Create filter by username
	filter := bson.M{
		"username" : username,
	}

	// Query DB for username
	err := m.DB.Collection("user").FindOne(context.Background(), filter).Decode(&auth)

	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			return "", "", 400
		}
		return "", "", 500
	}

	return auth.Password, auth.ID.String(), 0
}
