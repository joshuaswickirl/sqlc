// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0

package querytest

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Address struct {
	ID          int64
	AddressLine string
	Region      string
	City        string
	CreatedAt   pgtype.Timestamptz
	UpdatedAt   pgtype.Timestamptz
}

type User struct {
	ID             int64
	Username       string
	Email          string
	Password       string
	Telephone      int32
	DefaultPayment pgtype.Int8
	CreatedAt      pgtype.Timestamptz
	UpdatedAt      pgtype.Timestamptz
}

type UserAddress struct {
	UserID         int64
	AddressID      int64
	DefaultAddress pgtype.Int8
	CreatedAt      pgtype.Timestamptz
	UpdatedAt      pgtype.Timestamptz
}