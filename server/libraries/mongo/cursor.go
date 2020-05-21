package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

func CaseCursor(cur *mongo.Cursor) (cases []primitive.M ,err error) {

	// Iterate through the cursor
	for cur.Next(context.TODO()) {
		// Create a value into which the single document can be decoded
		var elem bson.M
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		cases = append(cases, elem)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	// Close the cursor once finished
	if err := cur.Close(context.TODO()); err != nil {
		return nil, err
	}

	return cases,nil
}