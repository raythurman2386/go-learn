//go:build ignore
// +build ignore

package main

import (
	"log/slog"
	"os"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	logger.Info("test log", "key", "value")
}
