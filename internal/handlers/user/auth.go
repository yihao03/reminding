package user

import (
	"errors"
	"net/http"

	firebase "firebase.google.com/go/v4"
	"github.com/jackc/pgx/v5"
	"github.com/yihao03/reminding/apperrors"
	"github.com/yihao03/reminding/internal/api"
	"github.com/yihao03/reminding/internal/database/sqlc"
	"github.com/yihao03/reminding/internal/views/userview"
)

const (
	ErrParseAuthView = "Error parsing auth view"
	ErrGetAuthClient = "failed to get firebase auth client"
	ErrInvalidToken  = "Token invalid"
)

func HandleAuthorizeUser(w http.ResponseWriter, r *http.Request, queries *sqlc.Queries, app *firebase.App) error {
	var authview userview.AuthView
	if err := api.Decode(r, &authview); err != nil {
		return apperrors.NewInternalError(err, ErrParseAuthView)
	}

	auth, err := app.Auth(r.Context())
	if err != nil {
		return apperrors.NewInternalError(err, ErrGetAuthClient)
	}

	token, err := auth.VerifyIDToken(r.Context(), authview.UserToken)
	if err != nil {
		api.WriteError(http.StatusUnauthorized, apperrors.Wrap(err, ErrInvalidToken), w, r.Context())
		return nil
	}

	userParams := sqlc.CreateUserIfAbsentParams{
		FirebaseUid: token.UID,
		DisplayName: authview.User.Name,
		Email:       authview.User.Email,
	}

	_, err = queries.CreateUserIfAbsent(r.Context(), userParams)
	if err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			return apperrors.Wrap(err, "Failed to create user if absent")
		}
	}

	api.WriteResponse(map[string]string{"status": "authorized"}, w)
	return nil
}
