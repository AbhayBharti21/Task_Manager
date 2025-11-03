package path

import "errors"

var (
	ErrInvalidPath = errors.New("invalid path: path parameter not found")
	ErrInvalidID   = errors.New("invalid ID: unable to convert path parameter to integer")
)
