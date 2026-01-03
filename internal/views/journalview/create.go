package journalview

import "github.com/yihao03/reminding/internal/database/sqlc"

type CreateView struct {
	// TODO: check validation length limit
	Title          string `json:"title" validate:"required,le=255"`
	JournalContent string `json:"journalContent" validate:"required"`
}

func (v *CreateView) ToCreateJournalParams(uid string) *sqlc.CreateJournalParams {
	return &sqlc.CreateJournalParams{
		UserUid:        uid,
		Title:          v.Title,
		JournalContent: v.JournalContent,
	}
}
