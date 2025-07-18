package ctxlogger

import (
	"context"
	"testing"
)

// testLogger 는 Logger 인터페이스를 구현하는 테스트 로거입니다.
type testLogger struct{}

func (l *testLogger) Debug(args ...any)                       {}
func (l *testLogger) Debugf(template string, args ...any)     {}
func (l *testLogger) Debugw(msg string, keysAndValues ...any) {}
func (l *testLogger) Info(args ...any)                        {}
func (l *testLogger) Infof(template string, args ...any)      {}
func (l *testLogger) Infow(msg string, keysAndValues ...any)  {}
func (l *testLogger) Warn(args ...any)                        {}
func (l *testLogger) Warnf(template string, args ...any)      {}
func (l *testLogger) Warnw(msg string, keysAndValues ...any)  {}
func (l *testLogger) Error(args ...any)                       {}
func (l *testLogger) Errorf(template string, args ...any)     {}
func (l *testLogger) Errorw(msg string, keysAndValues ...any) {}
func (l *testLogger) Fatal(args ...any)                       {}
func (l *testLogger) Fatalf(template string, args ...any)     {}
func (l *testLogger) Fatalw(msg string, keysAndValues ...any) {}
func (l *testLogger) With(args ...any) Logger                 { return l }

func TestContextWithLogger(t *testing.T) {
	logger := &testLogger{}

	ctx := context.Background()
	ctx = ContextWithLogger(ctx, logger)

	retrieved := LoggerFromContext(ctx)
	if retrieved == nil {
		t.Fatal("expected logger to be retrieved from context")
	}

	if _, ok := retrieved.(*testLogger); !ok {
		t.Fatal("expected retrieved logger to be testLogger")
	}
}

func TestLoggerFromContext_FallbackLogger(t *testing.T) {
	ctx := context.Background()
	retrieved := LoggerFromContext(ctx)

	if retrieved == nil {
		t.Fatal("expected no-op logger to be returned")
	}

	if _, ok := retrieved.(*noOpLogger); !ok {
		t.Fatal("expected retrieved logger to be noOpLogger")
	}
}

func TestNoOpLogger(t *testing.T) {
	logger := &noOpLogger{}

	// Test that all methods can be called without panic
	logger.Debug("test")
	logger.Debugf("test %s", "arg")
	logger.Debugw("test", "key", "value")
	logger.Info("test")
	logger.Infof("test %s", "arg")
	logger.Infow("test", "key", "value")
	logger.Warn("test")
	logger.Warnf("test %s", "arg")
	logger.Warnw("test", "key", "value")
	logger.Error("test")
	logger.Errorf("test %s", "arg")
	logger.Errorw("test", "key", "value")

	// With should return itself
	withLogger := logger.With("key", "value")
	if withLogger != logger {
		t.Fatal("expected With to return same instance")
	}
}
