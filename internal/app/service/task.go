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
type TaskRepository interface {
	Create(t api_models.CreateTask) (models.Task, error)
	Delete(id int64) error
	Find(id int64) (models.Task, error)
	Update(id int64, t api_models.UpdateTask) error
	UpdateDoer(id int64, t api_models.UpdateDoer) error
	UpdateDone(id int64, t api_models.UpdateDone) error
}

// TaskSearchRepository defines the datastore handling persisting Searchable Project records.
// type TaskSearchRepository interface {
// 	Delete(ctx context.Context, id int64) error
// 	Index(ctx context.Context, p models.Task) error
// }

// ProjectService
type TaskService struct {
	repo TaskRepository
}

// NewProjectService
func NewTaskService(repo TaskRepository) *TaskService {
	return &TaskService{
		repo: repo,
	}
}

// Create stores a new record.
func (s *TaskService) Create(
	ctx context.Context,
	p api_models.CreateTask,
) (models.Task, error) {
	if err := p.Validate(); err != nil {
		return models.Task{}, internal.WrapErrorf(err, internal.ErrorCodeInvalidArgument, "validate task")
	}

	task, err := s.repo.Create(p)
	if err != nil {
		return models.Task{}, fmt.Errorf("repo create: %w", err)
	}

	return task, nil
}

// Delete removes an existing Task from the datastore.
func (s *TaskService) Delete(ctx context.Context, id string) error {
	iD, err := strconv.Atoi(id)
	if err != nil {
		return fmt.Errorf("svc delete: %w", err)
	}

	if err := s.repo.Delete(int64(iD)); err != nil {
		return fmt.Errorf("svc delete: %w", err)
	}

	return nil
}

// Find gets an existing Task from the datastore.
func (s *TaskService) Find(id string) (models.Task, error) {
	iD, err := strconv.Atoi(id)
	if err != nil {
		return models.Task{}, fmt.Errorf("svc delete: %w", err)
	}

	project, err := s.repo.Find(int64(iD))
	if err != nil {
		return models.Task{}, fmt.Errorf("repo find: %w", err)
	}

	return project, nil
}

// Update updates an existing Task in the datastore.
func (s *TaskService) Update(
	ctx context.Context,
	id string,
	t api_models.UpdateTask,
) error {
	if err := t.Validate(); err != nil {
		return internal.WrapErrorf(err, internal.ErrorCodeInvalidArgument, "validate task")
	}

	iD, err := strconv.Atoi(id)
	if err != nil {
		return fmt.Errorf("svc update: %w", err)
	}

	if err := s.repo.Update(int64(iD), t); err != nil {
		return fmt.Errorf("repo update: %w", err)
	}

	return nil
}

// Update updates an Task Doer in the datastore.
func (s *TaskService) UpdateDoer(
	ctx context.Context,
	id string,
	t api_models.UpdateDoer,
) error {
	if err := t.Validate(); err != nil {
		return internal.WrapErrorf(err, internal.ErrorCodeInvalidArgument, "validate doer")
	}

	iD, err := strconv.Atoi(id)
	if err != nil {
		return fmt.Errorf("svc update doer: %w", err)
	}

	if err := s.repo.UpdateDoer(int64(iD), t); err != nil {
		return fmt.Errorf("repo update doer: %w", err)
	}

	return nil
}

// Update updates an Task Done in the datastore.
func (s *TaskService) UpdateDone(
	ctx context.Context,
	id string,
	t api_models.UpdateDone,
) error {
	if err := t.Validate(); err != nil {
		return internal.WrapErrorf(err, internal.ErrorCodeInvalidArgument, "validate done")
	}

	iD, err := strconv.Atoi(id)
	if err != nil {
		return fmt.Errorf("svc update done: %w", err)
	}

	if err := s.repo.UpdateDone(int64(iD), t); err != nil {
		return fmt.Errorf("repo update done: %w", err)
	}

	return nil
}
