package logger

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"time"
)

func Init() (*os.File, error) {
	const logDir = "logs"

	if err := os.MkdirAll(logDir, 0775); err != nil {
		return nil, fmt.Errorf("не удалось создать папку логов %w", err)
	}

	fileName := time.Now().Format("2006-01-02_15-04-05") + ".log"

	path := filepath.Join(logDir, fileName)

	file, err := os.OpenFile(
		path,
		os.O_CREATE|os.O_WRONLY|os.O_APPEND|os.O_EXCL,
		0644,
	)
	if err != nil {
		return nil, fmt.Errorf("Не удалось создать файл логов: %w", err)
	}

	handler := slog.NewJSONHandler(file, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	})

	slog.SetDefault(slog.New(handler))

	return file, nil
}
