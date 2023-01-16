// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package gardens

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Garden struct {
	ID            uuid.UUID
	GroupID       uuid.UUID
	Name          string
	Comment       string
	Location      string
	GrowType      string
	GrowSize      string
	GrowStyle     string
	ContainerSize string
	Tags          string
	CreatedAt     time.Time
	UpdatedAt     sql.NullTime
}

type Plant struct {
	ID              uuid.UUID
	GroupID         uuid.UUID
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
	UpdatedAt       sql.NullTime
}

type Recipe struct {
	ID                  uuid.UUID
	Name                string
	Comment             string
	Ingredients         string
	Instructions        string
	Ph                  float64
	MixTimeMilliseconds float64
	Tags                string
	CreatedAt           time.Time
	UpdatedAt           sql.NullTime
}

type Strain struct {
	ID         uuid.UUID
	GroupID    uuid.UUID
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
	UpdatedAt  sql.NullTime
}
