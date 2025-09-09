package lessons

// ConcurrencyLesson returns the lesson for goroutines and channels
func ConcurrencyLesson() Lesson {
	return Lesson{
		Title:       "Concurrency",
		Explanation: `Goroutines are lightweight threads started with the 'go' keyword. Channels communicate between goroutines. Use sync.WaitGroup to wait for goroutines.`,
		Example: `package main
import (
    "fmt"
    "sync"
)
func main(){ var wg sync.WaitGroup; wg.Add(1); go func(){ defer wg.Done(); fmt.Println("hi") }(); wg.Wait() }`,
		ProblemFile: "internal/problems/11_concurrency.go",
		TestCmd:     []string{"go", "-tags", "ignore", "run", "internal/problems/11_concurrency.go"},
	}
}
