package lessons

// DataStructuresLesson returns the lesson for slices, arrays, and maps
func DataStructuresLesson() Lesson {
	return Lesson{
		Title:       "Data Structures",
		Explanation: `Slices, arrays, and maps are the main collection types. Slices are dynamic and most commonly used. Maps are key/value stores.`,
		Example: `package main
import "fmt"

func main() {
    nums := []int{1,2,3}
    nums = append(nums, 4)
    m := map[string]int{"a": 1}
    fmt.Println(nums, m)
}`,
		ProblemFile: "internal/problems/8_datastructures.go",
		TestCmd:     []string{"go", "-tags", "ignore", "run", "internal/problems/8_datastructures.go"},
	}
}
