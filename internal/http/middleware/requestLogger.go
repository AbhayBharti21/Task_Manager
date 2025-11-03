package middleware

import (
	logger2 "github.com/AbhayBharti21/task-manager/internal/http/utils/logger"
	"net/http"
)

func LogRequest(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger2.Logger.Printf("%s %s %s \n", r.RemoteAddr, r.Method, r.URL)
		next(w, r)
	}
}
