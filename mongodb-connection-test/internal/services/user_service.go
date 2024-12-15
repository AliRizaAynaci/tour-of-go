package services

import (
	"context"
	"mongodb-connection-test/internal/models"
	"mongodb-connection-test/internal/repositories"
)

// UserService handles business logic
type UserService struct {
	repo *repositories.UserRepository
}

// NewUserService creates a new UserService
func NewUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{repo: repo}
}

// CreateUser creates a new user
func (service *UserService) CreateUser(ctx context.Context, user models.User) (string, error) {
	return service.repo.InsertUser(ctx, user)
}

// GetUserByID retrieves a user by ID
func (service *UserService) GetUserByID(ctx context.Context, id string) (*models.User, error) {
	return service.repo.FindUserByID(ctx, id)
}

// UpdateUser updates user data
func (service *UserService) UpdateUser(ctx context.Context, id string, update map[string]interface{}) error {
	return service.repo.UpdateUser(ctx, id, update)
}

// DeleteUser deletes a user by ID
func (service *UserService) DeleteUser(ctx context.Context, id string) error {
	return service.repo.DeleteUser(ctx, id)
}
