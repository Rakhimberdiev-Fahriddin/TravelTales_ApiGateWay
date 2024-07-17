package logs

import (
	"log"
	"log/slog"
	"os"
)

var Logger *slog.Logger

func InitLogger() {
	logFile, err := os.OpenFile("logs/app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}

	handler := slog.NewJSONHandler(logFile, nil)
	Logger = slog.New(handler)
}
