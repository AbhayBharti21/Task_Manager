package main

import (
	"fmt"
	"net/http"

	"github.com/AbhayBharti21/task-manager/internal/http/handlers"
	"github.com/AbhayBharti21/task-manager/internal/http/middleware"
	"github.com/AbhayBharti21/task-manager/internal/http/utils/logger"
)

const defaultPort = ":8080"

func main() {
	if err := logger.Init(); err != nil {
		fmt.Printf("Failed to initialize logger: %v\n", err)
		return
	}

	router := setupRoutes()

	logger.Infof("Server starting on port %s", defaultPort)

	if err := http.ListenAndServe(defaultPort, router); err != nil {
		logger.Fatalf("Failed to start server: %v", err)
	}
}

func setupRoutes() *http.ServeMux {
	router := http.NewServeMux()

	// Register all routes with logging middleware
	router.HandleFunc("POST /api/tasks", middleware.LogRequest(handlers.CreateTask))
	router.HandleFunc("GET /api/tasks/", middleware.LogRequest(handlers.GetTask))
	router.HandleFunc("PATCH /api/tasks/", middleware.LogRequest(handlers.UpdateTask))
	router.HandleFunc("DELETE /api/tasks/", middleware.LogRequest(handlers.DeleteTask))
	router.HandleFunc("GET /api/health", middleware.LogRequest(handlers.HealthStatus))

	return router
}
