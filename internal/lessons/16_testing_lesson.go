package lessons

// TestingLesson returns the lesson for tests
func TestingLesson() Lesson {
	return Lesson{
		Title:       "Testing",
		Explanation: `Tests live in _test.go files and use the testing package. Run with go test.`,
		Example: `package main
import "testing"
func TestSum(t *testing.T){ if 2+2!=4 { t.Fatal("math broken") } }`,
		ProblemFile: "internal/problems/16_testing.go",
		TestCmd:     []string{"go", "test", "./..."},
	}
}
