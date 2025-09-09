package lessons

// InterfacesLesson returns the lesson for interfaces
func InterfacesLesson() Lesson {
	return Lesson{
		Title:       "Interfaces",
		Explanation: `Interfaces define behavior via method sets. Types implement interfaces implicitly by implementing the methods.`,
		Example: `package main
import "fmt"

type Speaker interface{ Speak() string }
type Dog struct{}
func (Dog) Speak() string { return "woof" }

func main(){ var s Speaker = Dog{}; fmt.Println(s.Speak()) }`,
		ProblemFile: "internal/problems/9_interfaces.go",
		TestCmd:     []string{"go", "-tags", "ignore", "run", "internal/problems/9_interfaces.go"},
	}
}
