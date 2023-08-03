package object

type SearchRequest struct {
	// Word string
	OnlyMedia bool `db:"only_media"`
	MaxID     int  `db:"max_id"`
	SinceID   int  `db:"since_id"`
	Limit     int  `db:"limit"`
}
