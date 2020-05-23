package admin

import "github.com/R3l3ntl3ss/CarJacked/libraries/mongo"

// Controller : Admin routes
type Controller struct {
	M *mongo.Mongo
	Secret string
}
