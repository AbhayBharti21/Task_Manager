package main

import (
	"fmt"
	"github.com/AbhayBharti21/task-manager/internal/config"
	"github.com/AbhayBharti21/task-manager/internal/http/handlers"
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

	router.HandleFunc("POST /api/tasks", handlers.CreateTask)
	//router.HandleFunc("GET /api/tasks/:id", handlers.GetTask)

	server := http.Server{
		Addr:    cfg.Addr,
		Handler: router,
	}

	slog.Info("Server started at", slog.String("PORT", cfg.HTTPServer.Addr))
	logger2.Logger.Println()

	serverErr := server.ListenAndServe()
	if serverErr != nil {
		logger2.Logger.Fatalf("Failed to start server %v", serverErr)
	}

}
