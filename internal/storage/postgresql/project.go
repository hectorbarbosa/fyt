package postgresql

import (
	"context"
	"errors"
	"strings"

	"fyt/internal"
	"fyt/internal/api/api_models"
	"fyt/internal/app/models"
	"fyt/internal/storage/postgresql/db"

	"github.com/jackc/pgx/v5"
)

// ProjectRepositary represents the repository used for interacting with projects records.
type ProjectRepository struct {
	q *db.Queries
}

// NewFilm instantiates the Project repository.
func NewProjectRepo(d db.DBTX) *ProjectRepository {
	r := ProjectRepository{
		q: db.New(d),
	}

	return &r
}

// Create inserts a new record in db.
func (r *ProjectRepository) Create(p api_models.CreateProject) (models.Project, error) {
	row, err := r.q.InsertProject(context.Background(), db.InsertProjectParams{
		ProjectType: int32(p.ProjectType),
		Title:       p.Title,
		Description: p.Description,
		SocialUrl:   p.SocialUrl,
		SourceUrl:   newText(p.SourceUrl),
	})
	if err != nil {
		return models.Project{}, internal.WrapErrorf(err, internal.ErrorCodeUnknown, "insert Project")
	}

	return models.Project{
		Id:          row.ID,
		ProjectType: models.ProjectType(p.ProjectType),
		Title:       p.Title,
		Description: p.Description,
		SocialUrl:   p.SocialUrl,
		SourceUrl:   p.SourceUrl,
		CreatedAt:   row.CreatedAt.Time,
		UpdatedAt:   row.UpdatedAt.Time,
	}, nil
}

// Delete deletes the existing record matching the id.
func (r *ProjectRepository) Delete(id int64) error {
	result, err := r.q.DeleteProject(context.Background(), id)
	if err != nil {
		return internal.WrapErrorf(err, internal.ErrorCodeUnknown, "delete project")
	}

	if result == 0 {
		return internal.WrapErrorf(err, internal.ErrorCodeNotFound, "delete project")
	}

	return nil
}

func (r *ProjectRepository) Find(id int64) (models.Project, error) {
	project, err := r.q.SelectProject(context.Background(), id)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return models.Project{}, internal.WrapErrorf(err, internal.ErrorCodeNotFound, "project not found")
		}

		return models.Project{}, internal.WrapErrorf(err, internal.ErrorCodeUnknown, "select project")
	}

	return models.Project{
		Id:          project.ID,
		ProjectType: models.ProjectType(project.ProjectType),
		Title:       project.Title,
		Description: project.Description,
		SocialUrl:   project.SocialUrl,
		SourceUrl:   project.SourceUrl.String,
		CreatedAt:   project.CreatedAt.Time,
		UpdatedAt:   project.UpdatedAt.Time,
		Closed:      project.Closed.Bool,
		ClosedAt:    project.ClosedAt.Time,
	}, nil
}

func (r *ProjectRepository) Update(id int64, p api_models.UpdateProject) error {
	result, err := r.q.UpdateProject(context.Background(), db.UpdateProjectParams{
		ID:          id,
		ProjectType: int32(p.ProjectType),
		Title:       p.Title,
		Description: p.Description,
		SocialUrl:   p.SocialUrl,
		SourceUrl:   newText(p.SourceUrl),
		Closed:      newBool(p.Closed),
	})
	if err != nil {
		if strings.Contains(err.Error(), "unique constraint") {
			return internal.WrapErrorf(err, internal.ErrorCodeUniqueConstraints, "update project")
		} else {
			return internal.WrapErrorf(err, internal.ErrorCodeUnknown, "update project")
		}
	}

	if result == 0 {
		return internal.WrapErrorf(err, internal.ErrorCodeNotFound, "delete project")
	}

	return nil
}
