package middleware

import (
	"net/http"

	"github.com/AbhayBharti21/task-manager/internal/http/utils/logger"
)

func LogRequest(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Debugf("Request: %s %s %s", r.RemoteAddr, r.Method, r.URL)
		next(w, r)
	}
}
