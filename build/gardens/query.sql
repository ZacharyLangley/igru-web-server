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

-- name: GetPlants :many
SELECT * FROM plants;

-- name: GetPlant :one
SELECT * FROM plants
WHERE id = $1 LIMIT 1;

-- name: CreatePlant :one
INSERT INTO plants (
  name, comment, notes, grow_cycle_length, parentage, origin, gender, days_flowering, days_cured, harvested_weight, bud_density, bud_pistils, tags, acquired_at, created_at
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15
)
RETURNING *;

-- name: UpdatePlant :one
UPDATE plants
SET name = $2, comment = $3, notes = $4, grow_cycle_length = $5, parentage = $6, origin = $7, gender = $8, days_flowering = $9, days_cured = $10, harvested_weight = $11, bud_density = $12, bud_pistils = $13, tags = $14, acquired_at=$15, created_at = $16, updated_at=CURRENT_TIMESTAMP
WHERE id = $1
RETURNING *;

-- name: DeletePlant :exec
DELETE FROM plants
WHERE id = $1;

-- name: GetStrains :many
SELECT * FROM strains;

-- name: GetStrain :one
SELECT * FROM strains
WHERE id = $1 LIMIT 1;

-- name: CreateStrain :one
INSERT INTO strains (
  name, comment, notes, type, price, thc_percent, cbd_percent, parentage, aroma, taste, tags, created_at
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12
)
RETURNING *;

-- name: UpdateStrain :one
UPDATE strains
SET name = $2, comment = $3, notes = $4, type = $5, price = $6, thc_percent = $7, cbd_percent = $8, parentage = $9, aroma = $10, taste = $11, tags = $12, created_at = $13, updated_at=CURRENT_TIMESTAMP
WHERE id = $1
RETURNING *;

-- name: DeleteStrain :exec
DELETE FROM strains
WHERE id = $1;

-- name: GetRecipes :many
SELECT * FROM recipes;

-- name: GetRecipe :one
SELECT * FROM recipes
WHERE id = $1 LIMIT 1;

-- name: CreateRecipe :one
INSERT INTO recipes (
  name, comment, ingredients, instructions, ph, mix_time_milliseconds, tags, created_at
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8
)
RETURNING *;

-- name: UpdateRecipe :one
UPDATE recipes
SET name = $2, comment = $3, ingredients = $4, instructions = $5, ph = $6, mix_time_milliseconds = $7, tags = $8, created_at = $9, updated_at=CURRENT_TIMESTAMP
WHERE id = $1
RETURNING *;

-- name: DeleteRecipe :exec
DELETE FROM recipes
WHERE id = $1;