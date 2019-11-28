package database

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoClient *mongo.Client

func GetMongoClient() *mongo.Client {
	return mongoClient
}

func MongoConnect() {
	clientOptions := options.Client().ApplyURI("mongodb+srv://triluong:rbjjgfVYeYNR033a@cluster0-0rlhg.mongodb.net/test?retryWrites=true&w=majority")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	mongoClient = client

	fmt.Println("Connected to MongoDB!")
}
