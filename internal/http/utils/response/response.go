package response

import (
	"encoding/json"
	"net/http"

	"github.com/AbhayBharti21/task-manager/internal/http/utils/logger"
)

func WriteJSON(w http.ResponseWriter, status int, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		logger.Errorf("Error encoding JSON response: %v", err)
		return err
	}

	return nil
}

func WriteError(w http.ResponseWriter, status int, message string) {
	WriteJSON(w, status, map[string]string{"error": message})
}

func WriteSuccess(w http.ResponseWriter, status int, data interface{}) {
	response := map[string]interface{}{
		"success": true,
	}

	if message, ok := data.(string); ok {
		response["message"] = message
	} else {
		response["data"] = data
	}

	WriteJSON(w, status, response)
}

func WriteSuccessWithData(w http.ResponseWriter, status int, data interface{}) {
	WriteJSON(w, status, map[string]interface{}{
		"success": true,
		"data":    data,
	})
}
