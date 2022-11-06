-- name: GetUser :one
SELECT * FROM users
WHERE email = $1 LIMIT 1;

-- name: CreateUser :one
INSERT INTO users (
  email, first_name, last_name, salt, hash
) VALUES (
  $1, $2, $3, $4, $5
)
RETURNING *;

-- name: CreateSession :one
INSERT INTO sessions (
  user_id, created_at, expired_at
) VALUES (
  $1, $2, $3
)
RETURNING id;

-- name: GetSessionUserID :one
SELECT user_id FROM sessions
WHERE id = $1 LIMIT 1;

-- name: GetUserByID :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;
