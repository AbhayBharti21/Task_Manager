package logger

import (
	"log"
	"os"
	"path/filepath"
	"time"
)

var Logger *log.Logger

// Init initializes the logger and creates a log file
func Init() error {
	timestamp := time.Now().Format("2006-01-02_15-04-05")
	logFilePath := filepath.Join("logs", "log_file_"+timestamp+".txt")

	f, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return err
	}

	Logger = log.New(f, "", log.LstdFlags|log.Lshortfile)
	Println("Logger initialized")
	return nil
}

// Error logs an error message with formatting
func Error(format string, v ...interface{}) {
	if Logger != nil {
		Logger.Printf("ERROR: "+format, v...)
	}
}

// Errorf logs an error message (alias for consistency)
func Errorf(format string, v ...interface{}) {
	Error(format, v...)
}

// Info logs an informational message with formatting
func Info(format string, v ...interface{}) {
	if Logger != nil {
		Logger.Printf("INFO: "+format, v...)
	}
}

// Infof logs an informational message (alias for consistency)
func Infof(format string, v ...interface{}) {
	Info(format, v...)
}

// Debug logs a debug message with formatting
func Debug(format string, v ...interface{}) {
	if Logger != nil {
		Logger.Printf("DEBUG: "+format, v...)
	}
}

// Debugf logs a debug message (alias for consistency)
func Debugf(format string, v ...interface{}) {
	Debug(format, v...)
}

// Warn logs a warning message with formatting
func Warn(format string, v ...interface{}) {
	if Logger != nil {
		Logger.Printf("WARN: "+format, v...)
	}
}

// Warnf logs a warning message (alias for consistency)
func Warnf(format string, v ...interface{}) {
	Warn(format, v...)
}

// Fatal logs a fatal message and exits
func Fatal(v ...interface{}) {
	if Logger != nil {
		Logger.Fatal(v...)
	}
}

// Fatalf logs a fatal message with formatting and exits
func Fatalf(format string, v ...interface{}) {
	if Logger != nil {
		Logger.Fatalf(format, v...)
	}
}

// Print logs a message without prefix
func Print(v ...interface{}) {
	if Logger != nil {
		Logger.Print(v...)
	}
}

// Println logs a message without prefix and adds a newline
func Println(v ...interface{}) {
	if Logger != nil {
		Logger.Println(v...)
	}
}
