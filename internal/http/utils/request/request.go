package request

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/AbhayBharti21/task-manager/internal/http/utils/logger"
)

func DecodeJSON(r *http.Request, v interface{}) error {
	if r.Body == nil {
		return ErrEmptyBody
	}

	err := json.NewDecoder(r.Body).Decode(v)
	if errors.Is(err, io.EOF) {
		return ErrEmptyBody
	}
	if err != nil {
		logger.Errorf("Error decoding JSON: %v", err)
		return ErrInvalidJSON
	}

	return nil
}

func ValidateRequestBody(r *http.Request, v interface{}, validator func() error) error {
	if err := DecodeJSON(r, v); err != nil {
		return err
	}

	if validator != nil {
		if err := validator(); err != nil {
			return err
		}
	}

	return nil
}
