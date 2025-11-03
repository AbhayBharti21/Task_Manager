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
		response.WriteJson(w, http.StatusBadRequest, map[string]string{"Error": "empty body"})
		return
	}

	if err != nil {
		response.WriteJson(w, http.StatusBadRequest, err)
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
	fmt.Println(tasks)

	response.WriteJson(w, http.StatusCreated, map[string]string{"success": "OK", "message": msg})

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
	}

	id, err := strconv.Atoi(pathId[3])
	if err != nil {
		logger2.Logger.Println("Error: Unable to Convert path id to int !!")

	}

	taskData, ok := tasks[id]
	if !ok {
		response.WriteJson(w, http.StatusNotFound, map[string]string{"Error": "Task Id not Found!!"})
		return
	}

	response.WriteJson(w, http.StatusOK, map[string]any{"success": "OK", "task": taskData})
}
