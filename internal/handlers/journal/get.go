package journal

import (
	"database/sql"
	"errors"
	"net/http"
	"strconv"

	firebase "firebase.google.com/go/v4"
	"github.com/go-chi/chi/v5"
	"github.com/yihao03/reminding/apperrors"
	"github.com/yihao03/reminding/internal/api"
	"github.com/yihao03/reminding/internal/database/sqlc"
	"github.com/yihao03/reminding/internal/router/middleware"
	"github.com/yihao03/reminding/internal/views/journalview"
)

const (
	ErrGetJournal      = "error getting journal"
	ErrInvalidID       = "invalid journal ID"
	ErrJournalNotFound = "journal not found"
)

func HandleGetJournal(w http.ResponseWriter, r *http.Request, queries *sqlc.Queries, app *firebase.App) error {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		api.WriteError(http.StatusBadRequest, apperrors.Wrap(err, ErrInvalidID), w, r.Context())
		return nil
	}

	uid, ok := middleware.GetUserUIDFromContext(r.Context())
	if !ok {
		return apperrors.NewInternalError(nil, "User uid not found in context")
	}

	journal, err := queries.GetJournal(r.Context(), sqlc.GetJournalParams{
		ID:      int32(id),
		UserUid: uid,
	})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			api.WriteError(http.StatusNotFound, apperrors.Wrap(err, ErrJournalNotFound), w, r.Context())
			return nil
		}
		return apperrors.NewInternalError(err, ErrGetJournal)
	}

	view := journalview.ToReadView(&journal)
	api.WriteResponse(view, w)
	return nil
}
