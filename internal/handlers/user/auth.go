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
	ErrGetAuthClient = "failed to get firebase auth client"
	ErrInvalidToken  = "Token invalid"
)

func HandleAuthorizeUser(w http.ResponseWriter, r *http.Request, queries *sqlc.Queries, app *firebase.App) error {
	var authview userview.AuthView
	if err := api.Decode(r, &authview); err != nil {
		api.WriteError(http.StatusBadRequest,
			apperrors.DecodeError(err), w, r.Context())
		return nil
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

	user, err := queries.GetUserByUid(r.Context(), token.UID)
	if err != nil {
		api.WriteError(http.StatusInternalServerError, apperrors.Wrap(err, "failed to get user by uid"), w, r.Context())
		return nil
	}

	view := userview.ToUserView(&user)

	api.WriteResponse(view, w)
	return nil
}
