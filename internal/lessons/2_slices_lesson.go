package lessons

// SlicesLesson returns the lesson for slices
func SlicesLesson() Lesson {
	return Lesson{
		Title: "Slices",
		Explanation: `
Slices are Go's version of dynamic arrays (like Python lists).
They are more common and flexible than arrays. You can add items using the append() function.
Example: numbers := []int{10, 20}
         numbers = append(numbers, 30)`,
		Example: `
package main
import "fmt"

func main() {
	names := []string{"Alice", "Bob"}
	names = append(names, "Charlie")
	fmt.Println("Slice:", names)       // Prints [Alice Bob Charlie]
	fmt.Println("Length:", len(names)) // Prints 3
}`,
		ProblemFile: "internal/problems/2_slices.go",
		TestCmd:     []string{"go", "-tags", "ignore", "run", "internal/problems/2_slices.go"},
	}
}
