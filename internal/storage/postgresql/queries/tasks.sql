-- name: SelectTask :one
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
  id = @id
LIMIT 1;

-- name: InsertTask :one
INSERT INTO tasks (
  project_id,
  title,
  description,
  due_date
)
VALUES (
  @project_id,
  @title,
  @description,
  @due_date
)
RETURNING id, created_at, updated_at;

-- name: UpdateTask :one
UPDATE tasks SET
  title       = @title,
  description = @description,
  due_date    = @due_date,
  updated_at  = NOW()
WHERE id = @id
RETURNING id AS res;

-- name: UpdateDoer :one
UPDATE tasks SET
  doer = @doer,
  updated_at  = NOW()
WHERE id = @id
RETURNING id AS res;

-- name: UpdateDone :one
UPDATE tasks SET
  done = @done,
  updated_at  = NOW()
WHERE id = @id
RETURNING id AS res;

-- name: DeleteTask :one
DELETE FROM
  tasks
WHERE
  id = @id
RETURNING id AS res;