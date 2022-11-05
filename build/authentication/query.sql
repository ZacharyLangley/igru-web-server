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