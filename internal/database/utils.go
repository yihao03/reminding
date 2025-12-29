package database

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

func ToPGText(s string) pgtype.Text {
	return pgtype.Text{
		String: s,
		Valid:  true,
	}
}

func ToPGTime(t *time.Time) pgtype.Timestamptz {
	if t == nil {
		return pgtype.Timestamptz{Valid: false}
	}
	return pgtype.Timestamptz{
		Time:  *t,
		Valid: true,
	}
}
