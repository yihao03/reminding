package events

import (
	"net/http"

	firebase "firebase.google.com/go/v4"
	"github.com/yihao03/reminding/apperrors"
	"github.com/yihao03/reminding/internal/api"
	"github.com/yihao03/reminding/internal/database/sqlc"
	"github.com/yihao03/reminding/internal/views/eventview"
)

var (
	ErrCreateEvent     = "error creating event"
	SuccessCreateEvent = "event created successfully"
)

func HandleCreateEvents(w http.ResponseWriter, r *http.Request, queries *sqlc.Queries, app *firebase.App) error {
	var eventParams eventview.EventCreateView
	err := api.Decode(r, &eventParams)
	if err != nil {
		api.WriteError(http.StatusBadRequest, err, w, r.Context())
		return nil
	}

	event, err := queries.CreateEvent(r.Context(), *eventview.ToCreateParams(&eventParams))
	if err != nil {
		return apperrors.NewInternalError(err, ErrCreateEvent)
	}

	api.WriteResponse(event, w, SuccessCreateEvent)
	return nil
}
