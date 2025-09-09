package lessons

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"google.golang.org/genai"
)

// Lesson defines a single learning unit
type Lesson struct {
	Title       string
	Explanation string
	Example     string
	ProblemFile string   // The file for the user to solve
	TestCmd     []string // Command to verify the solution
}

// A map to hold all our lessons, keyed by number
var lessons = make(map[string]Lesson)

func initializeLessons() {
	lessons["1"] = VariablesLesson()
	lessons["2"] = SlicesLesson()
	lessons["3"] = FunctionsLesson()
	lessons["4"] = StructsLesson()
	lessons["5"] = PointersLesson()
	lessons["6"] = PackagesLesson()
	lessons["7"] = ControlFlowLesson()
	lessons["8"] = DataStructuresLesson()
	lessons["9"] = InterfacesLesson()
	lessons["10"] = ErrorsLesson()
	lessons["11"] = ConcurrencyLesson()
	lessons["12"] = GenericsLesson()
	lessons["13"] = ContextLesson()
	lessons["14"] = LoggingLesson()
	lessons["15"] = JSONLesson()
	lessons["16"] = TestingLesson()
}

func GetLessons() map[string]Lesson {
	initializeLessons()
	return lessons
}

func RunLesson(lesson Lesson, reader *bufio.Reader) {
	first := true
	for {
		// On first entry, show the explanation automatically
		if first {
			fmt.Println("\n--- Explanation ---")
			fmt.Println(lesson.Explanation)
			first = false
		}

		fmt.Printf("\n--- Topic: %s ---\n", lesson.Title)
		fmt.Println("1. Explain 🧠")
		fmt.Println("2. Show Example ✨")
		fmt.Println("3. Solve Problem 💻")
		fmt.Println("b. Back to main menu")
		fmt.Print("> ")

		input, _ := reader.ReadString('\n')
		choice := strings.TrimSpace(input)

		switch choice {
		case "1":
			fmt.Println("\n--- AI-Generated Explanation ---")
			aiExplanation, err := generateAIExplanation(lesson.Title)
			if err != nil {
				fmt.Printf("Could not generate AI explanation (%v)\n", err)
				fmt.Println("Falling back to static explanation:")
				fmt.Println("\n--- Explanation ---")
				fmt.Println(lesson.Explanation)
			} else {
				fmt.Println(aiExplanation)
			}
		case "2":
			fmt.Println("\n--- Example ---")
			fmt.Println(lesson.Example)
		case "3":
			solveProblem(lesson, reader)
		case "b":
			return
		default:
			fmt.Println("Invalid choice.")
		}
	}
}

func solveProblem(lesson Lesson, reader *bufio.Reader) {
	fmt.Println("\n--- 📝 Problem ---")
	fmt.Printf("Open the file '%s' in your editor and complete the code.\n", lesson.ProblemFile)
	fmt.Println("Once you are done, save the file and press [Enter] here to test your solution.")
	reader.ReadString('\n')

	fmt.Println("Running tests...")
	cmd := exec.Command(lesson.TestCmd[0], lesson.TestCmd[1:]...)
	output, err := cmd.CombinedOutput()

	fmt.Println("\n--- ✅ Results ---")
	fmt.Println(string(output))
	if err != nil {
		fmt.Println("Looks like there's an error. Keep trying! You got this. 👍")
	} else {
		fmt.Println("🎉 Success! Great job! You solved it correctly. 🎉")
	}
}

// generateAIExplanation uses Google Gemini to generate a comprehensive explanation
func generateAIExplanation(lessonTitle string) (string, error) {
	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		return "", fmt.Errorf("GEMINI_API_KEY environment variable not set")
	}

	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey: apiKey,
	})
	if err != nil {
		return "", fmt.Errorf("failed to create Gemini client: %v", err)
	}

	prompt := fmt.Sprintf(`Provide a comprehensive and beginner-friendly explanation of "%s" in Go programming. 
Include:
1. What it is and why it's important
2. Basic syntax and usage examples
3. Common patterns and best practices
4. Potential pitfalls to avoid
5. Real-world use cases

Make it educational and easy to understand for someone learning Go.`, lessonTitle)

	resp, err := client.Models.GenerateContent(ctx, "gemini-2.5-flash-lite", genai.Text(prompt), nil)
	if err != nil {
		return "", fmt.Errorf("failed to generate content: %v", err)
	}

	if resp == nil || len(resp.Candidates) == 0 {
		return "", fmt.Errorf("no content generated")
	}

	// Use the Text() method to get the generated text
	return resp.Text(), nil
}
