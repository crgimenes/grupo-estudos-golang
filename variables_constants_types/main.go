package main

import "fmt"

type LogLevel int

const (
	LevelDebug LogLevel = iota
	LevelInfo
	LevelWarn
	LevelError
)

func describeLevel(level LogLevel) string {
	switch level {
	case LevelDebug:
		return "debug"
	case LevelInfo:
		return "info"
	case LevelWarn:
		return "warn"
	case LevelError:
		return "error"
	default:
		return "unknown"
	}
}

func main() {
	short := 10
	var inferred = 20.5
	var zeroInt int
	converted := int(inferred)

	fmt.Printf("short=%d inferred=%.1f zeroInt=%d converted=%d level=%s\n", short, inferred, zeroInt, converted, describeLevel(LevelInfo))
}
