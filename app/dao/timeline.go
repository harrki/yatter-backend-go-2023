package dao

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/domain/repository"

	"github.com/jmoiron/sqlx"
)

type (
	// Implementation for repository.Account
	timeline struct {
		db *sqlx.DB
	}
)

func NewTimeline(db *sqlx.DB) repository.Timeline {
	return &timeline{db: db}
}

// Search : SearchRequestを受け取って条件に合致したstatusを取得
func (r *timeline) Search(ctx context.Context, sr *object.SearchRequest) (*[]object.Status, error) {
	entities := make([]object.Status, 0)

	query := "select * from status "
	params := []string{}
	args := make([]interface{}, 0)
	contain_maxid := sr.MaxID >= 0
	contain_sinceid := sr.SinceID >= 0

	if contain_maxid || contain_sinceid {
		query = query + "where "
		if contain_maxid {
			params = append(params, "id <= ?")
			args = append(args, strconv.Itoa(sr.MaxID))
		}
		if contain_sinceid {
			params = append(params, "id >= ?")
			args = append(args, strconv.Itoa(sr.SinceID))
		}
	}

	query = query + strings.Join(params, " and ") + " order by create_at desc limit ?"
	args = append(args, strconv.Itoa(sr.Limit))

	rows, err := r.db.QueryxContext(ctx, query, args...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("failed to find status from db: %w", err)
	}

	for rows.Next() {
		entity := object.Status{}
		err := rows.StructScan(&entity)
		if err != nil {
			return nil, fmt.Errorf("failed to convert status from db to struct: %w", err)
		}
		entities = append(entities, entity)
	}

	return &entities, nil
}
