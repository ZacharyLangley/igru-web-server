-- RPC Queries

-- name: GetNode :one
SELECT * FROM nodes
WHERE mac_address = $1 LIMIT 1;

-- name: DeleteNode :exec
DELETE FROM nodes
WHERE mac_address = $1;

-- name: GetNodes :many
SELECT * FROM nodes
WHERE owned_by = $1 LIMIT $2 OFFSET $3;

-- name: GetUnadoptedNodes :many
SELECT * FROM nodes
WHERE owned_by IS NULL  LIMIT $2 OFFSET $3;

-- name: AdoptNode :one
UPDATE nodes
SET owned_by=$2, updated_at=CURRENT_TIMESTAMP, adopted_at=CURRENT_TIMESTAMP
WHERE mac_address=$1
RETURNING *;

-- name: UnadoptNode :one
UPDATE nodes
SET owned_by=NULL, updated_at=CURRENT_TIMESTAMP, adopted_at=NULL
WHERE mac_address=$1
RETURNING *;

-- name: UpdateNode :one
UPDATE nodes
SET name=$2, custom_labels=$3, updated_at=CURRENT_TIMESTAMP
WHERE mac_address=$1
RETURNING *;

-- name: GetNodeSensors :many
SELECT * FROM sensors
WHERE node_id = $1;

-- name: UpdateNodeSensor :one
UPDATE sensors
SET name=$2, model=$3, category=$4
WHERE node_id=$1
RETURNING *;

-- Internal Queries

-- name: CreateNode :one
INSERT INTO nodes (
  mac_address, name, location, created_at
) VALUES (
  $1, $2, $3, CURRENT_TIMESTAMP
)
RETURNING *;

-- name: AddNodeSensor :one
INSERT INTO nodes (
  node_id, name, model, category
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;
