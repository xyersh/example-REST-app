package logging

import (
	"context"
	"log/slog"
)

// MultiHandler - кастомный обработчик, который отправляет логи в несколько получателей.
type MultiHandler struct {
	handlers []slog.Handler
}

// NewMultiHandler создает новый MultiHandler
func NewMultiHandler(handlers ...slog.Handler) *MultiHandler {
	return &MultiHandler{
		handlers: handlers,
	}
}

// Enabled реализует метод интерфейса slog.Handler.
// В этом примере, если хотя бы один обработчик включен, MultiHandler тоже будет включен.
func (h *MultiHandler) Enabled(ctx context.Context, level slog.Level) bool {
	for _, handler := range h.handlers {
		if handler.Enabled(ctx, level) {
			return true
		}
	}
	return false
}

// Handle реализует метод интерфейса slog.Handler.
// Отправляет запись лога каждому из вложенных обработчиков.
func (h *MultiHandler) Handle(ctx context.Context, record slog.Record) error {
	for _, handler := range h.handlers {
		if handler.Enabled(ctx, record.Level) {
			if err := handler.Handle(ctx, record); err != nil {
				return err
			}
		}
	}
	return nil
}

// WithAttrs реализует метод интерфейса slog.Handler.
func (h *MultiHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	newHandlers := make([]slog.Handler, len(h.handlers))
	for i, handler := range h.handlers {
		newHandlers[i] = handler.WithAttrs(attrs)
	}
	return NewMultiHandler(newHandlers...)
}

// WithGroup реализует метод интерфейса slog.Handler.
func (h *MultiHandler) WithGroup(name string) slog.Handler {
	newHandlers := make([]slog.Handler, len(h.handlers))
	for i, handler := range h.handlers {
		newHandlers[i] = handler.WithGroup(name)
	}
	return NewMultiHandler(newHandlers...)
}
