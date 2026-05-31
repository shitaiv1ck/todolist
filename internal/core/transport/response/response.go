package core_response

import (
	"encoding/json"
	"errors"
	"net/http"

	core_errors "github.com/shitaiv1ck/todolist/internal/core/errors"
)

type ResponseHandler struct {
	http.ResponseWriter
	statusCode int
}

func NewResponseHandler(w http.ResponseWriter) *ResponseHandler {
	return &ResponseHandler{
		ResponseWriter: w,
	}
}

func (rw *ResponseHandler) JsonResponse(responseBody any, statusCode int) {
	rw.statusCode = statusCode

	rw.WriteHeader(rw.statusCode)

	if err := json.NewEncoder(rw).Encode(responseBody); err != nil {
		panic(err)
	}
}

func (rw *ResponseHandler) ErrorResponse(msg string, err error) {
	rw.statusCode = setErrStatusCode(err)

	rw.WriteHeader(rw.statusCode)

	errResponse := ErrorDTO{
		Message: msg,
		Error:   err.Error(),
	}

	if err := json.NewEncoder(rw).Encode(errResponse); err != nil {
		panic(err)
	}
}

func (rw *ResponseHandler) GetStatusCode() int {
	return rw.statusCode
}

func setErrStatusCode(err error) int {
	if errors.Is(err, core_errors.ErrInvalidArgument) {
		return http.StatusBadRequest
	}
	if errors.Is(err, core_errors.ErrConflict) {
		return http.StatusConflict
	}
	if errors.Is(err, core_errors.ErrNotFound) {
		return http.StatusNotFound
	}
	if errors.Is(err, core_errors.ErrUnautorize) {
		return http.StatusUnauthorized
	}

	return http.StatusInternalServerError
}
