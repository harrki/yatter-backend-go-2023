package repository

import (
	"context"

	"yatter-backend-go/app/domain/object"
)

type Timeline interface {
	Search(ctx context.Context, sr *object.SearchRequest) (*[]object.Status, error)
}
