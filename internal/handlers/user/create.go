package user

import (
	"net/http"

	firebase "firebase.google.com/go/v4"
	"github.com/yihao03/reminding/apperrors"
	"github.com/yihao03/reminding/internal/api"
	database "github.com/yihao03/reminding/internal/database/sqlc"
	"github.com/yihao03/reminding/internal/views/userview"
)

var (
	ErrParseUserView = "Error parsing user view"
	ErrCreateUser    = "Error creating user"
)

func CreateUser(w http.ResponseWriter, r *http.Request, queries *database.Queries, app *firebase.App) error {
	var req userview.CreateUserView
	if err := api.Decode(r, &req); err != nil {
		api.WriteError(http.StatusBadRequest, apperrors.Wrap(err, ErrParseUserView), w, r.Context())
	}

	user, err := queries.CreateUser(r.Context(), *req.ToCreateUserParams())
	if err != nil {
		api.WriteError(http.StatusInternalServerError, apperrors.Wrap(err, ErrCreateUser), w, r.Context())
	}

	api.WriteResponse(user, w, "User created successfully")
	return nil
}
