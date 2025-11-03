package handlers

import (
	"net/http"

	"github.com/AbhayBharti21/task-manager/internal/http/utils/logger"
	"github.com/AbhayBharti21/task-manager/internal/http/utils/response"
)

func HealthStatus(w http.ResponseWriter, r *http.Request) {
	logger.Info("Health check: Service is up")
	response.WriteJSON(w, http.StatusOK, map[string]interface{}{
		"success": true,
		"message": "👌Sb Chal rha hai",
	})
}
