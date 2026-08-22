package main

import (
	"bytes"
	"fmt"
	"log/slog"
)

func NewBufferLogger() (*slog.Logger, *bytes.Buffer) {
	buf := new(bytes.Buffer)
	handler := slog.NewJSONHandler(buf, &slog.HandlerOptions{
		Level: slog.LevelInfo,
		ReplaceAttr: func(groups []string, attr slog.Attr) slog.Attr {
			if attr.Key == slog.TimeKey {
				return slog.Attr{}
			}
			return attr
		},
	})
	return slog.New(handler), buf
}

func LogRequest(logger *slog.Logger, path string, status int, durationMS int64) {
	logger.Info("request", "path", path, "status", status, "duration_ms", durationMS)
}

func main() {
	logger, buf := NewBufferLogger()
	LogRequest(logger, "/health", 200, 3)
	fmt.Print(buf.String())
}
