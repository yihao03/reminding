package event

import (
	"errors"
	"net/http"
	"strconv"

	firebase "firebase.google.com/go/v4"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/yihao03/reminding/apperrors"
	"github.com/yihao03/reminding/internal/api"
	"github.com/yihao03/reminding/internal/database/sqlc"
	"github.com/yihao03/reminding/internal/router/middleware"
)

const (
	ErrFailedToGetUserUID  = "Failed to retrieve user UID from context"
	ErrInvalidEventIDParam = "Invalid event ID parameter"
	ErrUserRegistered      = "User already registered for this event"
)

func HandleRegisterEvents(w http.ResponseWriter, r *http.Request, queries *sqlc.Queries, app *firebase.App) error {
	eventID := chi.URLParam(r, "id")
	if eventID == "" {
		api.WriteError(http.StatusBadRequest, apperrors.Wrap(nil, ErrMissingIDParam), w, r.Context())
		return nil
	}

	eventIDInt, err := strconv.Atoi(eventID)
	if err != nil {
		api.WriteError(http.StatusBadRequest, apperrors.Wrap(err, ErrInvalidEventIDParam), w, r.Context())
		return nil
	}

	userUID, ok := middleware.GetUserIDFromContext(r.Context())
	if !ok {
		api.WriteError(http.StatusUnauthorized, apperrors.New(ErrFailedToGetUserUID), w, r.Context())
		return nil
	}

	reg, err := queries.RegisterEvent(r.Context(), sqlc.RegisterEventParams{EventID: int32(eventIDInt), UserUid: userUID})
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == pgerrcode.UniqueViolation {
			api.WriteError(http.StatusConflict, apperrors.New(ErrUserRegistered), w, r.Context())
			return nil
		}
		return apperrors.Wrap(err, "Failed to register for event")
	}

	api.WriteResponse(reg, w)
	return nil
}
