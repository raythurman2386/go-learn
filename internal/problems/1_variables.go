//go:build ignore
// +build ignore

package main

import "fmt"

func main() {
	// Your goal: Declare a variable 'fruit' of type string and assign it the value "apple".
	// Then declare a variable 'quantity' of type int and assign it the value 5.
	// Use the short := syntax.

	// --- YOUR CODE HERE ---
	fruit := "banana"  // This is incorrect, fix it!
	quantity := 0     // This is incorrect, fix it!
	// --------------------


	// Verification code (do not change)
	if fruit == "apple" && quantity == 5 {
		fmt.Println("Success!")
	} else {
		fmt.Println("Not quite, try again!")
		// This os.Exit(1) is what makes the test fail in the main app
		os.Exit(1)
	}
}
