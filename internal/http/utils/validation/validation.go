package validation

import "errors"

var (
	ErrDescriptionRequired = errors.New("description is required")
)

func ValidateTaskCreation(description string) error {
	if description == "" {
		return ErrDescriptionRequired
	}
	return nil
}
