package dao

import (
	"context"
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

func (r *status) CreateStatus(ctx context.Context, status *object.Status) error {
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
		return err
	}

	return nil
}
