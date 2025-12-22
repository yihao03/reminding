package events

import (
	"net/http"

	firebase "firebase.google.com/go/v4"
	"github.com/yihao03/reminding/internal/database/sqlc"
)

func HandleRegisterEvents(w http.ResponseWriter, r *http.Request, queries *sqlc.Queries, app *firebase.App) error {
	return nil
}
