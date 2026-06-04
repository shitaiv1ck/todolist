package web_repository

import (
	"os"

	core_errors "github.com/shitaiv1ck/todolist/internal/core/errors"
)

type WebRepository struct{}

func NewRepository() *WebRepository {
	return &WebRepository{}
}

func (r *WebRepository) GetFile(path string) ([]byte, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, core_errors.ErrInvalidArgument
		}

		return nil, err
	}

	return file, nil
}
