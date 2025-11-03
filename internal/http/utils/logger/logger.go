package logger

import (
	"log"
	"os"
	"path/filepath"
	"time"
)

var Logger *log.Logger

func Init() error {
	timestamp := time.Now().Format("2006-01-02_15-04-05")
	logFilePath := filepath.Join("logs", "log_file_"+timestamp+".txt")

	f, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return err
	}

	Logger = log.New(f, "", log.LstdFlags|log.Lshortfile)
	Logger.Println("Logger initialized")
	return nil
}
