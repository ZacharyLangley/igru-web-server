-- name: GetGardens :many
SELECT * FROM gardens;

-- name: GetGarden :one
SELECT * FROM gardens
WHERE id = $1 LIMIT 1;

-- name: CreateGarden :one
INSERT INTO gardens (
  name, comment, location, grow_type, grow_size, grow_style, container_size, tags, created_at
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9
)
RETURNING *;

-- name: UpdateGarden :one
UPDATE gardens
SET name = $2, comment= $3, location = $4, grow_type = $5, grow_size = $6, grow_style = $7, container_size = $8, tags = $9, created_at = $10, updated_at=CURRENT_TIMESTAMP
WHERE id = $1
RETURNING *;

-- name: DeleteGarden :exec
DELETE FROM gardens
WHERE id = $1;
