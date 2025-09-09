package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"go-learn/internal/lessons"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	lessonsMap := lessons.GetLessons()

	// Main application loop
	for {
		printMainMenu(lessonsMap)
		input, _ := reader.ReadString('\n')
		choice := strings.TrimSpace(input)

		if choice == "q" {
			fmt.Println("Happy coding! 👋")
			break
		}

		if lesson, ok := lessonsMap[choice]; ok {
			lessons.RunLesson(lesson, reader)
		} else {
			fmt.Println("Invalid choice, please try again.")
		}
	}
}

func printMainMenu(lessonsMap map[string]lessons.Lesson) {
	fmt.Println("\n--- 🚀 Go Learning Environment 🚀 ---")
	fmt.Println("Choose a topic to learn, or type 'q' to quit:")
	// Ensure numeric ordering of lessons
	keys := make([]string, 0, len(lessonsMap))
	for k := range lessonsMap {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		// Try numeric comparison when keys are numeric strings like "1", "2", "10"
		a, errA := strconv.Atoi(keys[i])
		b, errB := strconv.Atoi(keys[j])
		if errA == nil && errB == nil {
			return a < b
		}
		// Fallback to lexicographic if not both integers
		return keys[i] < keys[j]
	})
	for _, key := range keys {
		lesson := lessonsMap[key]
		fmt.Printf("%s. %s\n", key, lesson.Title)
	}
	fmt.Print("> ")
}
