package service

import (
	"context"
	"fmt"
	"strconv"

	"fyt/internal"
	"fyt/internal/api/api_models"
	"fyt/internal/app/models"
)

// ProjectRepository defines the datastore handling Project records.
type ProjectRepository interface {
	Create(p api_models.CreateProject) (models.Project, error)
	Delete(id int64) error
	Find(id int64) (models.Project, error)
	Update(id int64, p api_models.UpdateProject) error
}

// ProjectSearchRepository defines the datastore handling persisting Searchable Project records.
type ProjectSearchRepository interface {
	Delete(ctx context.Context, id int64) error
	Index(ctx context.Context, p models.Project) error
}

// ProjectService
type ProjectService struct {
	repo ProjectRepository
}

// NewProjectService
func NewProjectService(repo ProjectRepository) *ProjectService {
	return &ProjectService{
		repo: repo,
	}
}

// Create stores a new record.
func (s *ProjectService) Create(
	ctx context.Context,
	p api_models.CreateProject,
) (models.Project, error) {
	if err := p.Validate(); err != nil {
		return models.Project{}, internal.WrapErrorf(err, internal.ErrorCodeInvalidArgument, "validate project")
	}

	project, err := s.repo.Create(p)
	if err != nil {
		return models.Project{}, fmt.Errorf("repo create: %w", err)
	}

	return project, nil
}

// Delete removes an existing Project from the datastore.
func (s *ProjectService) Delete(ctx context.Context, id string) error {
	iD, err := strconv.Atoi(id)
	if err != nil {
		return fmt.Errorf("svc delete: %w", err)
	}

	if err := s.repo.Delete(int64(iD)); err != nil {
		return fmt.Errorf("svc delete: %w", err)
	}

	return nil
}

// Find gets an existing Project from the datastore.
func (s *ProjectService) Find(id string) (models.Project, error) {
	iD, err := strconv.Atoi(id)
	if err != nil {
		return models.Project{}, fmt.Errorf("svc delete: %w", err)
	}

	project, err := s.repo.Find(int64(iD))
	if err != nil {
		return models.Project{}, fmt.Errorf("repo find: %w", err)
	}

	return project, nil
}

// Update updates an existing Project in the datastore.
func (s *ProjectService) Update(
	ctx context.Context,
	id string,
	p api_models.UpdateProject,
) error {
	if err := p.Validate(); err != nil {
		return internal.WrapErrorf(err, internal.ErrorCodeInvalidArgument, "validate project")
	}

	iD, err := strconv.Atoi(id)
	if err != nil {
		return fmt.Errorf("svc delete: %w", err)
	}

	if err := s.repo.Update(int64(iD), p); err != nil {
		return fmt.Errorf("repo update: %w", err)
	}

	return nil
}
