package lessons

// StructsLesson returns the lesson for structs
func StructsLesson() Lesson {
	return Lesson{
		Title: "Structs",
		Explanation: `
Structs are Go's way of creating custom data types. They group related data together.
You can attach methods to structs to define their behavior.
Example: type Person struct { Name string; Age int }`,
		Example: `
package main
import "fmt"

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func main() {
	rect := Rectangle{Width: 10, Height: 5}
	fmt.Println("Area:", rect.Area()) // Prints 50
}`,
		ProblemFile: "internal/problems/4_structs.go",
		TestCmd:     []string{"go", "-tags", "ignore", "run", "internal/problems/4_structs.go"},
	}
}
