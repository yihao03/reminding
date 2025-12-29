package events

import (
	"net/http"

	firebase "firebase.google.com/go/v4"
	"github.com/yihao03/reminding/apperrors"
	"github.com/yihao03/reminding/internal/api"
	"github.com/yihao03/reminding/internal/database/sqlc"
	"github.com/yihao03/reminding/internal/views/eventview"
)

const (
	ErrGetEvents     = "Error getting events"
	SuccessGetEvents = "Events retrieved successfully"
)

func HandleListEvents(w http.ResponseWriter, r *http.Request, queries *sqlc.Queries, app *firebase.App) error {
	events, err := queries.ListEventsAdmin(r.Context())
	if err != nil {
		return apperrors.Wrap(err, ErrGetEvents)
	}

	view := eventview.ToEventListAdminView(&events)

	api.WriteResponse(view, w, SuccessGetEvents)
	return nil
}
