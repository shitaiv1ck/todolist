package core_request

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func DecodeAndValidate(r *http.Request, body any) error {
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return err
	}

	if err := validate.Struct(body); err != nil {
		return err
	}

	return nil
}
