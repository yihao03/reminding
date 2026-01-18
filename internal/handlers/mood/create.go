// Package mood provides handlers for mood-related operations.
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

const ErrLogMood = "error logging mood"

func HandleLogMood(w http.ResponseWriter, r *http.Request, queries *sqlc.Queries, app *firebase.App) error {
	var logparams moodview.MoodLogView
	if err := api.Decode(r, &logparams); err != nil {
		api.WriteError(http.StatusBadRequest, apperrors.New("mood parameter is required"), w, r.Context())
		return err
	}

	uid, ok := middleware.GetUserUIDFromContext(r.Context())
	if !ok {
		return apperrors.NewInternalError(nil, "User uid not found in context")
	}

	params := sqlc.AddMoodParams{
		UserUid: uid,
		Mood:    logparams.Mood,
	}

	pool, ok := middleware.GetDBPoolFromContext(r.Context())
	if !ok {
		return apperrors.NewInternalError(nil, "database pool not found in context")
	}

	tx, err := pool.Begin(r.Context())
	if err != nil {
		return err
	}
	defer tx.Rollback(r.Context())
	qtx := queries.WithTx(tx)

	_, err = qtx.AddMood(r.Context(), params)
	if err != nil {
		return apperrors.NewInternalError(err, ErrLogMood)
	}

	startTime := time.Now().AddDate(0, 0, -30)

	paramsGet := sqlc.GetMonthlyMoodCountByUserUidParams{
		UserUid:   uid,
		CreatedAt: database.ToPGTime(&startTime),
	}

	count, err := qtx.GetMonthlyMoodCountByUserUid(r.Context(), paramsGet)
	if err != nil {
		return apperrors.NewInternalError(err, "error retrieving monthly mood count")
	}

	if err := tx.Commit(r.Context()); err != nil {
		return apperrors.NewInternalError(err, "error committing mood transaction")
	}

	viewPtr := moodview.ToMoodReadView(true, &count)

	api.WriteResponse(*viewPtr, w)
	return nil
}
