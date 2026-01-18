package mood

import (
	"net/http"
	"time"

	firebase "firebase.google.com/go/v4"
	"github.com/yihao03/reminding/apperrors"
	"github.com/yihao03/reminding/internal/api"
	"github.com/yihao03/reminding/internal/database"
	"github.com/yihao03/reminding/internal/database/sqlc"
	"github.com/yihao03/reminding/internal/router/middleware"
	"github.com/yihao03/reminding/internal/views/moodview"
)

func HandleCheckMood(w http.ResponseWriter, r *http.Request, queries *sqlc.Queries, app *firebase.App) error {
	uid, ok := middleware.GetUserUIDFromContext(r.Context())
	if !ok {
		return apperrors.NewInternalError(nil, "User uid not found in context")
	}

	logged, err := queries.CheckUserLoggedMoodToday(r.Context(), uid)
	if err != nil {
		return apperrors.NewInternalError(err, "Failed to check if user logged mood today")
	}

	if !logged {
		api.WriteResponse(*moodview.ToMoodReadView(logged, nil), w)
		return nil
	}

	startTime := time.Now().AddDate(0, 0, -30)
	paramsGet := sqlc.GetMonthlyMoodCountByUserUidParams{
		UserUid:   uid,
		CreatedAt: database.ToPGTime(&startTime),
	}

	moodCount, err := queries.GetMonthlyMoodCountByUserUid(r.Context(), paramsGet)
	if err != nil {
		return apperrors.NewInternalError(err, "Failed to get monthly mood count")
	}

	api.WriteResponse(*moodview.ToMoodReadView(logged, &moodCount), w)
	return nil
}
