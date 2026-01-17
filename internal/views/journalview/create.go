package journalview

import "github.com/yihao03/reminding/internal/database/sqlc"

type CreateView struct {
	// TODO: check validation length limit
	Title   string `json:"title" validate:"required,max=255"`
	Content string `json:"content" validate:"required"`
}

func (v *CreateView) ToCreateJournalParams(uid string) *sqlc.CreateJournalParams {
	return &sqlc.CreateJournalParams{
		UserUid:        uid,
		Title:          v.Title,
		JournalContent: v.Content,
	}
}
