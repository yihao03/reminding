package journal

import (
	"net/http"

	firebase "firebase.google.com/go/v4"
	"github.com/yihao03/reminding/apperrors"
	"github.com/yihao03/reminding/internal/api"
	"github.com/yihao03/reminding/internal/database/sqlc"
	"github.com/yihao03/reminding/internal/router/middleware"
	"github.com/yihao03/reminding/internal/views/journalview"
)

const ErrCreateJournal = "error creating journal"

func HandleCreateJournal(w http.ResponseWriter, r *http.Request, queries *sqlc.Queries, app *firebase.App) error {
	var createParams journalview.CreateView
	if err := api.Decode(r, &createParams); err != nil {
		api.WriteError(http.StatusBadRequest, apperrors.DecodeError(err), w, r.Context())
		return nil
	}

	uid, ok := middleware.GetUserUIDFromContext(r.Context())
	if !ok {
		return apperrors.NewInternalError(nil, "User uid not found in context")
	}

	createObj := createParams.ToCreateJournalParams(uid)
	res, err := queries.CreateJournal(r.Context(), *createObj)
	if err != nil {
		return apperrors.NewInternalError(err, ErrCreateJournal)
	}

	view := journalview.ToReadView(&res)
	api.WriteResponse(view, w)
	return nil
}
