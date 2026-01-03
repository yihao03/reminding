package event

import (
	"net/http"
	"strconv"

	firebase "firebase.google.com/go/v4"
	"github.com/go-chi/chi/v5"
	"github.com/yihao03/reminding/apperrors"
	"github.com/yihao03/reminding/internal/api"
	"github.com/yihao03/reminding/internal/database/sqlc"
	"github.com/yihao03/reminding/internal/router/middleware"
	"github.com/yihao03/reminding/internal/views/eventview"
)

var ErrMissingIDParam = "missing id parameter"

func HandleReadEvents(w http.ResponseWriter, r *http.Request, queries *sqlc.Queries, app *firebase.App) error {
	id := chi.URLParam(r, "id")
	if id == "" {
		api.WriteError(http.StatusBadRequest, apperrors.Wrap(nil, ErrMissingIDParam), w, r.Context())
		return nil
	}

	intID, err := strconv.Atoi(id)
	if err != nil {
		api.WriteError(http.StatusBadRequest, apperrors.Wrap(err, ErrMissingIDParam), w, r.Context())
		return nil
	}

	uid, ok := middleware.GetUserIDFromContext(r.Context())
	if !ok {
		api.WriteError(http.StatusUnauthorized, apperrors.New(ErrFailedToGetUserUID), w, r.Context())
		return nil
	}

	params := sqlc.GetEventByIdAndUidParams{
		ID:      int32(intID),
		UserUid: uid,
	}

	event, err := queries.GetEventByIdAndUid(r.Context(), params)
	if err != nil {
		return apperrors.NewInternalError(err, "failed to get event by id")
	}

	eventView := eventview.ToDetailedEventView(&event)
	api.WriteResponse(eventView, w)
	return nil
}
