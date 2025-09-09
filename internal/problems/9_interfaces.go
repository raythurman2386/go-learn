//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"os"
)

// Goal: Implement Speaker interface for a Dog type and print its Speak output.

type Speaker interface{ Speak() string }

type Dog struct{}

func (Dog) Speak() string { return "woof" }

func main() {
	var s Speaker = Dog{}
	fmt.Println(s.Speak())
	_ = os.Stdout
}
