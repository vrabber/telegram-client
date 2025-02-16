package setup

import (
	"log/slog"
	"testing"
)

func TestConfigureLogLevel(t *testing.T) {
	tests := []struct {
		name     string
		level    string
		expected slog.Level
	}{
		{
			name:     "DEBUG level",
			level:    "DEBUG",
			expected: slog.LevelDebug,
		},
		{
			name:     "INFO level",
			level:    "INFO",
			expected: slog.LevelInfo,
		},
		{
			name:     "WARN level",
			level:    "WARN",
			expected: slog.LevelWarn,
		},
		{
			name:     "ERROR level",
			level:    "ERROR",
			expected: slog.LevelError,
		},
		{
			name:     "Unknown level",
			level:    "UNKNOWN",
			expected: slog.LevelInfo,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ConfigureLogLevel(tt.level)

			if !slog.Default().Handler().Enabled(nil, tt.expected) {
				t.Errorf("ConfigureLogLevel(%s) = %v, want %v", tt.level, slog.Default().Handler(), tt.expected)
			}
		})
	}
}
