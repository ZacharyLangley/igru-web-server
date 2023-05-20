// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0

package authentication

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Group struct {
	ID        uuid.UUID
	Name      string
	UserGroup bool
	CreatedAt time.Time
	UpdatedAt sql.NullTime
}

type GroupMember struct {
	UserID    uuid.UUID
	GroupID   uuid.UUID
	Role      int32
	CreatedAt time.Time
	UpdatedAt sql.NullTime
}

type Session struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	CreatedAt time.Time
	ExpiredAt time.Time
}

type User struct {
	ID         uuid.UUID
	Email      string
	GroupID    uuid.UUID
	FullName   sql.NullString
	GlobalRole sql.NullInt32
	Active     sql.NullBool
	Salt       string
	Hash       string
	CreatedAt  time.Time
	UpdatedAt  sql.NullTime
}
