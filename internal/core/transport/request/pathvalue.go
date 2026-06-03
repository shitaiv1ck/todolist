package core_request

import (
	"fmt"
	"net/http"
	"strconv"

	core_errors "github.com/shitaiv1ck/todolist/internal/core/errors"
)

func GetIntPathValue(r *http.Request, key string) (int, error) {
	value := r.PathValue(key)

	num, err := strconv.Atoi(value)
	if err != nil {
		return -1, fmt.Errorf("%v: %w", err, core_errors.ErrInvalidArgument)
	}

	return num, nil
}
