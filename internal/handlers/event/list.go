package event

import (
	"net/http"

	firebase "firebase.google.com/go/v4"
	"github.com/yihao03/reminding/apperrors"
	"github.com/yihao03/reminding/internal/api"
	"github.com/yihao03/reminding/internal/database/sqlc"
	"github.com/yihao03/reminding/internal/router/middleware"
	"github.com/yihao03/reminding/internal/views/eventview"
)

const (
	ErrGetEvents     = "Error getting events"
	SuccessGetEvents = "Events retrieved successfully"
)

func HandleGetEvents(w http.ResponseWriter, r *http.Request, queries *sqlc.Queries, app *firebase.App) error {
	uid, ok := middleware.GetUserIDFromContext(r.Context())
	if !ok {
		api.WriteError(http.StatusBadRequest, apperrors.New("User ID not found in context"), w, r.Context())
		return nil
	}

	events, err := queries.ListEventsUser(r.Context(), uid)
	if err != nil {
		return apperrors.Wrap(err, ErrGetEvents)
	}

	view := eventview.ToUserEventList(&events)

	api.WriteResponse(view, w, SuccessGetEvents)
	return nil
}
