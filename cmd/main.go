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
	fmt.Printf("🚀 Server starting on port %s\n", defaultPort)

	if err := http.ListenAndServe(defaultPort, router); err != nil {
		fmt.Printf("❌ Failed to start server: %v\n", err)
		logger.Fatalf("Failed to start server: %v", err)
	}
}

func setupRoutes() *http.ServeMux {
	router := http.NewServeMux()

	// Register all routes with logging middleware - specific routes first
	
	router.HandleFunc("POST /api/tasks/create", middleware.LogRequest(handlers.CreateTask))
	router.HandleFunc("GET /api/tasks/get/", middleware.LogRequest(handlers.GetTask))
	router.HandleFunc("PATCH /api/tasks/update/", middleware.LogRequest(handlers.UpdateTask))
	router.HandleFunc("DELETE /api/tasks/delete/", middleware.LogRequest(handlers.DeleteTask))
	router.HandleFunc("GET /api/health", middleware.LogRequest(handlers.HealthStatus))

	// Explicit handlers for /api/tasks and /api/tasks/ to prevent wrong matching
	router.HandleFunc("GET /api/tasks", func(w http.ResponseWriter, r *http.Request) {
		logger.Errorf("Method not allowed: %s %s", r.Method, r.URL.Path)
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	})
	router.HandleFunc("GET /api/tasks/", func(w http.ResponseWriter, r *http.Request) {
		logger.Errorf("Method not allowed: %s %s", r.Method, r.URL.Path)
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	})

	// Catch-all handler for unmatched routes - must be last
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		logger.Errorf("Unmatched route: %s %s", r.Method, r.URL.Path)
		http.NotFound(w, r)
	})

	return router
}
