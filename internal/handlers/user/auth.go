package user

import (
	"net/http"

	firebase "firebase.google.com/go/v4"
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

func AuthorizeUser(w http.ResponseWriter, r *http.Request, queries *sqlc.Queries, app *firebase.App) error {
	var authview userview.AuthView
	if err := api.Decode(r, &authview); err != nil {
		return apperrors.NewInternalError(err, ErrParseAuthView)
	}

	auth, err := app.Auth(r.Context())
	if err != nil {
		return apperrors.NewInternalError(err, ErrGetAuthClient)
	}

	_, err = auth.VerifyIDToken(r.Context(), authview.UserToken)
	if err != nil {
		api.WriteError(http.StatusUnauthorized, apperrors.Wrap(err, ErrInvalidToken), w, r.Context())
		return nil
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "session",
		Value:    authview.UserToken,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
		MaxAge:   60 * 60 * 24 * 5, // 5 days
	})

	return nil
}
