package handlers

import (
	"fmt"
	"net/http"

	types "github.com/AbhayBharti21/task-manager/internal/http/models"
	"github.com/AbhayBharti21/task-manager/internal/http/utils/logger"
	"github.com/AbhayBharti21/task-manager/internal/http/utils/path"
	requestutil "github.com/AbhayBharti21/task-manager/internal/http/utils/request"
	"github.com/AbhayBharti21/task-manager/internal/http/utils/response"
	"github.com/AbhayBharti21/task-manager/internal/http/utils/validation"
	"github.com/AbhayBharti21/task-manager/internal/repository"
)

var taskRepo *repository.TaskRepository

func init() {
	taskRepo = repository.NewTaskRepository()
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	var task types.Task

	validator := func() error {
		return validation.ValidateCreateTaskBody(task)
	}

	if err := requestutil.ValidateRequestBody(r, &task, validator); err != nil {
		if err == requestutil.ErrEmptyBody {
			logger.Error("Empty request body")
			response.WriteError(w, http.StatusBadRequest, "empty body")
			return
		}
		if err == requestutil.ErrInvalidJSON {
			logger.Errorf("Error decoding JSON: %v", err)
			response.WriteError(w, http.StatusBadRequest, "invalid JSON format")
			return
		}
		// Validation error
		logger.Errorf("Validation error: %v", err)
		response.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	createdTask := taskRepo.Create(task)
	msg := fmt.Sprintf("Task created successfully with task Id %d", createdTask.TaskId)
	logger.Infof("Task created successfully with task ID %d", createdTask.TaskId)

	response.WriteSuccess(w, http.StatusCreated, map[string]interface{}{
		"message": msg,
		"task":    createdTask,
	})
}

func GetTask(w http.ResponseWriter, r *http.Request) {
	id, err := path.ExtractTaskID(r)
	if err != nil {
		logger.Errorf("Error extracting task ID: %v", err)
		response.WriteError(w, http.StatusBadRequest, "invalid task ID")
		return
	}

	task, err := taskRepo.GetByID(id)
	if err != nil {
		if err == repository.ErrTaskNotFound {
			logger.Warnf("Task ID %d not found", id)
			response.WriteError(w, http.StatusNotFound, "task not found")
			return
		}
		logger.Errorf("Error getting task: %v", err)
		response.WriteError(w, http.StatusInternalServerError, "failed to retrieve task")
		return
	}

	logger.Infof("Task with ID %d retrieved successfully", id)
	response.WriteSuccessWithData(w, http.StatusOK, map[string]interface{}{"task": task})
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	id, err := path.ExtractTaskID(r)
	if err != nil {
		logger.Errorf("Error extracting task ID: %v", err)
		response.WriteError(w, http.StatusBadRequest, "invalid task ID")
		return
	}

	var updateFields types.Task
	validator := func() error {
		return validation.ValidateUpdateTaskBody(updateFields)
	}

	if err := requestutil.ValidateRequestBody(r, &updateFields, validator); err != nil {
		if err == requestutil.ErrEmptyBody {
			logger.Error("Empty request body")
			response.WriteError(w, http.StatusBadRequest, "empty body")
			return
		}
		if err == requestutil.ErrInvalidJSON {
			logger.Errorf("Error decoding update fields: %v", err)
			response.WriteError(w, http.StatusBadRequest, "invalid JSON format")
			return
		}
		// Validation error
		logger.Errorf("Validation error: %v", err)
		response.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	// Verify ownership before updating
	if err := taskRepo.VerifyOwner(id, updateFields.OwnerId); err != nil {
		if err == repository.ErrUnauthorized {
			logger.Warnf("Unauthorized access attempt for task ID %d by owner ID %d", id, updateFields.OwnerId)
			response.WriteError(w, http.StatusUnauthorized, "unauthorized")
			return
		}
		if err == repository.ErrTaskNotFound {
			logger.Warnf("Task ID %d not found for update", id)
			response.WriteError(w, http.StatusNotFound, "task not found")
			return
		}
	}

	// Get existing task to preserve fields that aren't being updated
	existingTask, _ := taskRepo.GetByID(id)

	// Only update if new values are provided
	if updateFields.Description == "" {
		updateFields.Description = existingTask.Description
	}
	if !updateFields.IsCompleted {
		updateFields.IsCompleted = existingTask.IsCompleted
	}

	updatedTask, err := taskRepo.Update(id, updateFields)
	if err != nil {
		logger.Errorf("Error updating task ID %d: %v", id, err)
		response.WriteError(w, http.StatusInternalServerError, "failed to update task")
		return
	}

	logger.Infof("Task with ID %d successfully updated", updatedTask.TaskId)
	response.WriteSuccessWithData(w, http.StatusOK, map[string]interface{}{"task": updatedTask})
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	id, err := path.ExtractTaskID(r)
	if err != nil {
		logger.Errorf("Error extracting task ID: %v", err)
		response.WriteError(w, http.StatusBadRequest, "invalid task ID")
		return
	}

	var deleteRequest types.Task
	validator := func() error {
		return validation.ValidateDeleteTaskBody(deleteRequest)
	}

	if err := requestutil.ValidateRequestBody(r, &deleteRequest, validator); err != nil {
		if err == requestutil.ErrEmptyBody {
			logger.Error("Empty request body")
			response.WriteError(w, http.StatusBadRequest, "empty body")
			return
		}
		if err == requestutil.ErrInvalidJSON {
			logger.Errorf("Error decoding delete request: %v", err)
			response.WriteError(w, http.StatusBadRequest, "invalid JSON format")
			return
		}
		// Validation error
		logger.Errorf("Validation error: %v", err)
		response.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	// Verify ownership before deleting
	if err := taskRepo.VerifyOwner(id, deleteRequest.OwnerId); err != nil {
		if err == repository.ErrUnauthorized {
			logger.Warnf("Unauthorized delete attempt for task ID %d by owner ID %d", id, deleteRequest.OwnerId)
			response.WriteError(w, http.StatusUnauthorized, "unauthorized")
			return
		}
		if err == repository.ErrTaskNotFound {
			logger.Warnf("Task ID %d not found for deletion", id)
			response.WriteError(w, http.StatusNotFound, "task not found")
			return
		}
	}

	if err := taskRepo.Delete(id); err != nil {
		if err == repository.ErrTaskNotFound {
			logger.Warnf("Task ID %d not found for deletion", id)
			response.WriteError(w, http.StatusNotFound, "task not found")
			return
		}
		logger.Errorf("Error deleting task ID %d: %v", id, err)
		response.WriteError(w, http.StatusInternalServerError, "failed to delete task")
		return
	}

	logger.Infof("Task with ID %d deleted successfully", id)
	response.WriteSuccess(w, http.StatusOK, fmt.Sprintf("Task with Id %d deleted successfully", id))
}
