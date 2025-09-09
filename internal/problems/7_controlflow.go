//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"os"
)

// Goal: Print numbers 0..2 each on their own line using a for loop.

func main() {
	// --- YOUR CODE HERE ---
	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}
	// ----------------------

	// Verification (do not change)
	_ = os.Stdout
}
