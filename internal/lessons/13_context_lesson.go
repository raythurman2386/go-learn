package lessons

// ContextLesson returns the lesson for context
func ContextLesson() Lesson {
	return Lesson{
		Title:       "Context",
		Explanation: `The context package carries deadlines, cancellations, and request-scoped values across API boundaries and goroutines.`,
		Example: `package main
import (
    "context"
    "fmt"
    "time"
)
func main(){ ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*100); defer cancel(); select{case <-time.After(time.Second): fmt.Println("done"); case <-ctx.Done(): fmt.Println("canceled") } }`,
		ProblemFile: "internal/problems/13_context.go",
		TestCmd:     []string{"go", "-tags", "ignore", "run", "internal/problems/13_context.go"},
	}
}
