//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"os"
)

// Your goal: Define a Rectangle struct with Width and Height fields.
// Then implement the Area method to calculate the area of the rectangle.

type Rectangle struct {
	// --- YOUR CODE HERE ---
	// Add Width and Height fields of type float64
	// --------------------
}

// Area method should return the area (Width * Height)
func (r Rectangle) Area() float64 {
	// --- YOUR CODE HERE ---
	return 0 // This is incorrect, fix it!
	// --------------------
}

// Verification code (do not change)
func main() {
	rect := Rectangle{Width: 10, Height: 5}
	area := rect.Area()
	if area != 50 {
		fmt.Printf("Area calculation failed: expected 50, got %f\n", area)
		os.Exit(1)
	}

	fmt.Println("Success!")
}
