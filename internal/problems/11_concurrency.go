//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"sync"
)

// Goal: Start a goroutine that prints "worker" and use WaitGroup to wait.

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("worker")
	}()
	wg.Wait()
}
