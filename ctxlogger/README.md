# ctxlogger

Go를 위한 범용 컨텍스트 기반 로거 패키지로, 로거 인터페이스와 컨텍스트 기반 로거 관리 기능을 제공합니다.

## 주요 기능

- 컨텍스트 기반 로거 저장 및 검색
- 내장 Zap 로거 구현
- 로거가 없을 때 사용되는 no-op 로거 폴백

## 설치

```bash
go get github.com/delightroom/go-shared/ctxlogger
```

## 사용법

### 기본 사용법 (ZapLogger 사용 예시)

```go
package main

import (
    "context"
    "github.com/delightroom/go-shared/ctxlogger"
)

func main() {
    // 새로운 ZapLogger 로거 생성
    zapLogger := zap.New().Sugar()
    
    // 컨텍스트에 로거 추가
    ctx := ctxlogger.ContextWithLogger(context.Background(), zapLogger)
    
    // 컨텍스트에서 로거 추출
    logger := ctxlogger.LoggerFromContext(ctx, nil)
    logger.Info("컨텍스트 로거에서 안녕하세요")
    
    // 추가 필드가 있는 자식 로거 생성
    childLogger := logger.With("request_id", "123")
    childLogger.Info("요청 처리 중")
}
```

### 함수에서 사용법

```go
func ProcessRequest(ctx context.Context) {
    // 컨텍스트에서 로거 가져오기 (폴백 포함)
    logger := ctxlogger.LoggerFromContext(ctx, nil)
    
    logger.Info("요청 처리 시작")
    logger.Debug("디버그 정보", "key", "value")
    
    if err := doSomething(); err != nil {
        log.Error("처리 실패", "error", err)
    }
}
```

## 로거 인터페이스

다음 인터페이스를 구현하는 모든 로거는 ctxlogger와 함께 사용할 수 있습니다.

```go
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
```