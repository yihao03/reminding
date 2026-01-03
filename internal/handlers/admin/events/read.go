package events

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	firebase "firebase.google.com/go/v4"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"
	"github.com/yihao03/reminding/apperrors"
	"github.com/yihao03/reminding/internal/api"
	"github.com/yihao03/reminding/internal/database/sqlc"
	"github.com/yihao03/reminding/internal/views/eventview"
)

const (
	ErrMissingIDParam = "missing id parameter"
	ErrEventNotFound  = "event with id %d not found"
)

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

	event, err := queries.GetEventById(r.Context(), int32(intID))
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			api.WriteError(http.StatusNotFound, apperrors.New(fmt.Sprintf(ErrEventNotFound, intID)), w, r.Context())
			return nil
		}
		return apperrors.NewInternalError(err, "failed to get event by id")
	}

	users, err := queries.GetEventRegisteredUsers(r.Context(), int32(intID))
	if err != nil {
		return apperrors.NewInternalError(err, "failed to get event registered users")
	}

	eventView := eventview.ToAdminEventView(&event, &users)
	api.WriteResponse(eventView, w)
	return nil
}
