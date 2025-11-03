package validation

import "errors"

var (
	ErrDescriptionRequired = errors.New("description is required")
)

// ValidateTaskCreation validates required fields for task creation
func ValidateTaskCreation(description string) error {
	if description == "" {
		return ErrDescriptionRequired
	}
	return nil
}
