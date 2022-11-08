// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: query.sql

package authentication

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createSession = `-- name: CreateSession :one
INSERT INTO sessions (
  user_id, created_at, expired_at
) VALUES (
  $1, $2, $3
)
RETURNING id
`

type CreateSessionParams struct {
	UserID    uuid.UUID
	CreatedAt time.Time
	ExpiredAt time.Time
}

func (q *Queries) CreateSession(ctx context.Context, arg CreateSessionParams) (uuid.UUID, error) {
	row := q.db.QueryRowContext(ctx, createSession, arg.UserID, arg.CreatedAt, arg.ExpiredAt)
	var id uuid.UUID
	err := row.Scan(&id)
	return id, err
}

const createUser = `-- name: CreateUser :one
INSERT INTO users (
  email, first_name, last_name, salt, hash
) VALUES (
  $1, $2, $3, $4, $5
)
RETURNING id, email, first_name, last_name, active, salt, hash, created_at, updated_at, deleted_at
`

type CreateUserParams struct {
	Email     string
	FirstName string
	LastName  string
	Salt      string
	Hash      string
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.Email,
		arg.FirstName,
		arg.LastName,
		arg.Salt,
		arg.Hash,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.FirstName,
		&i.LastName,
		&i.Active,
		&i.Salt,
		&i.Hash,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const deleteSession = `-- name: DeleteSession :exec
DELETE FROM sessions
WHERE id = $1
`

func (q *Queries) DeleteSession(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteSession, id)
	return err
}

const getExpiredSessions = `-- name: GetExpiredSessions :many
SELECT id, user_id, created_at, expired_at FROM sessions
WHERE expired_at < NOW() LIMIT 100
`

func (q *Queries) GetExpiredSessions(ctx context.Context) ([]Session, error) {
	rows, err := q.db.QueryContext(ctx, getExpiredSessions)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Session
	for rows.Next() {
		var i Session
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.CreatedAt,
			&i.ExpiredAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getSessionUserID = `-- name: GetSessionUserID :one
SELECT user_id, expired_at FROM sessions
WHERE id = $1 LIMIT 1
`

type GetSessionUserIDRow struct {
	UserID    uuid.UUID
	ExpiredAt time.Time
}

func (q *Queries) GetSessionUserID(ctx context.Context, id uuid.UUID) (GetSessionUserIDRow, error) {
	row := q.db.QueryRowContext(ctx, getSessionUserID, id)
	var i GetSessionUserIDRow
	err := row.Scan(&i.UserID, &i.ExpiredAt)
	return i, err
}

const getUser = `-- name: GetUser :one
SELECT id, email, first_name, last_name, active, salt, hash, created_at, updated_at, deleted_at FROM users
WHERE email = $1 LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUser, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.FirstName,
		&i.LastName,
		&i.Active,
		&i.Salt,
		&i.Hash,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const getUserByID = `-- name: GetUserByID :one
SELECT id, email, first_name, last_name, active, salt, hash, created_at, updated_at, deleted_at FROM users
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetUserByID(ctx context.Context, id uuid.NullUUID) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByID, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.FirstName,
		&i.LastName,
		&i.Active,
		&i.Salt,
		&i.Hash,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}
