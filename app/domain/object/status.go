package object

import (
	"time"
)

type Status struct {
	ID        int64     `json:"id,omitempty" db:"id"`
	AccountID int64     `json:"account_id,omitempty" db:"account_id"`
	Content   string    `json:"content,omitempty" db:"content"`
	URL       string    `json:"url,omitempty" db:"url"`
	CreateAt  time.Time `json:"create_at,omitempty" db:"create_at"`
}
