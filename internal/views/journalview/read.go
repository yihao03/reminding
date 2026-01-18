package journalview

import (
	"time"

	"github.com/yihao03/reminding/internal/database/sqlc"
)

type ReadView struct {
	ID             int32     `json:"id"`
	UpdatedAt      time.Time `json:"updatedAt"`
	CreatedAt      time.Time `json:"createdAt"`
	Title          string    `json:"title"`
	JournalContent string    `json:"journalContent"`
}

func ToReadView(j *sqlc.Journal) *ReadView {
	return &ReadView{
		ID:             j.ID,
		UpdatedAt:      j.UpdatedAt.Time,
		CreatedAt:      j.CreatedAt.Time,
		Title:          j.Title,
		JournalContent: j.JournalContent,
	}
}
