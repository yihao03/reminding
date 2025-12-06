package api

import (
	"encoding/json"
	"net/http"
)

func Decode(r *http.Request, v any) error {
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(v); err != nil {
		return err
	}

	return nil
}
