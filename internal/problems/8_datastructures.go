//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"os"
)

// Goal: Create a slice with 1,2,3 and append 4. Create a map with key "a" -> 1.

func main() {
	// --- YOUR CODE HERE ---
	nums := []int{1, 2, 3}
	nums = append(nums, 4)
	m := map[string]int{"a": 1}
	fmt.Println(nums, m)
	// ----------------------

	_ = os.Stdout
}
