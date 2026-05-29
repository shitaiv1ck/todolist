package core_response

import (
	"encoding/json"
	"errors"
	"net/http"

	core_errors "github.com/shitaiv1ck/todolist/internal/core/errors"
)

type ResponseHandler struct {
	w http.ResponseWriter
}

func NewResponseHandler(w http.ResponseWriter) *ResponseHandler {
	return &ResponseHandler{
		w: w,
	}
}

func (r *ResponseHandler) JsonResponse(responseBody any, statusCode int) {
	r.w.WriteHeader(statusCode)

	if err := json.NewEncoder(r.w).Encode(responseBody); err != nil {
		panic(err)
	}
}

func (r *ResponseHandler) ErrorResponse(msg string, err error) {
	statusCode := setErrStatusCode(err)

	r.w.WriteHeader(statusCode)

	errResponse := ErrorDTO{
		Message: msg,
		Error:   err.Error(),
	}

	if err := json.NewEncoder(r.w).Encode(errResponse); err != nil {
		panic(err)
	}
}

func setErrStatusCode(err error) int {
	if errors.Is(err, core_errors.ErrInvalidArgument) {
		return http.StatusBadRequest
	}
	if errors.Is(err, core_errors.ErrConflict) {
		return http.StatusConflict
	}

	return http.StatusInternalServerError
}
