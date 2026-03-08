package main

import (
	"bytes"
	"log/slog"
)

func NewBufferLogger() (*slog.Logger, *bytes.Buffer) {
	buf := new(bytes.Buffer)
	h := slog.NewJSONHandler(buf, &slog.HandlerOptions{Level: slog.LevelInfo})
	return slog.New(h), buf
}

func LogRequest(logger *slog.Logger, path string, status int, durationMS int64) {
	logger.Info("request", "path", path, "status", status, "duration_ms", durationMS)
}
