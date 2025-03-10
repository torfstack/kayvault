// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package sqlc

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Secret struct {
	ID            int32
	Value         []byte
	Key           string
	Url           string
	Tags          string
	UserID        int32
	SecretSharing pgtype.Int4
	CreatedAt     pgtype.Timestamp
	UpdatedAt     pgtype.Timestamp
}

type User struct {
	ID        int32
	Subject   string
	Email     string
	FullName  string
	CreatedAt pgtype.Timestamp
	UpdatedAt pgtype.Timestamp
}
