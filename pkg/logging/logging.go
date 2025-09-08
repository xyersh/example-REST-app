package logging

import (
	"log/slog"
	"os"
)

var logger *slog.Logger

func init() {
	slog.Info("Log initialization starts")
	// Создаем JSON-обработчик для вывода в файл

	jsonFile, err := os.Create("logs.json")
	if err != nil {
		slog.Error("failed to create log file", "error", err)
		os.Exit(1)
	}
	defer jsonFile.Close()

	jsonHandler := slog.NewJSONHandler(jsonFile, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	})

	// Создаем текстовый обработчик для вывода в консоль (stdout)
	textHandler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})

	// Объединяем обработчики в MultiHandler
	multiHandler := NewMultiHandler(jsonHandler, textHandler)

	// Создаем логгер, который использует MultiHandler
	logger = slog.New(multiHandler)

	// Устанавливаем наш логгер как глобальный, чтобы его могли использовать все функции
	slog.SetDefault(logger)

	slog.Info("Log initialization ends")
}

