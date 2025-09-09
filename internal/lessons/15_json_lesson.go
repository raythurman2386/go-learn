package lessons

// JSONLesson returns the lesson for working with JSON
func JSONLesson() Lesson {
	return Lesson{
		Title:       "JSON",
		Explanation: `encoding/json marshals and unmarshals Go structs. Use struct tags to control field names and omitempty.`,
		Example: `package main
import (
    "encoding/json"
    "fmt"
)
type User struct { ID int ` + "`json:\"id\"`" + `; Name string ` + "`json:\"name\"`" + ` }
func main(){ u := User{ID:1, Name:"g"}; b,_:=json.Marshal(u); fmt.Println(string(b)) }`,
		ProblemFile: "internal/problems/15_json.go",
		TestCmd:     []string{"go", "-tags", "ignore", "run", "internal/problems/15_json.go"},
	}
}
