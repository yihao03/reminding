package api

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func Decode(r *http.Request, v any, except ...string) error {
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(v); err != nil {
		return err
	}

	validate := validator.New(validator.WithRequiredStructEnabled())

	return validate.StructExcept(v, except...)
}
