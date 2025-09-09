//go:build ignore
// +build ignore

package main

import (
	"encoding/json"
	"fmt"
)

// Goal: Marshal a simple struct to JSON and print it.

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	u := User{ID: 1, Name: "gopher"}
	b, _ := json.Marshal(u)
	fmt.Println(string(b))
}
