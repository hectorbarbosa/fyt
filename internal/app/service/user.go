package service

import (
	"context"
	"fmt"
	"strconv"

	"fyt/internal"
	"fyt/internal/api/api_models"
	"fyt/internal/app/models"
)

// UserRepository defines the datastore handling Project records.
type UserRepository interface {
	Create(p api_models.CreateUser) (models.User, error)
	Delete(id int32) error
	Find(id int32) (models.User, error)
	Update(id int32, p api_models.UpdateUser) error
}

// UserService
type UserService struct {
	repo UserRepository
}

// New User Service
func NewUserService(repo UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

// Create stores a new record.
func (s *UserService) Create(
	ctx context.Context,
	u api_models.CreateUser,
) (models.User, error) {
	if err := u.Validate(); err != nil {
		return models.User{}, internal.WrapErrorf(err, internal.ErrorCodeInvalidArgument, "validate user")
	}

	user, err := s.repo.Create(u)
	if err != nil {
		return models.User{}, fmt.Errorf("repo create: %w", err)
	}

	return user, nil
}

// Delete removes an existing User from the datastore.
func (s *UserService) Delete(ctx context.Context, id string) error {
	iD, err := strconv.Atoi(id)
	if err != nil {
		return fmt.Errorf("svc delete: %w", err)
	}

	if err := s.repo.Delete(int32(iD)); err != nil {
		return fmt.Errorf("svc delete: %w", err)
	}

	return nil
}

// Find gets an existing User from the datastore.
func (s *UserService) Find(id string) (models.User, error) {
	iD, err := strconv.Atoi(id)
	if err != nil {
		return models.User{}, fmt.Errorf("svc delete: %w", err)
	}

	user, err := s.repo.Find(int32(iD))
	if err != nil {
		return models.User{}, fmt.Errorf("repo find: %w", err)
	}

	return user, nil
}

// Update updates an existing User in the datastore.
func (s *UserService) Update(
	ctx context.Context,
	id string,
	u api_models.UpdateUser,
) error {
	if err := u.Validate(); err != nil {
		return internal.WrapErrorf(err, internal.ErrorCodeInvalidArgument, "validate user")
	}

	iD, err := strconv.Atoi(id)
	if err != nil {
		return fmt.Errorf("svc delete: %w", err)
	}

	if err := s.repo.Update(int32(iD), u); err != nil {
		return fmt.Errorf("repo update: %w", err)
	}

	return nil
}
