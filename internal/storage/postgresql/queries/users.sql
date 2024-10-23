-- name: SelectUser :one
SELECT
  id,
  email,
  user_name,
  password,
  created_at
FROM
  users
WHERE
  id = @id
LIMIT 1;

-- name: InsertUser :one
INSERT INTO users (
  email,
  user_name,
  password
)
VALUES (
  @email,
  @user_name,
  @password
)
RETURNING id, created_at;

-- name: UpdateUser :one
UPDATE users SET
  email     = @email,
  user_name = @user_name,
  password  = @password
WHERE id = @id
RETURNING id AS res;

-- name: DeleteUser :one
DELETE FROM
  users
WHERE
  id = @id
RETURNING id AS res;