package lessons

// PointersLesson returns the lesson for pointers
func PointersLesson() Lesson {
	return Lesson{
		Title: "Pointers",
		Explanation: `
Pointers hold the memory address of a value. The & operator gets the address of a variable,
and the * operator dereferences a pointer to get the value it points to.
Pointers are useful for modifying values in functions and working with large data structures efficiently.`,
		Example: `
package main
import "fmt"

func increment(ptr *int) {
	*ptr++ // Increment the value that ptr points to
}

func swap(a, b *int) {
	*a, *b = *b, *a // Swap the values
}

func main() {
	x := 5
	increment(&x)
	fmt.Println(x) // Prints 6

	a, b := 10, 20
	swap(&a, &b)
	fmt.Println(a, b) // Prints 20 10
}`,
		ProblemFile: "internal/problems/5_pointers.go",
		TestCmd:     []string{"go", "-tags", "ignore", "run", "internal/problems/5_pointers.go"},
	}
}
