package lessons

// FunctionsLesson returns the lesson for functions
func FunctionsLesson() Lesson {
	return Lesson{
		Title: "Functions",
		Explanation: `
Functions in Go are defined with the 'func' keyword. They can take parameters and return values.
You can have multiple return values, which is useful for error handling.
Example: func add(a, b int) int { return a + b }`,
		Example: `
package main
import "fmt"

func multiply(a, b int) int {
	return a * b
}

func isEven(n int) bool {
	return n%2 == 0
}

func main() {
	fmt.Println(multiply(4, 5)) // Prints 20
	fmt.Println(isEven(4))      // Prints true
	fmt.Println(isEven(5))      // Prints false
}`,
		ProblemFile: "internal/problems/3_functions.go",
		TestCmd:     []string{"go", "-tags", "ignore", "run", "internal/problems/3_functions.go"},
	}
}
