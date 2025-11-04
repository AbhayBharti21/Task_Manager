package logger

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

var Logger *log.Logger

func Init() error {
	logsDir := "logs"
	if err := os.MkdirAll(logsDir, 0755); err != nil {
		return fmt.Errorf("failed to create logs directory: %w", err)
	}

	timestamp := time.Now().Format("2006-01-02_15-04-05")
	logFilePath := filepath.Join(logsDir, "log_file_"+timestamp+".txt")

	f, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return fmt.Errorf("failed to open log file: %w", err)
	}

	Logger = log.New(f, "", log.LstdFlags|log.Lshortfile)
	Println("Logger initialized")
	return nil
}

func Error(format string, v ...interface{}) {
	if Logger != nil {
		Logger.Printf("ERROR: "+format, v...)
	}
}

func Errorf(format string, v ...interface{}) {
	Error(format, v...)
}

func Info(format string, v ...interface{}) {
	if Logger != nil {
		Logger.Printf("INFO: "+format, v...)
	}
}

func Infof(format string, v ...interface{}) {
	Info(format, v...)
}

func Debug(format string, v ...interface{}) {
	if Logger != nil {
		Logger.Printf("DEBUG: "+format, v...)
	}
}

func Debugf(format string, v ...interface{}) {
	Debug(format, v...)
}

func Warn(format string, v ...interface{}) {
	if Logger != nil {
		Logger.Printf("WARN: "+format, v...)
	}
}

func Warnf(format string, v ...interface{}) {
	Warn(format, v...)
}

func Fatal(v ...interface{}) {
	if Logger != nil {
		Logger.Fatal(v...)
	}
}

func Fatalf(format string, v ...interface{}) {
	if Logger != nil {
		Logger.Fatalf(format, v...)
	}
}

func Print(v ...interface{}) {
	if Logger != nil {
		Logger.Print(v...)
	}
}

func Println(v ...interface{}) {
	if Logger != nil {
		Logger.Println(v...)
	}
}
