package ctxlogger

import (
	"context"
)

// Logger 는 모든 로거 구현이 충족해야 하는 일반 로거 인터페이스입니다.
type Logger interface {
	Debug(args ...any)
	Debugf(template string, args ...any)
	Debugw(msg string, keysAndValues ...any)

	Info(args ...any)
	Infof(template string, args ...any)
	Infow(msg string, keysAndValues ...any)

	Warn(args ...any)
	Warnf(template string, args ...any)
	Warnw(msg string, keysAndValues ...any)

	Error(args ...any)
	Errorf(template string, args ...any)
	Errorw(msg string, keysAndValues ...any)

	Fatal(args ...any)
	Fatalf(template string, args ...any)
	Fatalw(msg string, keysAndValues ...any)

	With(args ...any) Logger
}

type ctxLoggerKey struct{}

// LoggerFromContext 는 컨텍스트에서 로거를 추출합니다.
// 컨텍스트에 로거가 없으면 제공된 폴백 로거를 반환합니다.
// 폴백이 nil이면 널 로거를 반환합니다.
func LoggerFromContext(ctx context.Context, fallback Logger) Logger {
	if l, ok := ctx.Value(ctxLoggerKey{}).(Logger); ok {
		return l
	}
	if fallback != nil {
		return fallback
	}
	return &noOpLogger{}
}

// ContextWithLogger 는 로거를 컨텍스트에 추가합니다.
// 컨텍스트에 로거가 이미 있으면 덮어씁니다.
func ContextWithLogger(ctx context.Context, l Logger) context.Context {
	return context.WithValue(ctx, ctxLoggerKey{}, l)
}

// noOpLogger 는 아무것도 하지 않는 로거입니다.
type noOpLogger struct{}

func (n *noOpLogger) Debug(args ...any)                       {}
func (n *noOpLogger) Debugf(template string, args ...any)     {}
func (n *noOpLogger) Debugw(msg string, keysAndValues ...any) {}
func (n *noOpLogger) Info(args ...any)                        {}
func (n *noOpLogger) Infof(template string, args ...any)      {}
func (n *noOpLogger) Infow(msg string, keysAndValues ...any)  {}
func (n *noOpLogger) Warn(args ...any)                        {}
func (n *noOpLogger) Warnf(template string, args ...any)      {}
func (n *noOpLogger) Warnw(msg string, keysAndValues ...any)  {}
func (n *noOpLogger) Error(args ...any)                       {}
func (n *noOpLogger) Errorf(template string, args ...any)     {}
func (n *noOpLogger) Errorw(msg string, keysAndValues ...any) {}
func (n *noOpLogger) Fatal(args ...any)                       {}
func (n *noOpLogger) Fatalf(template string, args ...any)     {}
func (n *noOpLogger) Fatalw(msg string, keysAndValues ...any) {}
func (n *noOpLogger) With(args ...any) Logger                 { return n }
