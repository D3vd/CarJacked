package auth

import "../../libraries/mongo"

// Controller : User login and sign-up
type Controller struct {
	M *mongo.Mongo
}
