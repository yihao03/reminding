package journalview

import (
	"time"

	"github.com/yihao03/reminding/internal/database/sqlc"
)

type ListView struct {
	ID        int32     `json:"id"`
	UpdatedAt time.Time `json:"updatedAt"`
	CreatedAt time.Time `json:"createdAt"`
	Title     string    `json:"title"`
}

func ToListView(j *sqlc.ListJournalsRow) *ListView {
	return &ListView{
		ID:        j.ID,
		UpdatedAt: j.UpdatedAt.Time,
		CreatedAt: j.CreatedAt.Time,
		Title:     j.Title,
	}
}

func ToListViewList(journals []sqlc.ListJournalsRow) *[]ListView {
	res := make([]ListView, len(journals))
	for i, j := range journals {
		res[i] = *ToListView(&j)
	}
	return &res
}
