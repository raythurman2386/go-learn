package lessons

// PackagesLesson returns the lesson for packages & imports
func PackagesLesson() Lesson {
	return Lesson{
		Title: "Packages & Imports",
		Explanation: `Go code is organized into packages. The special package ` + "`main`" + ` is the program entry point.
You import standard and third-party packages using the import statement. Exported names are capitalized.
Example: import "fmt"`,
		Example: `package main
import "fmt"

func main() {
    fmt.Println("Hello from packages lesson")
}`,
		ProblemFile: "internal/problems/6_packages.go",
		TestCmd:     []string{"go", "-tags", "ignore", "run", "internal/problems/6_packages.go"},
	}
}
