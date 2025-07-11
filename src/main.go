package main

import (
	"log/slog"
	"os"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	child := logger.WithGroup("info").With(slog.String("method", "GET"))
	logger.Info("Hello - Logger")
	child.Debug("Hello - Logger", slog.String("key string", "value"), slog.Int("key int", 100), slog.Group("key group", slog.String("nested 1", "hello"), slog.String("nested 2", "world")))
	logger.Warn("Hello - Logger")
	child.Error("Hello - Logger")
}
