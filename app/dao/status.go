package dao

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/domain/repository"

	"github.com/jmoiron/sqlx"
)

type (
	// Implementation for repository.Account
	status struct {
		db *sqlx.DB
	}
)

func NewStatus(db *sqlx.DB) repository.Status {
	return &status{db: db}
}

func (r *status) FindByID(ctx context.Context, id string) (*object.Status, error) {
	entity := new(object.Status)
	err := r.db.QueryRowxContext(ctx, "SELECT * FROM status WHERE id = ?", id).StructScan(entity)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("failed to find status from db: %w", err)
	}

	return entity, nil
}

func (r *status) CreateStatus(ctx context.Context, status *object.Status) (*object.Status, error) {
	tx, _ := r.db.Beginx()
	var err error
	defer func() {
		switch r := recover(); {
		case r != nil:
			tx.Rollback()
			panic(r)
		case err != nil:
			tx.Rollback()
		}
	}()

	if _, err = r.db.NamedExec("INSERT INTO status (account_id, content, url) VALUES (:account_id, :content, :url)", status); err != nil {
		return nil, err
	}

	entity := new(object.Status)
	err_get := r.db.QueryRowxContext(ctx, "SELECT * FROM status WHERE id = LAST_INSERT_ID()").StructScan(entity)
	if err_get != nil {
		return nil, err_get
	}

	return entity, nil
}
