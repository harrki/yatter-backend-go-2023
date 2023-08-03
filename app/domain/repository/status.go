package repository

import (
	"context"

	"yatter-backend-go/app/domain/object"
)

type Status interface {
	CreateStatus(ctx context.Context, status *object.Status) error
	FindByID(ctx context.Context, id string) (*object.Status, error)
}
