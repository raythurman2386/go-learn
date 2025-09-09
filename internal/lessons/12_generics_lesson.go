package lessons

// GenericsLesson returns the lesson for generics
func GenericsLesson() Lesson {
	return Lesson{
		Title:       "Generics",
		Explanation: `Generics let you write type-parameterized functions and types (Go 1.18+). Use constraints to limit types.`,
		Example: `package main
import "fmt"
func SumInts[T int | int64](s []T) T { var sum T; for _, v := range s { sum += v }; return sum }
func main(){ fmt.Println(SumInts([]int{1,2,3})) }`,
		ProblemFile: "internal/problems/12_generics.go",
		TestCmd:     []string{"go", "-tags", "ignore", "run", "internal/problems/12_generics.go"},
	}
}
