package main

import (
	"fmt"
	"go.uber.org/zap"
	"time"
)

// zap 实现极快、结构化、分级的日志记录。

// https://pkg.go.dev/go.uber.org/zap

func main() {

}

// SugaredLogger
func zap1() {
	sugar := zap.NewExample().Sugar()
	defer sugar.Sync()
	sugar.Infow("failed to fetch URL",
		"url", "http://example.com",
		"attempt", 3,
		"backoff", time.Second,
	)
	sugar.Infof("failed to fetch URL: %s", "http://example.com")
}

// Logger
func zap2() {
	logger := zap.NewExample()
	defer logger.Sync()
	logger.Info("failed to fetch URL",
		zap.String("url", "http://example.com"),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)
}

// Logger 和 SugaredLogger 之间转换
func zap3() {
	logger := zap.NewExample()
	defer logger.Sync()
	sugar := logger.Sugar()
	plain := sugar.Desugar()

	fmt.Println(plain)
}
