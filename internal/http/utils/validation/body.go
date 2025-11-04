package validation

import (
	"errors"

	types "github.com/AbhayBharti21/task-manager/internal/http/models"
)

var (
	ErrDescriptionRequired = errors.New("description is required")
)

func ValidateTaskCreation(description string) error {
	if description == "" {
		return ErrDescriptionRequired
	}
	return nil
}

func ValidateCreateTaskBody(task types.Task) error {
	return ValidateTaskCreation(task.Description)
}

func ValidateUpdateTaskBody(task types.Task) error {
	if task.OwnerId == 0 {
		return errors.New("ownerId is required for update")
	}
	return nil
}

func ValidateDeleteTaskBody(task types.Task) error {
	if task.OwnerId == 0 {
		return errors.New("ownerId is required for deletion")
	}
	return nil
}
