package handlers

import (
	logger2 "github.com/AbhayBharti21/task-manager/internal/http/utils/logger"
	"github.com/AbhayBharti21/task-manager/internal/http/utils/response"
	"net/http"
)

func HealthStatus(w http.ResponseWriter, r *http.Request) {
	logger2.Logger.Println("Service is Up")
	response.WriteJson(w, http.StatusOK, map[string]any{"success": true, "message": "👌Sb Chal rha hai"})
}
