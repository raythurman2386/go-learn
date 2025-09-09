package lessons

// VariablesLesson returns the lesson for variables and types
func VariablesLesson() Lesson {
	return Lesson{
		Title: "Variables and Types",
		Explanation: `
Go is a statically typed language. You declare variables with a specific type.
The most common way is using the := short declaration statement, which infers the type.
Example: name := "Alice" // Go infers this is a string`,
		Example: `
var name string = "Alice" // Long form
age := 30                 // Short form (inferred int)
isReady := true           // Short form (inferred bool)
fmt.Println(name, age, isReady)`,
		ProblemFile: "internal/problems/1_variables.go",
		TestCmd:     []string{"go", "-tags", "ignore", "run", "internal/problems/1_variables.go"},
	}
}
