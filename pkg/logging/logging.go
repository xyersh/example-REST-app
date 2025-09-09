package logging

import (
	"log/slog"
	"os"
)

var (
	logger       *slog.Logger
	jsonFileInfo *os.File
	jsonFileWarn *os.File
)

func init() {
	// Создаем JSON-обработчик для вывода в файл
	var err error
	jsonFileInfo, err = os.OpenFile("info_logs.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644) //os.Create("info_logs.json")
	if err != nil {
		slog.Error("failed to create log file", "error", err)
		os.Exit(1)
	}

	jsonFileWarn, err = os.OpenFile("warning_logs.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644) //os.Create("warning_logs.json")
	if err != nil {
		slog.Error("failed to create log file", "error", err)
		os.Exit(1)
	}

	jsonHandlerInfo := slog.NewJSONHandler(jsonFileInfo, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	})

	jsonHandlerWarn := slog.NewJSONHandler(jsonFileWarn, &slog.HandlerOptions{
		Level: slog.LevelWarn,
	})

	// Создаем текстовый обработчик для вывода в консоль (stdout)
	textHandler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})

	// Объединяем обработчики в MultiHandler
	multiHandler := NewMultiHandler(jsonHandlerInfo, jsonHandlerWarn, textHandler)

	// Создаем логгер, который использует MultiHandler
	logger = slog.New(multiHandler)

	// Устанавливаем наш логгер как глобальный, чтобы его могли использовать все функции
	slog.SetDefault(logger)

	slog.Info("Log initialization complete")
}
