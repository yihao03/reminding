package events

import (
	"net/http"

	firebase "firebase.google.com/go/v4"
	"github.com/yihao03/reminding/apperrors"
	"github.com/yihao03/reminding/internal/api"
	"github.com/yihao03/reminding/internal/database/sqlc"
	"github.com/yihao03/reminding/internal/router/middleware"
	"github.com/yihao03/reminding/internal/views/eventview"
)

var (
	ErrGetEvents     = "Error getting events"
	SuccessGetEvents = "Events retrieved successfully"
)

func HandleGetEvents(w http.ResponseWriter, r *http.Request, queries *sqlc.Queries, app *firebase.App) error {
	uid := middleware.GetUserIDFromContext(r.Context())

	events, err := queries.ListEventsWithUserRegistration(r.Context(), uid)
	if err != nil {
		api.WriteError(http.StatusInternalServerError, apperrors.Wrap(err, ErrGetEvents), w, r.Context())
		return nil
	}

	view := eventview.ToEventListView(&events)

	api.WriteResponse(view, w, SuccessGetEvents)
	return nil
}
