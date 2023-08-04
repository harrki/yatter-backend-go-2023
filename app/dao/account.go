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
	account struct {
		db *sqlx.DB
	}
)

// Create accout repository
func NewAccount(db *sqlx.DB) repository.Account {
	return &account{db: db}
}

// FindByUsername : ユーザ名からユーザを取得
func (r *account) FindByUsername(ctx context.Context, username string) (*object.Account, error) {
	entity := new(object.Account)
	err := r.db.QueryRowxContext(ctx, "SELECT * FROM account WHERE username = ?", username).StructScan(entity)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("failed to find account from db: %w", err)
	}

	return entity, nil
}

// CreateUser : ユーザを追加
func (r *account) CreateUser(ctx context.Context, account *object.Account) (*object.Account, error) {
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

	if _, err = r.db.NamedExec("INSERT INTO account (username, password_hash) VALUES (:username, :password_hash)", account); err != nil {
		return nil, err
	}

	entity := new(object.Account)
	err_get := r.db.QueryRowxContext(ctx, "SELECT * FROM account WHERE id = LAST_INSERT_ID()").StructScan(entity)
	if err_get != nil {
		return nil, err_get
	}

	return entity, nil
}

func (r *account) UpdateCredentials(ctx context.Context, account *object.Account) (*object.Account, error) {
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

	fmt.Println(account.ID)
	if _, err = r.db.NamedExec("UPDATE account SET display_name = :display_name, avatar = :avatar, header = :header, note = :note WHERE id = :id", account); err != nil {
		return nil, err
	}

	entity := new(object.Account)
	err_get := r.db.QueryRowxContext(ctx, "SELECT * FROM account WHERE id = ?", account.ID).StructScan(entity)
	if err_get != nil {
		return nil, err_get
	}

	return entity, nil
}
