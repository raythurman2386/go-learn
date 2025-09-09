package lessons

// ErrorsLesson returns the lesson for error handling
func ErrorsLesson() Lesson {
	return Lesson{
		Title:       "Error Handling",
		Explanation: `Go returns errors as the last return value. Check against nil to detect errors. Use errors.Is and errors.As for advanced checks.`,
		Example: `package main
import (
    "errors"
    "fmt"
)
func mightFail() (string, error) { return "", errors.New("fail") }
func main(){ if _, err := mightFail(); err != nil { fmt.Println("error:", err) } }`,
		ProblemFile: "internal/problems/10_errors.go",
		TestCmd:     []string{"go", "-tags", "ignore", "run", "internal/problems/10_errors.go"},
	}
}
