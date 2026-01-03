// Package user contains handlers related to user operations.
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
	ErrParseUserView = "Error parsing user view"
	ErrCreateUser    = "Error creating user"
)

func CreateUser(w http.ResponseWriter, r *http.Request, queries *sqlc.Queries, app *firebase.App) error {
	var req userview.CreateUserView
	if err := api.Decode(r, &req); err != nil {
		api.WriteError(http.StatusBadRequest, apperrors.DecodeError(err), w, r.Context())
		return nil
	}

	user, err := queries.CreateUser(r.Context(), *req.ToCreateUserParams())
	if err != nil {
		return apperrors.Wrap(err, ErrCreateUser)
	}

	api.WriteResponse(user, w, "User created successfully")
	return nil
}
