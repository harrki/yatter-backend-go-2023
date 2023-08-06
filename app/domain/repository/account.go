package repository

import (
	"context"

	"yatter-backend-go/app/domain/object"
)

type Account interface {
	// Fetch account which has specified username
	FindByUsername(ctx context.Context, username string) (*object.Account, error)
	CreateUser(ctx context.Context, account *object.Account) (*object.Account, error)
	UpdateCredentials(ctx context.Context, account *object.Account) (*object.Account, error)
	// TODO: Add Other APIs
}
