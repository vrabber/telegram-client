package setup

import "log/slog"

func ConfigureLogLevel(level string) {
	switch level {
	case "DEBUG":
		slog.SetLogLoggerLevel(slog.LevelDebug)
	case "INFO":
		slog.SetLogLoggerLevel(slog.LevelInfo)
	case "WARN":
		slog.SetLogLoggerLevel(slog.LevelWarn)
	case "ERROR":
		slog.SetLogLoggerLevel(slog.LevelError)
	default:
		slog.Warn("unknown log level, set INFO by default", "level", level)
		slog.SetLogLoggerLevel(slog.LevelInfo)
	}
}
