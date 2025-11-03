package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/AbhayBharti21/task-manager/internal/http/types"
	"github.com/AbhayBharti21/task-manager/internal/http/utils/response"
	"io"
	"net/http"
	"sync"
)

var (
	Tasks map[int]types.Task
	Mu    *sync.Mutex
)

func CreateTask(w http.ResponseWriter, r *http.Request) {
	var task types.Task

	err := json.NewDecoder(r.Body).Decode(&task)
	if errors.Is(err, io.EOF) {
		response.WriteJson(w, http.StatusBadRequest, fmt.Errorf("empty body"))
		return
	}

	if err != nil {
		response.WriteJson(w, http.StatusBadRequest, err)
	}

	fmt.Println(task)

	response.WriteJson(w, http.StatusCreated, map[string]string{"success": "OK"})
}
