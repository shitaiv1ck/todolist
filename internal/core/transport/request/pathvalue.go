package core_request

import (
	"net/http"
	"strconv"
)

func GetIntPathValue(r *http.Request, key string) (int, error) {
	value := r.PathValue(key)

	num, err := strconv.Atoi(value)
	if err != nil {
		return -1, err
	}

	return num, nil
}
