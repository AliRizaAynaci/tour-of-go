package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://admin:password@localhost:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	// Close the connection once the function returns
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			log.Fatal(err)
		}
		fmt.Println("Connection to MongoDB closed.")
	}()

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	collection := client.Database("testdb").Collection("testcollection")

	// Create
	newDocument := bson.D{
		{Key: "name", Value: "John Doe"},
		{Key: "age", Value: 30},
		{Key: "city", Value: "San Francisco"},
	}

	insertResult, err := collection.InsertOne(context.TODO(), newDocument)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a single document: ", insertResult.InsertedID)

	// Read
	var result bson.D
	filter := bson.D{{Key: "name", Value: "John Doe"}}

	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Found a single document: ", result)

	// Update
	filter = bson.D{{Key: "name", Value: "John Doe"}}
	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "age", Value: 31},
		}},
	}

	updateResult, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)

	// Delete
	filter = bson.D{{Key: "name", Value: "John Doe"}}

	deleteResult, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Deleted %v documents in the collection\n", deleteResult.DeletedCount)
}
