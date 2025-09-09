//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"os"
)

// Your goal: Complete the sumSlice function to return the sum of all integers in the slice.
func sumSlice(numbers []int) int {
	// --- YOUR CODE HERE ---
	// Hint: Use a for...range loop to iterate over the numbers.
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return sum
	// --------------------
}


// Verification code (do not change)
func main() {
	total := sumSlice([]int{10, 20, 30})
	if total == 60 {
		fmt.Println("Success!")
	} else {
		fmt.Printf("Failure! Expected 60, but got %d\n", total)
		os.Exit(1)
	}
}