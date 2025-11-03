package repository

import (
	"errors"
	"sync"

	types "github.com/AbhayBharti21/task-manager/internal/http/models"
)

var (
	ErrTaskNotFound = errors.New("task not found")
	ErrUnauthorized = errors.New("unauthorized: task owner mismatch")
)

// TaskRepository handles task storage operations
type TaskRepository struct {
	tasks    map[int]types.Task
	ownerInc int
	taskInc  int
	mu       sync.RWMutex
}

// NewTaskRepository creates a new task repository instance
func NewTaskRepository() *TaskRepository {
	return &TaskRepository{
		tasks:    make(map[int]types.Task),
		ownerInc: 1,
		taskInc:  1,
	}
}

// Create creates a new task and returns it
func (tr *TaskRepository) Create(task types.Task) types.Task {
	tr.mu.Lock()
	defer tr.mu.Unlock()

	var ownerID int
	if task.OwnerId != 0 {
		ownerID = task.OwnerId
	} else {
		ownerID = tr.ownerInc
		tr.ownerInc++
	}

	newTask := types.Task{
		TaskId:      tr.taskInc,
		OwnerId:     ownerID,
		Description: task.Description,
		IsCompleted: false,
	}

	tr.tasks[tr.taskInc] = newTask
	tr.taskInc++

	return newTask
}

// GetByID retrieves a task by its ID
func (tr *TaskRepository) GetByID(id int) (types.Task, error) {
	tr.mu.RLock()
	defer tr.mu.RUnlock()

	task, exists := tr.tasks[id]
	if !exists {
		return types.Task{}, ErrTaskNotFound
	}

	return task, nil
}

// Update updates an existing task
func (tr *TaskRepository) Update(id int, updates types.Task) (types.Task, error) {
	tr.mu.Lock()
	defer tr.mu.Unlock()

	existing, exists := tr.tasks[id]
	if !exists {
		return types.Task{}, ErrTaskNotFound
	}

	// Update description if provided
	if updates.Description != "" {
		existing.Description = updates.Description
	}

	// Update completion status if provided (check if it's explicitly set to true)
	// Note: This logic assumes false means "not provided" in updates
	// You might want to use pointers or a different approach for optional fields
	if updates.IsCompleted {
		existing.IsCompleted = updates.IsCompleted
	}

	tr.tasks[id] = existing
	return existing, nil
}

// Delete removes a task by its ID
func (tr *TaskRepository) Delete(id int) error {
	tr.mu.Lock()
	defer tr.mu.Unlock()

	if _, exists := tr.tasks[id]; !exists {
		return ErrTaskNotFound
	}

	delete(tr.tasks, id)
	return nil
}

// VerifyOwner checks if the given owner ID matches the task's owner
func (tr *TaskRepository) VerifyOwner(taskID int, ownerID int) error {
	tr.mu.RLock()
	defer tr.mu.RUnlock()

	task, exists := tr.tasks[taskID]
	if !exists {
		return ErrTaskNotFound
	}

	if task.OwnerId != ownerID {
		return ErrUnauthorized
	}

	return nil
}
