package lessons

// LoggingLesson returns the lesson for structured logging
func LoggingLesson() Lesson {
	return Lesson{
		Title:       "Structured Logging",
		Explanation: `Use structured logging for machine-readable logs. The slog package (Go 1.21) provides JSON handlers.`,
		Example: `package main
import (
    "log/slog"
    "os"
)
func main(){ logger := slog.New(slog.NewJSONHandler(os.Stdout, nil)); logger.Info("hello", "user", "gopher") }`,
		ProblemFile: "internal/problems/14_logging.go",
		TestCmd:     []string{"go", "-tags", "ignore", "run", "internal/problems/14_logging.go"},
	}
}
