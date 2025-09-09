package lessons

// ControlFlowLesson returns the lesson for control flow
func ControlFlowLesson() Lesson {
	return Lesson{
		Title:       "Control Flow",
		Explanation: `if/else, for loops (including range), and switch are the primary control flow constructs in Go.`,
		Example: `package main
import "fmt"

func main() {
    for i := 0; i < 3; i++ {
        fmt.Println(i)
    }
}`,
		ProblemFile: "internal/problems/7_controlflow.go",
		TestCmd:     []string{"go", "-tags", "ignore", "run", "internal/problems/7_controlflow.go"},
	}
}
