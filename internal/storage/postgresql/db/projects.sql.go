// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: projects.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const DeleteProject = `-- name: DeleteProject :one
DELETE FROM
  projects
WHERE
  id = $1
RETURNING id AS res
`

func (q *Queries) DeleteProject(ctx context.Context, id int64) (int64, error) {
	row := q.db.QueryRow(ctx, DeleteProject, id)
	var res int64
	err := row.Scan(&res)
	return res, err
}

const InsertProject = `-- name: InsertProject :one
INSERT INTO projects (
  owner,
  project_type,
  title,
  description,
  social_url,
  source_url 
)
VALUES (
  $1,
  $2,
  $3,
  $4,
  $5, 
  $6
)
RETURNING id, created_at
`

type InsertProjectParams struct {
	Owner       int32
	ProjectType int32
	Title       string
	Description string
	SocialUrl   []string
	SourceUrl   pgtype.Text
}

type InsertProjectRow struct {
	ID        int64
	CreatedAt pgtype.Timestamptz
}

func (q *Queries) InsertProject(ctx context.Context, arg InsertProjectParams) (InsertProjectRow, error) {
	row := q.db.QueryRow(ctx, InsertProject,
		arg.Owner,
		arg.ProjectType,
		arg.Title,
		arg.Description,
		arg.SocialUrl,
		arg.SourceUrl,
	)
	var i InsertProjectRow
	err := row.Scan(&i.ID, &i.CreatedAt)
	return i, err
}

const SelectProject = `-- name: SelectProject :one
SELECT
  id,
  owner,
  project_type,
  title,
  description,
  created_at,
  updated_at,
  social_url,
  source_url,
  closed,
  closed_at
FROM
  projects
WHERE
  id = $1
LIMIT 1
`

func (q *Queries) SelectProject(ctx context.Context, id int64) (Projects, error) {
	row := q.db.QueryRow(ctx, SelectProject, id)
	var i Projects
	err := row.Scan(
		&i.ID,
		&i.Owner,
		&i.ProjectType,
		&i.Title,
		&i.Description,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.SocialUrl,
		&i.SourceUrl,
		&i.Closed,
		&i.ClosedAt,
	)
	return i, err
}

const UpdateProject = `-- name: UpdateProject :one
UPDATE projects SET
  project_type = $1,
  title        = $2,
  description  = $3,
  social_url   = $4,
  source_url   = $5,
  closed       = $6,
  updated_at   = NOW()
WHERE id = $7
RETURNING id AS res
`

type UpdateProjectParams struct {
	ProjectType int32
	Title       string
	Description string
	SocialUrl   []string
	SourceUrl   pgtype.Text
	Closed      pgtype.Bool
	ID          int64
}

func (q *Queries) UpdateProject(ctx context.Context, arg UpdateProjectParams) (int64, error) {
	row := q.db.QueryRow(ctx, UpdateProject,
		arg.ProjectType,
		arg.Title,
		arg.Description,
		arg.SocialUrl,
		arg.SourceUrl,
		arg.Closed,
		arg.ID,
	)
	var res int64
	err := row.Scan(&res)
	return res, err
}
