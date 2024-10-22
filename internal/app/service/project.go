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
	// Search(
	// 	ctx context.Context,
	// 	name *string,
	// 	description *string,
	// ) ([]models.Project, error)
}

// ProjectService defines the application service in charge of interacting with Tasks.
type ProjectService struct {
	repo ProjectRepository
	// search ProjectSearchRepository
}

// NewProjectService
func NewProjectService(repo ProjectRepository) *ProjectService {
	return &ProjectService{
		repo: repo,
		// search: search,
	}
}

// Search gets all existing Project from the datastore.
// func (s *ProjectService) Search(
// 	ctx context.Context,
// 	name string,
// 	description string,
// 	releaseYear uint16,
// 	rating float32,
// ) ([]models.Project, error) {
// 	films, err := s.search.Search(ctx, &name, &description)
// 	if err != nil {
// 		return nil, fmt.Errorf("search: %w", err)
// 	}

// 	return films, nil
// }

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

	// _ = s.search.Index(ctx, film) // Ignoring errors on purpose

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

	// _ = s.search.Delete(ctx, id) // Ignoring errors on purpose

	return nil
}

// Find gets an existing Project from the datastore.
func (s *ProjectService) Find(id string) (models.Project, error) {
	iD, err := strconv.Atoi(id)
	if err != nil {
		return models.Project{}, fmt.Errorf("svc delete: %w", err)
	}

	task, err := s.repo.Find(int64(iD))
	if err != nil {
		return models.Project{}, fmt.Errorf("repo find: %w", err)
	}

	return task, nil
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

	// _, err := s.repo.Find(id)
	// if err == nil {
	// 	// _ = s.search.Index(ctx, film) // Ignoring errors on purpose
	// }
	return nil
}
