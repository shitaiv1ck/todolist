package core_errors

import "errors"

var (
	ErrInvalidArgument = errors.New("invalid argument")
	ErrConflict        = errors.New("already exists")
)
