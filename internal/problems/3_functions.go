//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"os"
)

// Your goal: Complete the multiply function to return the product of two numbers.
// Then complete the isEven function to return true if the number is even.

func multiply(a, b int) int {
	// --- YOUR CODE HERE ---
	return 0 // This is incorrect, fix it!
	// --------------------
}

func isEven(n int) bool {
	// --- YOUR CODE HERE ---
	return false // This is incorrect, fix it!
	// --------------------
}

// Verification code (do not change)
func main() {
	product := multiply(4, 5)
	if product != 20 {
		fmt.Printf("multiply function failed: expected 20, got %d\n", product)
		os.Exit(1)
	}

	if !isEven(4) {
		fmt.Println("isEven function failed: 4 should be even")
		os.Exit(1)
	}

	if isEven(5) {
		fmt.Println("isEven function failed: 5 should be odd")
		os.Exit(1)
	}

	fmt.Println("Success!")
}
