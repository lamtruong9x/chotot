// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0

package repository

import (
	"context"
	"database/sql"
)

type Querier interface {
	CreateUser(ctx context.Context, arg CreateUserParams) (sql.Result, error)
	GetUserByPhone(ctx context.Context, phone string) (*User, error)
}

var _ Querier = (*Queries)(nil)