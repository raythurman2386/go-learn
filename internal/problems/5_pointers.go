//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"os"
)

// Your goal: Complete the increment function to increment the value pointed to by the pointer.
// Then complete the swap function to swap the values of two pointers.

func increment(ptr *int) {
	// --- YOUR CODE HERE ---
	// Increment the value that ptr points to
	// --------------------
}

func swap(a, b *int) {
	// --- YOUR CODE HERE ---
	// Swap the values that a and b point to
	// --------------------
}

// Verification code (do not change)
func main() {
	x := 5
	increment(&x)
	if x != 6 {
		fmt.Printf("increment failed: expected 6, got %d\n", x)
		os.Exit(1)
	}

	a, b := 10, 20
	swap(&a, &b)
	if a != 20 || b != 10 {
		fmt.Printf("swap failed: expected a=20, b=10, got a=%d, b=%d\n", a, b)
		os.Exit(1)
	}

	fmt.Println("Success!")
}
