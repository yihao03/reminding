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

const ErrListJournals = "error listing journals"

func HandleListJournals(w http.ResponseWriter, r *http.Request, queries *sqlc.Queries, app *firebase.App) error {
	uid, ok := middleware.GetUserUIDFromContext(r.Context())
	if !ok {
		return apperrors.NewInternalError(nil, "User uid not found in context")
	}

	journals, err := queries.ListJournals(r.Context(), uid)
	if err != nil {
		return apperrors.NewInternalError(err, ErrListJournals)
	}

	views := journalview.ToListViewList(journals)

	api.WriteResponse(views, w)
	return nil
}
