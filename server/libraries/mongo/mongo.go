package mongo

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo/options"

	"go.mongodb.org/mongo-driver/mongo"
)

// Mongo : MongoDB Struct with client
type Mongo struct {
	Client mongo.Client
	DB     mongo.Database
}

// Init : Helper function to initialize MongoDB
func (m *Mongo) Init() {
	mongoURI := os.Getenv("MONGODB_URI")

	if mongoURI == "" {
		mongoURI = "mongodb://127.0.0.1:27017/carjack"
	}

	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))

	if err != nil {
		log.Fatal("Error while intializing MongoDB server. Make sure the URI is correct..")
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	err = client.Connect(ctx)

	if err != nil {
		log.Fatal("Error while connecting to MongoDB server. Make sure it is running..")
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Connected to MongoDB Successfully")
	}

	m.Client = *client
	m.DB = *client.Database("carjack")
}
