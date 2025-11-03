package main

import (
	"fmt"
	"github.com/AbhayBharti21/task-manager/internal/config"
	"github.com/AbhayBharti21/task-manager/internal/http/handlers"
	"github.com/AbhayBharti21/task-manager/internal/http/middleware"
	logger2 "github.com/AbhayBharti21/task-manager/internal/http/utils/logger"
	"log/slog"
	"net/http"
)

func main() {
	cfg := config.MustLoad()
	router := http.NewServeMux()

	err := logger2.Init()
	if err != nil {
		fmt.Println("unable to open logger")
		return
	}

	router.HandleFunc("POST /api/tasks", middleware.LogRequest(handlers.CreateTask))
	router.HandleFunc("GET /api/tasks/", middleware.LogRequest(handlers.GetTask))
	router.HandleFunc("PATCH /api/tasks/", middleware.LogRequest(handlers.UpdateTask))
	router.HandleFunc("DELETE /api/tasks/", middleware.LogRequest(handlers.DeleteTask))
	router.HandleFunc("GET /api/health", middleware.LogRequest(handlers.HealthStatus))

	server := http.Server{
		Addr:    cfg.Addr,
		Handler: router,
	}

	slog.Info("Server started at", slog.String("PORT", cfg.HTTPServer.Addr))

	serverErr := server.ListenAndServe()
	if serverErr != nil {
		logger2.Logger.Fatalf("Failed to start server %v", serverErr)
	}

}
