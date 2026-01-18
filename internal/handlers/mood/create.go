package mood

import (
	"net/http"
	"strconv"

	firebase "firebase.google.com/go/v4"
	"github.com/go-chi/chi/v5"
	"github.com/yihao03/reminding/apperrors"
	"github.com/yihao03/reminding/internal/api"
	"github.com/yihao03/reminding/internal/database/sqlc"
	"github.com/yihao03/reminding/internal/router/middleware"
	"github.com/yihao03/reminding/internal/views/moodview"
)

const ErrLogMood = "error logging mood"

func HandleLogMood(w http.ResponseWriter, r *http.Request, queries *sqlc.Queries, app *firebase.App) error {
	mood := chi.URLParam(r, "mood")
	if mood == "" {
		api.WriteError(http.StatusBadRequest, apperrors.New("mood parameter is required"), w, r.Context())
	}

	// Mood is stored as int enum in the database
	// the moods are currently hardcoded in the frontned
	moodEnum, err := strconv.Atoi(mood)
	if err != nil {
		return apperrors.NewInternalError(err, "invalid mood parameter")
	}

	uid, ok := middleware.GetUserUIDFromContext(r.Context())
	if !ok {
		return apperrors.NewInternalError(nil, "User uid not found in context")
	}

	params := sqlc.AddMoodParams{
		UserUid: uid,
		Mood:    int32(moodEnum),
	}

	_, err = queries.AddMood(r.Context(), params)
	if err != nil {
		return apperrors.NewInternalError(err, ErrLogMood)
	}

	count, err := queries.GetMonthlyMoodCountByUserUid(r.Context(), uid)
	if err != nil {
		return apperrors.NewInternalError(err, "error retrieving monthly mood count")
	}

	view := moodview.ToMoodReadViewArray(count)

	api.WriteResponse(view, w)
	return nil
}
