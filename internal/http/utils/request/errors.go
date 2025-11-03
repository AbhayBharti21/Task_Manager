package request

import "errors"

var (
	ErrEmptyBody   = errors.New("empty request body")
	ErrInvalidJSON = errors.New("invalid JSON format")
)
