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

type TaskRepository struct {
	tasks    map[int]types.Task
	ownerInc int
	taskInc  int
	mu       sync.RWMutex
}

func NewTaskRepository() *TaskRepository {
	return &TaskRepository{
		tasks:    make(map[int]types.Task),
		ownerInc: 1,
		taskInc:  1,
	}
}

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

func (tr *TaskRepository) GetByID(id int) (types.Task, error) {
	tr.mu.RLock()
	defer tr.mu.RUnlock()

	task, exists := tr.tasks[id]
	if !exists {
		return types.Task{}, ErrTaskNotFound
	}

	return task, nil
}

func (tr *TaskRepository) Update(id int, updates types.Task) (types.Task, error) {
	tr.mu.Lock()
	defer tr.mu.Unlock()

	existing, exists := tr.tasks[id]
	if !exists {
		return types.Task{}, ErrTaskNotFound
	}

	if updates.Description != "" {
		existing.Description = updates.Description
	}

	if updates.IsCompleted {
		existing.IsCompleted = updates.IsCompleted
	}

	tr.tasks[id] = existing
	return existing, nil
}

func (tr *TaskRepository) Delete(id int) error {
	tr.mu.Lock()
	defer tr.mu.Unlock()

	if _, exists := tr.tasks[id]; !exists {
		return ErrTaskNotFound
	}

	delete(tr.tasks, id)
	return nil
}

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
