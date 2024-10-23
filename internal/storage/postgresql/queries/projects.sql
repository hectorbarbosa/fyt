-- name: SelectProject :one
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
  id = @id
LIMIT 1;

-- name: InsertProject :one
INSERT INTO projects (
  owner,
  project_type,
  title,
  description,
  social_url,
  source_url 
)
VALUES (
  @owner,
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
  closed       = @closed,
  updated_at   = NOW()
WHERE id = @id
RETURNING id AS res;

-- name: DeleteProject :one
DELETE FROM
  projects
WHERE
  id = @id
RETURNING id AS res;