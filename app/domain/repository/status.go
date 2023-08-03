package repository

import (
	"context"

	"yatter-backend-go/app/domain/object"
)

type Status interface {
	FindByID(ctx context.Context, id string) (*object.Status, error)
	CreateStatus(ctx context.Context, status *object.Status) (*object.Status, error)
}
