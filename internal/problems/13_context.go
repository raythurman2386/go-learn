//go:build ignore
// +build ignore

package main

import (
	"context"
	"fmt"
	"time"
)

// Goal: Use a context with timeout and observe cancellation.

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()
	select {
	case <-time.After(50 * time.Millisecond):
		fmt.Println("done")
	case <-ctx.Done():
		fmt.Println("canceled")
	}
}
