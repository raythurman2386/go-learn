//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"os"
)

// Goal: Write a main that prints "Hello, packages!" using fmt.Println

func main() {
	// --- YOUR CODE HERE ---
	fmt.Println("Hello, packages!")
	// ----------------------

	// Verification (do not change)
	// The test runner checks stdout; exit non-zero on mismatch.
	// (Simple crude verification: ensure the program runs.)
	_ = os.Stdout
}
