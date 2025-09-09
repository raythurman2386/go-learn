//go:build ignore
// +build ignore

package main

import (
	"errors"
	"fmt"
	"os"
)

// Goal: Implement mightFail to return an error and handle it in main.

func mightFail() (string, error) {
	return "", errors.New("simulated")
}

func main() {
	if _, err := mightFail(); err != nil {
		fmt.Println("Error occurred")
	} else {
		fmt.Println("No error")
	}
	_ = os.Stdout
}
