package postgresql

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"fyt/internal"
	"fyt/internal/api/api_models"
	"fyt/internal/app/models"
	"fyt/internal/storage/postgresql/db"

	"github.com/jackc/pgx/v5"
)

// TaskRepository represents the repository used for interacting with tasks records.
type TaskRepository struct {
	q *db.Queries
}

// NewFilm instantiates the Task repository.
func NewTaskRepo(d db.DBTX) *TaskRepository {
	r := TaskRepository{
		q: db.New(d),
	}

	return &r
}

// Create inserts a new record in db.
func (r *TaskRepository) Create(p api_models.CreateTask) (models.Task, error) {
	d, err := newDateFromString(p.DueDate)
	if err != nil {
		return models.Task{}, fmt.Errorf("create task error %w", err)
	}

	row, err := r.q.InsertTask(context.Background(), db.InsertTaskParams{
		ProjectID:   p.ProjectId,
		Title:       p.Title,
		Description: p.Description,
		DueDate:     d,
	})
	if err != nil {
		return models.Task{}, internal.WrapErrorf(err, internal.ErrorCodeUnknown, "insert Task")
	}

	return models.Task{
		Id:          row.ID,
		ProjectId:   p.ProjectId,
		Title:       p.Title,
		Description: p.Description,
		DueDate:     d.Time,
		CreatedAt:   row.CreatedAt.Time,
		UpdatedAt:   row.UpdatedAt.Time,
	}, nil
}

// Delete deletes the existing record matching the id.
func (r *TaskRepository) Delete(id int64) error {
	result, err := r.q.DeleteTask(context.Background(), id)
	if err != nil {
		return internal.WrapErrorf(err, internal.ErrorCodeUnknown, "delete task")
	}

	if result == 0 {
		return internal.WrapErrorf(err, internal.ErrorCodeNotFound, "delete task")
	}

	return nil
}

func (r *TaskRepository) Find(id int64) (models.Task, error) {
	task, err := r.q.SelectTask(context.Background(), id)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return models.Task{}, internal.WrapErrorf(err, internal.ErrorCodeNotFound, "task not found")
		}

		return models.Task{}, internal.WrapErrorf(err, internal.ErrorCodeUnknown, "select task")
	}

	return models.Task{
		Id:          task.ID,
		ProjectId:   task.ProjectID,
		Title:       task.Title,
		Description: task.Description,
		DueDate:     task.DueDate.Time,
		Doer:        task.Doer.Int32,
		CreatedAt:   task.CreatedAt.Time,
		UpdatedAt:   task.UpdatedAt.Time,
		Done:        task.Done.Bool,
	}, nil
}

func (r *TaskRepository) Update(id int64, p api_models.UpdateTask) error {
	d, err := newDateFromString(p.DueDate)
	if err != nil {
		return fmt.Errorf("update task error %w", err)
	}

	result, err := r.q.UpdateTask(context.Background(), db.UpdateTaskParams{
		ID:          id,
		Title:       p.Title,
		Description: p.Description,
		DueDate:     d,
	})
	if err != nil {
		if strings.Contains(err.Error(), "unique constraint") {
			return internal.WrapErrorf(err, internal.ErrorCodeUniqueConstraints, "update task")
		} else {
			return internal.WrapErrorf(err, internal.ErrorCodeUnknown, "update task")
		}
	}

	if result == 0 {
		return internal.WrapErrorf(err, internal.ErrorCodeNotFound, "update task")
	}

	return nil
}

func (r *TaskRepository) UpdateDoer(id int64, t api_models.UpdateDoer) error {
	result, err := r.q.UpdateDoer(context.Background(), db.UpdateDoerParams{
		ID:   id,
		Doer: newInt4(t.Doer),
	})
	if err != nil {
		return internal.WrapErrorf(err, internal.ErrorCodeUnknown, "update doer")
	}

	if result == 0 {
		return internal.WrapErrorf(err, internal.ErrorCodeNotFound, "update doer")
	}

	return nil
}

func (r *TaskRepository) UpdateDone(id int64, t api_models.UpdateDone) error {
	result, err := r.q.UpdateDone(context.Background(), db.UpdateDoneParams{
		ID:   id,
		Done: newBool(t.Done),
	})
	if err != nil {
		return internal.WrapErrorf(err, internal.ErrorCodeUnknown, "update done")
	}

	if result == 0 {
		return internal.WrapErrorf(err, internal.ErrorCodeNotFound, "update done")
	}

	return nil
}
