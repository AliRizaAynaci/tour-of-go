package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// User represents a user document in MongoDB
type User struct {
	ID   primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name string             `bson:"name" json:"name"`
	Age  int                `bson:"age" json:"age"`
	City string             `bson:"city" json:"city"`
}
