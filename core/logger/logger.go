package logger

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"time"
)

func Init() (*os.File, error) {
	if err := os.MkdirAll("logs", 0755); err != nil {
		return nil, fmt.Errorf("не удалось создать папку логов: %w", err)
	}

	filename := time.Now().Format("2006-01-02_15-04-05.000000") + ".log"
	path := filepath.Join("logs", filename)

	file, err := os.OpenFile(
		path,
		os.O_CREATE|os.O_WRONLY|os.O_APPEND,
		0644,
	)
	if err != nil {
		return nil, fmt.Errorf("не удалось создать файл логов: %w", err)
	}

	slog.SetDefault(slog.New(slog.NewJSONHandler(file, nil)))
	return file, nil
}
