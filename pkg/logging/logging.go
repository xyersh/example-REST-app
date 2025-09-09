package logging

import (
	"log/slog"
	"os"
	"path/filepath"
)

var (
	logger       *slog.Logger
	jsonFileInfo *os.File
	jsonFileWarn *os.File
)

func init() {
	// Создаем JSON-обработчик для вывода в файл
	var err error

	//формируем пути к файлам логов
	absPath, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	absPath = filepath.Join(absPath, "logs")

	infoLogPath := filepath.Join(absPath, "info_logs.json")
	warnLogPath := filepath.Join(absPath, "warning_logs.json")

	err = os.MkdirAll(absPath, 0750)

	// формируем файлы логов
	jsonFileInfo, err = os.OpenFile(infoLogPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644) //os.Create("info_logs.json")
	if err != nil {
		slog.Error("failed to create log file", "error", err)
		os.Exit(1)
	}

	jsonFileWarn, err = os.OpenFile(warnLogPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644) //os.Create("warning_logs.json")
	if err != nil {
		slog.Error("failed to create log file", "error", err)
		os.Exit(1)
	}

	// создаем JSON-обработчики
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
