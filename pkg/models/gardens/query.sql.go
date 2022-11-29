// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: query.sql

package gardens

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createGarden = `-- name: CreateGarden :one
INSERT INTO gardens (
  name, comment, location, grow_type, grow_size, grow_style, container_size, tags, created_at
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9
)
RETURNING id, name, comment, location, grow_type, grow_size, grow_style, container_size, tags, created_at, updated_at
`

type CreateGardenParams struct {
	Name          string
	Comment       string
	Location      string
	GrowType      string
	GrowSize      string
	GrowStyle     string
	ContainerSize string
	Tags          string
	CreatedAt     time.Time
}

func (q *Queries) CreateGarden(ctx context.Context, arg CreateGardenParams) (Garden, error) {
	row := q.db.QueryRow(ctx, createGarden,
		arg.Name,
		arg.Comment,
		arg.Location,
		arg.GrowType,
		arg.GrowSize,
		arg.GrowStyle,
		arg.ContainerSize,
		arg.Tags,
		arg.CreatedAt,
	)
	var i Garden
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Comment,
		&i.Location,
		&i.GrowType,
		&i.GrowSize,
		&i.GrowStyle,
		&i.ContainerSize,
		&i.Tags,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const createPlant = `-- name: CreatePlant :one
INSERT INTO plants (
  name, comment, notes, grow_cycle_length, parentage, origin, gender, days_flowering, days_cured, harvested_weight, bud_density, bud_pistils, tags, acquired_at, created_at
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15
)
RETURNING id, name, comment, notes, grow_cycle_length, parentage, origin, gender, days_flowering, days_cured, harvested_weight, bud_density, bud_pistils, tags, acquired_at, created_at, updated_at
`

type CreatePlantParams struct {
	Name            string
	Comment         string
	Notes           string
	GrowCycleLength string
	Parentage       uuid.UUID
	Origin          string
	Gender          string
	DaysFlowering   float64
	DaysCured       float64
	HarvestedWeight string
	BudDensity      float64
	BudPistils      bool
	Tags            string
	AcquiredAt      time.Time
	CreatedAt       time.Time
}

func (q *Queries) CreatePlant(ctx context.Context, arg CreatePlantParams) (Plant, error) {
	row := q.db.QueryRow(ctx, createPlant,
		arg.Name,
		arg.Comment,
		arg.Notes,
		arg.GrowCycleLength,
		arg.Parentage,
		arg.Origin,
		arg.Gender,
		arg.DaysFlowering,
		arg.DaysCured,
		arg.HarvestedWeight,
		arg.BudDensity,
		arg.BudPistils,
		arg.Tags,
		arg.AcquiredAt,
		arg.CreatedAt,
	)
	var i Plant
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Comment,
		&i.Notes,
		&i.GrowCycleLength,
		&i.Parentage,
		&i.Origin,
		&i.Gender,
		&i.DaysFlowering,
		&i.DaysCured,
		&i.HarvestedWeight,
		&i.BudDensity,
		&i.BudPistils,
		&i.Tags,
		&i.AcquiredAt,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const createStrain = `-- name: CreateStrain :one
INSERT INTO strains (
  name, comment, notes, type, price, thc_percent, cbd_percent, parentage, aroma, taste, tags, created_at
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12
)
RETURNING id, name, comment, notes, type, price, thc_percent, cbd_percent, parentage, aroma, taste, tags, created_at, updated_at
`

type CreateStrainParams struct {
	Name       string
	Comment    string
	Notes      string
	Type       string
	Price      float64
	ThcPercent float64
	CbdPercent float64
	Parentage  uuid.UUID
	Aroma      string
	Taste      string
	Tags       string
	CreatedAt  time.Time
}

func (q *Queries) CreateStrain(ctx context.Context, arg CreateStrainParams) (Strain, error) {
	row := q.db.QueryRow(ctx, createStrain,
		arg.Name,
		arg.Comment,
		arg.Notes,
		arg.Type,
		arg.Price,
		arg.ThcPercent,
		arg.CbdPercent,
		arg.Parentage,
		arg.Aroma,
		arg.Taste,
		arg.Tags,
		arg.CreatedAt,
	)
	var i Strain
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Comment,
		&i.Notes,
		&i.Type,
		&i.Price,
		&i.ThcPercent,
		&i.CbdPercent,
		&i.Parentage,
		&i.Aroma,
		&i.Taste,
		&i.Tags,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteGarden = `-- name: DeleteGarden :exec
DELETE FROM gardens
WHERE id = $1
`

func (q *Queries) DeleteGarden(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.Exec(ctx, deleteGarden, id)
	return err
}

const deletePlant = `-- name: DeletePlant :exec
DELETE FROM plants
WHERE id = $1
`

func (q *Queries) DeletePlant(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.Exec(ctx, deletePlant, id)
	return err
}

const deleteStrain = `-- name: DeleteStrain :exec
DELETE FROM strains
WHERE id = $1
`

func (q *Queries) DeleteStrain(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.Exec(ctx, deleteStrain, id)
	return err
}

const getGarden = `-- name: GetGarden :one
SELECT id, name, comment, location, grow_type, grow_size, grow_style, container_size, tags, created_at, updated_at FROM gardens
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetGarden(ctx context.Context, id uuid.UUID) (Garden, error) {
	row := q.db.QueryRow(ctx, getGarden, id)
	var i Garden
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Comment,
		&i.Location,
		&i.GrowType,
		&i.GrowSize,
		&i.GrowStyle,
		&i.ContainerSize,
		&i.Tags,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getGardens = `-- name: GetGardens :many
SELECT id, name, comment, location, grow_type, grow_size, grow_style, container_size, tags, created_at, updated_at FROM gardens
`

func (q *Queries) GetGardens(ctx context.Context) ([]Garden, error) {
	rows, err := q.db.Query(ctx, getGardens)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Garden
	for rows.Next() {
		var i Garden
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Comment,
			&i.Location,
			&i.GrowType,
			&i.GrowSize,
			&i.GrowStyle,
			&i.ContainerSize,
			&i.Tags,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getPlant = `-- name: GetPlant :one
SELECT id, name, comment, notes, grow_cycle_length, parentage, origin, gender, days_flowering, days_cured, harvested_weight, bud_density, bud_pistils, tags, acquired_at, created_at, updated_at FROM plants
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetPlant(ctx context.Context, id uuid.UUID) (Plant, error) {
	row := q.db.QueryRow(ctx, getPlant, id)
	var i Plant
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Comment,
		&i.Notes,
		&i.GrowCycleLength,
		&i.Parentage,
		&i.Origin,
		&i.Gender,
		&i.DaysFlowering,
		&i.DaysCured,
		&i.HarvestedWeight,
		&i.BudDensity,
		&i.BudPistils,
		&i.Tags,
		&i.AcquiredAt,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getPlants = `-- name: GetPlants :many
SELECT id, name, comment, notes, grow_cycle_length, parentage, origin, gender, days_flowering, days_cured, harvested_weight, bud_density, bud_pistils, tags, acquired_at, created_at, updated_at FROM plants
`

func (q *Queries) GetPlants(ctx context.Context) ([]Plant, error) {
	rows, err := q.db.Query(ctx, getPlants)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Plant
	for rows.Next() {
		var i Plant
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Comment,
			&i.Notes,
			&i.GrowCycleLength,
			&i.Parentage,
			&i.Origin,
			&i.Gender,
			&i.DaysFlowering,
			&i.DaysCured,
			&i.HarvestedWeight,
			&i.BudDensity,
			&i.BudPistils,
			&i.Tags,
			&i.AcquiredAt,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getStrain = `-- name: GetStrain :one
SELECT id, name, comment, notes, type, price, thc_percent, cbd_percent, parentage, aroma, taste, tags, created_at, updated_at FROM strains
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetStrain(ctx context.Context, id uuid.UUID) (Strain, error) {
	row := q.db.QueryRow(ctx, getStrain, id)
	var i Strain
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Comment,
		&i.Notes,
		&i.Type,
		&i.Price,
		&i.ThcPercent,
		&i.CbdPercent,
		&i.Parentage,
		&i.Aroma,
		&i.Taste,
		&i.Tags,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getStrains = `-- name: GetStrains :many
SELECT id, name, comment, notes, type, price, thc_percent, cbd_percent, parentage, aroma, taste, tags, created_at, updated_at FROM strains
`

func (q *Queries) GetStrains(ctx context.Context) ([]Strain, error) {
	rows, err := q.db.Query(ctx, getStrains)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Strain
	for rows.Next() {
		var i Strain
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Comment,
			&i.Notes,
			&i.Type,
			&i.Price,
			&i.ThcPercent,
			&i.CbdPercent,
			&i.Parentage,
			&i.Aroma,
			&i.Taste,
			&i.Tags,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateGarden = `-- name: UpdateGarden :one
UPDATE gardens
SET name = $2, comment= $3, location = $4, grow_type = $5, grow_size = $6, grow_style = $7, container_size = $8, tags = $9, created_at = $10, updated_at=CURRENT_TIMESTAMP
WHERE id = $1
RETURNING id, name, comment, location, grow_type, grow_size, grow_style, container_size, tags, created_at, updated_at
`

type UpdateGardenParams struct {
	ID            uuid.UUID
	Name          string
	Comment       string
	Location      string
	GrowType      string
	GrowSize      string
	GrowStyle     string
	ContainerSize string
	Tags          string
	CreatedAt     time.Time
}

func (q *Queries) UpdateGarden(ctx context.Context, arg UpdateGardenParams) (Garden, error) {
	row := q.db.QueryRow(ctx, updateGarden,
		arg.ID,
		arg.Name,
		arg.Comment,
		arg.Location,
		arg.GrowType,
		arg.GrowSize,
		arg.GrowStyle,
		arg.ContainerSize,
		arg.Tags,
		arg.CreatedAt,
	)
	var i Garden
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Comment,
		&i.Location,
		&i.GrowType,
		&i.GrowSize,
		&i.GrowStyle,
		&i.ContainerSize,
		&i.Tags,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updatePlant = `-- name: UpdatePlant :one
UPDATE plants
SET name = $2, comment = $3, notes = $4, grow_cycle_length = $5, parentage = $6, origin = $7, gender = $8, days_flowering = $9, days_cured = $10, harvested_weight = $11, bud_density = $12, bud_pistils = $13, tags = $14, acquired_at=$15, created_at = $16, updated_at=CURRENT_TIMESTAMP
WHERE id = $1
RETURNING id, name, comment, notes, grow_cycle_length, parentage, origin, gender, days_flowering, days_cured, harvested_weight, bud_density, bud_pistils, tags, acquired_at, created_at, updated_at
`

type UpdatePlantParams struct {
	ID              uuid.UUID
	Name            string
	Comment         string
	Notes           string
	GrowCycleLength string
	Parentage       uuid.UUID
	Origin          string
	Gender          string
	DaysFlowering   float64
	DaysCured       float64
	HarvestedWeight string
	BudDensity      float64
	BudPistils      bool
	Tags            string
	AcquiredAt      time.Time
	CreatedAt       time.Time
}

func (q *Queries) UpdatePlant(ctx context.Context, arg UpdatePlantParams) (Plant, error) {
	row := q.db.QueryRow(ctx, updatePlant,
		arg.ID,
		arg.Name,
		arg.Comment,
		arg.Notes,
		arg.GrowCycleLength,
		arg.Parentage,
		arg.Origin,
		arg.Gender,
		arg.DaysFlowering,
		arg.DaysCured,
		arg.HarvestedWeight,
		arg.BudDensity,
		arg.BudPistils,
		arg.Tags,
		arg.AcquiredAt,
		arg.CreatedAt,
	)
	var i Plant
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Comment,
		&i.Notes,
		&i.GrowCycleLength,
		&i.Parentage,
		&i.Origin,
		&i.Gender,
		&i.DaysFlowering,
		&i.DaysCured,
		&i.HarvestedWeight,
		&i.BudDensity,
		&i.BudPistils,
		&i.Tags,
		&i.AcquiredAt,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateStrain = `-- name: UpdateStrain :one
UPDATE strains
SET name = $2, comment = $3, notes = $4, type = $5, price = $6, thc_percent = $7, cbd_percent = $8, parentage = $9, aroma = $10, taste = $11, tags = $12, created_at = $13, updated_at=CURRENT_TIMESTAMP
WHERE id = $1
RETURNING id, name, comment, notes, type, price, thc_percent, cbd_percent, parentage, aroma, taste, tags, created_at, updated_at
`

type UpdateStrainParams struct {
	ID         uuid.UUID
	Name       string
	Comment    string
	Notes      string
	Type       string
	Price      float64
	ThcPercent float64
	CbdPercent float64
	Parentage  uuid.UUID
	Aroma      string
	Taste      string
	Tags       string
	CreatedAt  time.Time
}

func (q *Queries) UpdateStrain(ctx context.Context, arg UpdateStrainParams) (Strain, error) {
	row := q.db.QueryRow(ctx, updateStrain,
		arg.ID,
		arg.Name,
		arg.Comment,
		arg.Notes,
		arg.Type,
		arg.Price,
		arg.ThcPercent,
		arg.CbdPercent,
		arg.Parentage,
		arg.Aroma,
		arg.Taste,
		arg.Tags,
		arg.CreatedAt,
	)
	var i Strain
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Comment,
		&i.Notes,
		&i.Type,
		&i.Price,
		&i.ThcPercent,
		&i.CbdPercent,
		&i.Parentage,
		&i.Aroma,
		&i.Taste,
		&i.Tags,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
