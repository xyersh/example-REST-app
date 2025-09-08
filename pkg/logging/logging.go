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

// type writerHook struct {
// 	Writer    []io.Writer
// 	LogLevels []logrus.Level
// }

// func (hook *writerHook) Fire(entry *logrus.Entry) error {
// 	line, err := entry.String()
// 	if err != nil {
// 		return err
// 	}

// 	for _, w := range hook.Writer {
// 		w.Write([]byte(line))
// 	}

// 	return nil
// }

// func (hook *writerHook) Levels() []logrus.Level {
// 	return hook.LogLevels
// }

// var e *logrus.Entry

// type Logger struct {
// 	*logrus.Entry
// }

// func GetLogger() {
// 	return Logger{e}
// }

// func init() {
// 	l := logrus.New()
// 	l.SetReportCaller(true)
// 	l.Formatter = &logrus.TextFormatter{
// 		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
// 			filename := path.Base(frame.File)
// 			return fmt.Sprintf("%s()", frame.Function), fmt.Sprintf("%s %d", filename, frame.Line)
// 		},
// 		DisableColors: false,
// 		FullTimestamp: true,
// 	}

// 	err := os.Mkdir("logs", 0644)
// 	if err != nil {
// 		panic(err)
// 	}

// 	allFile, err := os.OpenFile("logs/all.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
// 	if err != nil {
// 		panic(err)
// 	}

// 	l.SetOutput(io.Discard)
// 	l.AddHook()(&writerHook{
// 		Writer:    []io.Writer{os.Stdout, allFile},
// 		LogLevels: logrus.AllLevels,
// 	})

// 	l.SetLevel(logrus.TraceLevel)
// }

// // kafka 	-- info, debug
// // file  	-- error, trace
// // stdout 	-- warning, critical
