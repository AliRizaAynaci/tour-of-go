package repositories

import (
	"context"
	"mongodb-connection-test/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	collection *mongo.Collection
}

// NewUserRepository creates a new UserRepository
func NewUserRepository(client *mongo.Client) *UserRepository {
	collection := client.Database("testdb").Collection("testdbcollection")
	return &UserRepository{collection: collection}
}

// InsertUser adds a new user to the database
func (repo *UserRepository) InsertUser(ctx context.Context, user models.User) (string, error) {
	result, err := repo.collection.InsertOne(ctx, user)
	if err != nil {
		return "", err
	}
	return result.InsertedID.(primitive.ObjectID).Hex(), nil
}

// FindUserByID retrieves a user by ID
func (repo *UserRepository) FindUserByID(ctx context.Context, id string) (*models.User, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	var user models.User
	err = repo.collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&user)
	return &user, err
}

// UpdateUser updates user data
func (repo *UserRepository) UpdateUser(ctx context.Context, id string, update bson.M) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = repo.collection.UpdateOne(ctx, bson.M{"_id": objectID}, bson.M{"$set": update})
	return err
}

// DeleteUser removes a user by ID
func (repo *UserRepository) DeleteUser(ctx context.Context, id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = repo.collection.DeleteOne(ctx, bson.M{"_id": objectID})
	return err
}
