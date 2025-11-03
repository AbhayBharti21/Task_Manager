package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/AbhayBharti21/task-manager/internal/http/models"
	logger2 "github.com/AbhayBharti21/task-manager/internal/http/utils/logger"
	"github.com/AbhayBharti21/task-manager/internal/http/utils/response"
	"io"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

var (
	tasks        = make(map[int]types.Task)
	ownerInc int = 1
	taskInc  int = 1
	mu       sync.Mutex
)

func CreateTask(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()
	var task types.Task

	err := json.NewDecoder(r.Body).Decode(&task)
	if errors.Is(err, io.EOF) {
		logger2.Logger.Println("Error: Empty Body")
		response.WriteJson(w, http.StatusBadRequest, map[string]string{"Error": "empty body"})
		return
	}

	if err != nil {
		logger2.Logger.Printf("Error: %v", err)
		response.WriteJson(w, http.StatusBadRequest, err)
	}

	if task.Description == "" {
		logger2.Logger.Printf("Error: Description Not Found")
		response.WriteJson(w, http.StatusBadRequest, map[string]any{"Error": "Description is required"})
	}

	isOwnerId := task.OwnerId != 0

	if isOwnerId {
		tasks[taskInc] = types.Task{
			TaskId:      taskInc,
			OwnerId:     task.OwnerId,
			Description: task.Description,
			IsCompleted: false,
		}
	} else {
		tasks[taskInc] = types.Task{
			TaskId:      taskInc,
			OwnerId:     ownerInc,
			Description: task.Description,
			IsCompleted: false,
		}
	}

	msg := fmt.Sprintf("Task created successfully with task Id %d", taskInc)
	logger2.Logger.Println(msg)

	response.WriteJson(w, http.StatusCreated, map[string]any{"success": true, "message": msg})

	if !isOwnerId {
		ownerInc++
	}
	taskInc++
}

func GetTask(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	path := r.URL.Path
	pathId := strings.Split(path, "/")

	if len(pathId) < 4 {
		logger2.Logger.Println("Error: Path params not found!!")
		response.WriteJson(w, http.StatusBadRequest, map[string]string{"Error": "Path params not found!!"})
		return
	}

	id, err := strconv.Atoi(pathId[3])
	if err != nil {
		logger2.Logger.Println("Error: Unable to Convert path id to int !!")
		response.WriteJson(w, http.StatusBadRequest, map[string]string{"Error": "Path params not found!!"})
		return
	}

	taskData, ok := tasks[id]
	if !ok {
		logger2.Logger.Println("Error: Task Id not Found!!")
		response.WriteJson(w, http.StatusNotFound, map[string]string{"Error": "Task Id not Found!!"})
		return
	}

	logger2.Logger.Printf("Task with id %d return successfully", id)
	response.WriteJson(w, http.StatusOK, map[string]any{"success": true, "task": taskData})
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	var updateFields types.Task

	path := r.URL.Path
	pathId := strings.Split(path, "/")

	if len(pathId) < 4 {
		logger2.Logger.Println("Error: Path params not found!!")
		response.WriteJson(w, http.StatusBadRequest, map[string]string{"Error": "Path params not found!!"})
		return
	}

	id, err := strconv.Atoi(pathId[3])
	if err != nil {
		logger2.Logger.Println("Error: Unable to Convert path id to int !!")
		response.WriteJson(w, http.StatusBadRequest, map[string]any{"Error": "Params can't be empty"})
		return
	}

	taskData, ok := tasks[id]
	if !ok {
		logger2.Logger.Println("Error: Task Id not found")
		response.WriteJson(w, http.StatusNotFound, map[string]string{"Error": "Task Id not Found!!"})
		return
	} else {
		if tasks[id].OwnerId != updateFields.OwnerId {
			logger2.Logger.Println("Error: Unauthorized")
			response.WriteJson(w, http.StatusUnauthorized, map[string]string{"Error": "Unauthorized"})
			return
		}
	}

	bodyErr := json.NewDecoder(r.Body).Decode(&updateFields)
	if bodyErr != nil {
		logger2.Logger.Println("Update fields not found")
		response.WriteJson(w, http.StatusBadRequest, map[string]string{"Error": "😭 Update Fields Not Found"})
		return
	}

	var description string
	var isComplete bool

	if updateFields.Description == "" {
		description = taskData.Description
	} else {
		description = updateFields.Description
	}

	if updateFields.IsCompleted == false {
		isComplete = taskData.IsCompleted
	} else {
		isComplete = updateFields.IsCompleted
	}

	tasks[id] = types.Task{
		TaskId:      taskData.TaskId,
		OwnerId:     taskData.OwnerId,
		Description: description,
		IsCompleted: isComplete,
	}

	logger2.Logger.Printf("Task With id %d Successfully Updated\n", taskData.TaskId)
	response.WriteJson(w, http.StatusOK, map[string]any{"success": true, "task": tasks[id]})
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	var deleteFields types.Task

	err := json.NewDecoder(r.Body).Decode(&deleteFields)

	if err != nil {
		logger2.Logger.Println("Error: Empty data")
		response.WriteJson(w, http.StatusBadRequest, map[string]any{"Error": "Empty Data"})
		return
	}

	path := r.URL.Path
	pathId := strings.Split(path, "/")

	if len(pathId) < 4 {
		logger2.Logger.Println("Error: Path params not found!!")
		response.WriteJson(w, http.StatusBadRequest, map[string]string{"Error": "Path params not found!!"})
		return
	}

	id, err := strconv.Atoi(pathId[3])
	if err != nil {
		logger2.Logger.Println("Error: Unable to Convert path id to int !!")
		response.WriteJson(w, http.StatusBadRequest, map[string]string{"Error": "Path params not found!!"})
		return
	}

	_, ok := tasks[id]
	if !ok {
		logger2.Logger.Println("Error: Task Id not Found!!")
		response.WriteJson(w, http.StatusNotFound, map[string]string{"Error": "Task Id not Found!!"})
		return
	} else {
		if tasks[id].OwnerId != deleteFields.OwnerId {
			logger2.Logger.Println("Error: Unauthorized")
			response.WriteJson(w, http.StatusUnauthorized, map[string]string{"Error": "Unauthorized"})
			return
		}
	}

	delete(tasks, id)

	logger2.Logger.Printf("Task with id %d deleted successfully\n", id)
	deleteMsg := fmt.Sprintf("Task with Id %d deleted successfully", id)
	response.WriteJson(w, http.StatusOK, map[string]any{"success": true, "message": deleteMsg})
}
