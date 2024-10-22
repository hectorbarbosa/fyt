-- name: SelectProject :one
SELECT
  id,
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
  id = @id
LIMIT 1;

-- name: InsertProject :one
INSERT INTO projects (
  project_type,
  title,
  description,
  social_url,
  source_url 
)
VALUES (
  @project_type,
  @title,
  @description,
  @social_url, 
  @source_url
)
RETURNING id, created_at, updated_at;

-- name: UpdateProject :one
UPDATE projects SET
  project_type = @project_type,
  title        = @title,
  description  = @description,
  social_url   = @social_url,
  source_url   = @source_url,
  closed       = @closed
WHERE id = @id
RETURNING id AS res;

-- name: DeleteProject :one
DELETE FROM
  projects
WHERE
  id = @id
RETURNING id AS res;