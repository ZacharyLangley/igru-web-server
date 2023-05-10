package database

import (
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

func NewFromNullableUUID(in uuid.NullUUID) pgtype.UUID {
	if !in.Valid {
		return pgtype.UUID{
			Valid: false,
		}
	}
	return pgtype.UUID{
		Bytes: in.UUID,
		Valid: true,
	}
}

func NewFromUUID(in uuid.UUID) pgtype.UUID {
	return pgtype.UUID{
		Bytes: in,
		Valid: true,
	}
}
