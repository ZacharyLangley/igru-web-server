-- name: GetGardens :many
SELECT * FROM gardens
WHERE group_id = $1;

-- name: GetGarden :one
SELECT * FROM gardens
WHERE id = $1 AND group_id = $2 LIMIT 1;

-- name: CreateGarden :one
INSERT INTO gardens (
  name, group_id, comment, location, grow_type, grow_size, grow_style, container_size, tags, created_at
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9, $10
)
RETURNING *;

-- name: UpdateGarden :one
UPDATE gardens
SET name = $3, comment= $4, location = $5, grow_type = $6, grow_size = $7, grow_style = $8, container_size = $9, tags = $10, created_at = $11, updated_at=CURRENT_TIMESTAMP
WHERE id = $1 AND group_id = $2
RETURNING *;

-- name: DeleteGarden :exec
DELETE FROM gardens
WHERE id = $1 AND group_id = $2;

-- name: GetPlants :many
SELECT * FROM plants
WHERE group_id = $1;

-- name: GetPlant :one
SELECT * FROM plants
WHERE id = $1 AND group_id = $2 LIMIT 1;

-- name: CreatePlant :one
INSERT INTO plants (
  name, group_id, comment, notes, grow_cycle_length, parentage, origin, gender, days_flowering, days_cured, harvested_weight, bud_density, bud_pistils, tags, acquired_at, created_at
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, CURRENT_TIMESTAMP
)
RETURNING *;

-- name: UpdatePlant :one
UPDATE plants
SET name = $3, comment = $4, notes = $5, grow_cycle_length = $6, parentage = $7, origin = $8, gender = $9, days_flowering = $10, days_cured = $11, harvested_weight = $12, bud_density = $13, bud_pistils = $14, tags = $15, acquired_at=$16, updated_at=CURRENT_TIMESTAMP
WHERE id = $1 AND group_id = $2
RETURNING *;

-- name: DeletePlant :exec
DELETE FROM plants
WHERE id = $1 AND group_id = $2;

-- name: GetStrains :many
SELECT * FROM strains
WHERE group_id = $1;

-- name: GetStrain :one
SELECT * FROM strains
WHERE id = $1 AND group_id = $2 LIMIT 1;

-- name: CreateStrain :one
INSERT INTO strains (
  name, group_id, comment, notes, type, price, thc_percent, cbd_percent, parentage, aroma, taste, tags, created_at
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13
)
RETURNING *;

-- name: UpdateStrain :one
UPDATE strains
SET name = $3, comment = $4, notes = $5, type = $6, price = $7, thc_percent = $8, cbd_percent = $9, parentage = $10, aroma = $11, taste = $12, tags = $13, created_at = $14, updated_at=CURRENT_TIMESTAMP
WHERE id = $1 AND group_id = $2
RETURNING *;

-- name: DeleteStrain :exec
DELETE FROM strains
WHERE id = $1 AND group_id = $2;

-- name: GetRecipes :many
SELECT * FROM recipes
WHERE group_id = $1;

-- name: GetRecipe :one
SELECT * FROM recipes
WHERE id = $1 AND group_id = $2 LIMIT 1;

-- name: CreateRecipe :one
INSERT INTO recipes (
  name, group_id, comment, ingredients, instructions, ph, mix_time_milliseconds, tags, created_at
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9
)
RETURNING *;

-- name: UpdateRecipe :one
UPDATE recipes
SET name = $3, comment = $4, ingredients = $5, instructions = $6, ph = $7, mix_time_milliseconds = $8, tags = $9, created_at = $10, updated_at=CURRENT_TIMESTAMP
WHERE id = $1 AND group_id = $2
RETURNING *;

-- name: DeleteRecipe :exec
DELETE FROM recipes
WHERE id = $1 AND group_id = $2;