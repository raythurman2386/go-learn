//go:build ignore
// +build ignore

package main

import (
	"fmt"
)

// Goal: Implement SumInts generic-like function (use concrete for compatibility).

func SumInts(s []int) int {
	var sum int
	for _, v := range s {
		sum += v
	}
	return sum
}

func main() {
	fmt.Println(SumInts([]int{1, 2, 3}))
}
