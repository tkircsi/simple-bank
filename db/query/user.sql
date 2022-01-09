-- name: CreateUser :one
INSERT INTO users (
  username,
  hashed_password,
  full_name,
  email,
  password_changed_at
) VALUES (
  $1, $2, $3, $4, $5
) RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE username = $1 LIMIT 1;

-- name: GetUserForUpdate :one
SELECT * FROM users
WHERE username = $1 LIMIT 1
FOR NO KEY UPDATE;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY created_at
LIMIT $1
OFFSET $2;

-- name: UpdateUser :one
UPDATE users
SET 
  full_name = $2,
  email = $3
WHERE username = $1
RETURNING *;

-- name: UpdateUserPassword :one
UPDATE users
SET 
  hashed_password = $2,
  password_changed_at = $3
WHERE username = $1
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users WHERE username = $1;