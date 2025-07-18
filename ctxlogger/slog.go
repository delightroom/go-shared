package ctxlogger

import (
	"fmt"
	"log/slog"
	"os"
	"strings"
)

// Slog 는 slog 로거를 ctxlogger.Logger 인터페이스로 래핑한 struct입니다.
// ctxlogger에 담을 로거로 slog 패키지 사용시 이 struct를 사용하세요.
type Slog struct {
	logger *slog.Logger
}

// NewSlog 는 환경 변수에서 로그 레벨을 읽어와서 slog 로거를 초기화합니다.
func NewSlog() *Slog {
	logLevel := os.Getenv("LOG_LEVEL")
	var level slog.Level
	switch strings.ToLower(logLevel) {
	case "debug":
		level = slog.LevelDebug
	case "info":
		level = slog.LevelInfo
	case "warn", "warning":
		level = slog.LevelWarn
	case "error":
		level = slog.LevelError
	default:
		level = slog.LevelInfo
	}

	opts := &slog.HandlerOptions{
		Level: level,
	}

	handler := slog.NewJSONHandler(os.Stdout, opts)
	logger := slog.New(handler)

	return &Slog{
		logger: logger,
	}
}

// Debug logs a message with Debug level
func (l *Slog) Debug(args ...any) {
	l.logger.Debug(fmt.Sprint(args...))
}

// Debugf logs a formatted message with Debug level
func (l *Slog) Debugf(template string, args ...any) {
	l.logger.Debug(fmt.Sprintf(template, args...))
}

// Debugw logs a message with key-value pairs with Debug level
func (l *Slog) Debugw(msg string, keysAndValues ...any) {
	l.logger.Debug(msg, keysAndValues...)
}

// Info logs a message with Info level
func (l *Slog) Info(args ...any) {
	l.logger.Info(fmt.Sprint(args...))
}

// Infof logs a formatted message with Info level
func (l *Slog) Infof(template string, args ...any) {
	l.logger.Info(fmt.Sprintf(template, args...))
}

// Infow logs a message with key-value pairs with Info level
func (l *Slog) Infow(msg string, keysAndValues ...any) {
	l.logger.Info(msg, keysAndValues...)
}

// Warn logs a message with Warn level
func (l *Slog) Warn(args ...any) {
	l.logger.Warn(fmt.Sprint(args...))
}

// Warnf logs a formatted message with Warn level
func (l *Slog) Warnf(template string, args ...any) {
	l.logger.Warn(fmt.Sprintf(template, args...))
}

// Warnw logs a message with key-value pairs with Warn level
func (l *Slog) Warnw(msg string, keysAndValues ...any) {
	l.logger.Warn(msg, keysAndValues...)
}

// Error logs a message with Error level
func (l *Slog) Error(args ...any) {
	l.logger.Error(fmt.Sprint(args...))
}

// Errorf logs a formatted message with Error level
func (l *Slog) Errorf(template string, args ...any) {
	l.logger.Error(fmt.Sprintf(template, args...))
}

// Errorw logs a message with key-value pairs with Error level
func (l *Slog) Errorw(msg string, keysAndValues ...any) {
	l.logger.Error(msg, keysAndValues...)
}

// Fatal logs a message with Fatal level and then calls os.Exit(1)
func (l *Slog) Fatal(args ...any) {
	l.logger.Error(fmt.Sprint(args...))
	os.Exit(1)
}

// Fatalf logs a formatted message with Fatal level and then calls os.Exit(1)
func (l *Slog) Fatalf(template string, args ...any) {
	l.logger.Error(fmt.Sprintf(template, args...))
	os.Exit(1)
}

// Fatalw logs a message with key-value pairs with Fatal level and then calls os.Exit(1)
func (l *Slog) Fatalw(msg string, keysAndValues ...any) {
	l.logger.Error(msg, keysAndValues...)
	os.Exit(1)
}

// With creates a child logger with additional fields
func (l *Slog) With(args ...any) Logger {
	return &Slog{
		logger: l.logger.With(args...),
	}
}

// SugaredLogger returns nil as slog doesn't have a direct equivalent
func (l *Slog) SugaredLogger() *slog.Logger {
	return l.logger
}
