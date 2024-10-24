// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: tasks.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const DeleteTask = `-- name: DeleteTask :one
DELETE FROM
  tasks
WHERE
  id = $1
RETURNING id AS res
`

func (q *Queries) DeleteTask(ctx context.Context, id int64) (int64, error) {
	row := q.db.QueryRow(ctx, DeleteTask, id)
	var res int64
	err := row.Scan(&res)
	return res, err
}

const InsertTask = `-- name: InsertTask :one
INSERT INTO tasks (
  project_id,
  title,
  description,
  due_date
)
VALUES (
  $1,
  $2,
  $3,
  $4
)
RETURNING id, created_at, updated_at
`

type InsertTaskParams struct {
	ProjectID   int64
	Title       string
	Description string
	DueDate     pgtype.Timestamptz
}

type InsertTaskRow struct {
	ID        int64
	CreatedAt pgtype.Timestamptz
	UpdatedAt pgtype.Timestamptz
}

func (q *Queries) InsertTask(ctx context.Context, arg InsertTaskParams) (InsertTaskRow, error) {
	row := q.db.QueryRow(ctx, InsertTask,
		arg.ProjectID,
		arg.Title,
		arg.Description,
		arg.DueDate,
	)
	var i InsertTaskRow
	err := row.Scan(&i.ID, &i.CreatedAt, &i.UpdatedAt)
	return i, err
}

const SelectTask = `-- name: SelectTask :one
SELECT
  id,
  project_id,
  title,
  description,
  due_date,
  doer,
  done,
  created_at,
  updated_at
FROM
  tasks
WHERE
  id = $1
LIMIT 1
`

type SelectTaskRow struct {
	ID          int64
	ProjectID   int64
	Title       string
	Description string
	DueDate     pgtype.Timestamptz
	Doer        pgtype.Int4
	Done        pgtype.Bool
	CreatedAt   pgtype.Timestamptz
	UpdatedAt   pgtype.Timestamptz
}

func (q *Queries) SelectTask(ctx context.Context, id int64) (SelectTaskRow, error) {
	row := q.db.QueryRow(ctx, SelectTask, id)
	var i SelectTaskRow
	err := row.Scan(
		&i.ID,
		&i.ProjectID,
		&i.Title,
		&i.Description,
		&i.DueDate,
		&i.Doer,
		&i.Done,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const UpdateDoer = `-- name: UpdateDoer :one
UPDATE tasks SET
  doer = $1,
  updated_at  = NOW()
WHERE id = $2
RETURNING id AS res
`

type UpdateDoerParams struct {
	Doer pgtype.Int4
	ID   int64
}

func (q *Queries) UpdateDoer(ctx context.Context, arg UpdateDoerParams) (int64, error) {
	row := q.db.QueryRow(ctx, UpdateDoer, arg.Doer, arg.ID)
	var res int64
	err := row.Scan(&res)
	return res, err
}

const UpdateDone = `-- name: UpdateDone :one
UPDATE tasks SET
  done = $1,
  updated_at  = NOW()
WHERE id = $2
RETURNING id AS res
`

type UpdateDoneParams struct {
	Done pgtype.Bool
	ID   int64
}

func (q *Queries) UpdateDone(ctx context.Context, arg UpdateDoneParams) (int64, error) {
	row := q.db.QueryRow(ctx, UpdateDone, arg.Done, arg.ID)
	var res int64
	err := row.Scan(&res)
	return res, err
}

const UpdateTask = `-- name: UpdateTask :one
UPDATE tasks SET
  title       = $1,
  description = $2,
  due_date    = $3,
  updated_at  = NOW()
WHERE id = $4
RETURNING id AS res
`

type UpdateTaskParams struct {
	Title       string
	Description string
	DueDate     pgtype.Timestamptz
	ID          int64
}

func (q *Queries) UpdateTask(ctx context.Context, arg UpdateTaskParams) (int64, error) {
	row := q.db.QueryRow(ctx, UpdateTask,
		arg.Title,
		arg.Description,
		arg.DueDate,
		arg.ID,
	)
	var res int64
	err := row.Scan(&res)
	return res, err
}
