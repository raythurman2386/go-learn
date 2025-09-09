# Contributing to Go Learning Environment

Thank you for your interest in contributing to the Go Learning Environment! 🎉

This document provides guidelines and information for contributors. Whether you're fixing bugs, adding features, improving documentation, or helping with lessons, your contributions are welcome and appreciated.

## 📋 Table of Contents

- [Code of Conduct](#code-of-conduct)
- [Getting Started](#getting-started)
- [Development Setup](#development-setup)
- [How to Contribute](#how-to-contribute)
- [Adding New Lessons](#adding-new-lessons)
- [Code Style Guidelines](#code-style-guidelines)
- [Testing](#testing)
- [Pull Request Process](#pull-request-process)
- [Reporting Issues](#reporting-issues)

## 🤝 Code of Conduct

This project follows a code of conduct to ensure a welcoming environment for all contributors. By participating, you agree to:

- Be respectful and inclusive
- Focus on constructive feedback
- Accept responsibility for mistakes
- Show empathy towards other contributors
- Help create a positive community

## 🚀 Getting Started

1. **Fork** the repository on GitHub
2. **Clone** your fork locally:
   ```bash
   git clone https://github.com/yourusername/go-learn.git
   cd go-learn
   ```
3. **Create** a new branch for your changes:
   ```bash
   git checkout -b feature/your-feature-name
   ```

## 🛠️ Development Setup

### Prerequisites

- Go 1.18 or later
- Git
- A text editor or IDE (we recommend VS Code with Go extensions)

### Setup Steps

1. **Install dependencies**:
   ```bash
   go mod download
   ```

2. **Verify installation**:
   ```bash
   go run ./cmd
   ```

3. **Run tests**:
   ```bash
   go test ./...
   ```

### Optional: AI Features

If you want to work on AI-related features:

1. Get a Google Gemini API key from [Google AI Studio](https://makersuite.google.com/app/apikey)
2. Set the environment variable:
   ```bash
   export GEMINI_API_KEY="your-api-key-here"
   ```

## 💡 How to Contribute

### Types of Contributions

- **🐛 Bug Fixes**: Fix issues in existing code
- **✨ New Features**: Add new functionality
- **📚 Lessons**: Create new learning content
- **📖 Documentation**: Improve docs, README, comments
- **🧪 Tests**: Add or improve test coverage
- **🔧 Tooling**: Improve build process, CI/CD, etc.

### Contribution Workflow

1. **Choose an issue** or **create one** describing what you want to work on
2. **Discuss** the approach if it's a significant change
3. **Implement** your changes
4. **Test** thoroughly
5. **Submit a pull request**

## 📚 Adding New Lessons

The Go Learning Environment supports adding new lessons to expand the learning content. Follow these steps:

### Step 1: Understand the Lesson Structure

Each lesson consists of three main components:

```go
type Lesson struct {
    Title       string   // Display name of the lesson
    Explanation string   // Brief explanation text
    Example     string   // Code example
    ProblemFile string   // Path to practice problem file
    TestCmd     []string // Command to test the solution
}
```

### Step 2: Create the Lesson File

1. **Create a new lesson file** in `internal/lessons/`:
   ```
   internal/lessons/N_topic_lesson.go
   ```
   Where `N` is the next lesson number and `topic` is a descriptive name.

2. **Implement the lesson constructor**:
   ```go
   package lessons

   func TopicLesson() Lesson {
       return Lesson{
           Title: "Your Topic Title",
           Explanation: `Your explanation text here.
           This should be educational and concise.`,
           Example: `package main

   import "fmt"

   func main() {
       // Your example code here
       fmt.Println("Hello, World!")
   }`,
           ProblemFile: "internal/problems/N_topic.go",
           TestCmd: []string{"go", "run", "internal/problems/N_topic.go"},
       }
   }
   ```

### Step 3: Create the Practice Problem

1. **Create a problem file** in `internal/problems/`:
   ```
   internal/problems/N_topic.go
   ```

2. **Write a practice problem**:
   ```go
   package main

   import "fmt"

   func main() {
       // TODO: Implement your solution here
       // This file should contain a coding challenge
       // that teaches the lesson topic

       fmt.Println("Implement me!")
   }
   ```

### Step 4: Register the Lesson

1. **Open** `internal/lessons/lessons.go`
2. **Add your lesson** to the `initializeLessons()` function:
   ```go
   func initializeLessons() {
       lessons["1"] = VariablesLesson()
       lessons["2"] = SlicesLesson()
       // ... existing lessons ...
       lessons["17"] = TopicLesson() // Add your new lesson
   }
   ```

### Step 5: Test Your Lesson

1. **Build and run** the application:
   ```bash
   go run ./cmd
   ```

2. **Navigate** to your new lesson
3. **Test all options**: Explain, Example, Solve Problem
4. **Verify** the problem runs and provides feedback

### Lesson Guidelines

- **Progressive Difficulty**: Ensure lessons build upon previous knowledge
- **Clear Objectives**: Each lesson should have a specific learning goal
- **Practical Examples**: Include real-world applicable code
- **Comprehensive Coverage**: Cover common use cases and pitfalls
- **Beginner-Friendly**: Write explanations for learners new to Go

## 🎨 Code Style Guidelines

### Go Code Style

- Follow the [official Go style guide](https://golang.org/doc/effective_go.html)
- Use `gofmt` to format your code
- Run `go vet` to check for common issues
- Use meaningful variable and function names
- Add comments for complex logic

### Commit Messages

- Use clear, descriptive commit messages
- Start with a verb (Add, Fix, Update, etc.)
- Keep the first line under 50 characters
- Add detailed description if needed

Examples:
```
Add new lesson on error handling
Fix panic in lesson navigation
Update README with better installation instructions
```

### File Organization

- Keep lesson files organized by topic
- Use consistent naming conventions
- Group related functionality together
- Add package comments for new files

## 🧪 Testing

### Running Tests

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run tests for specific package
go test ./internal/lessons
```

### Writing Tests

- Add tests for new functionality
- Use table-driven tests for multiple test cases
- Test both success and error scenarios
- Include edge cases

Example test structure:
```go
func TestLessonCreation(t *testing.T) {
    lesson := NewLesson()

    if lesson.Title == "" {
        t.Error("Lesson title should not be empty")
    }

    // Add more test cases...
}
```

## 🔄 Pull Request Process

1. **Ensure** your code builds and tests pass:
   ```bash
   go build ./cmd
   go test ./...
   ```

2. **Update documentation** if needed (README, comments, etc.)

3. **Create a pull request** with:
   - Clear title describing the changes
   - Detailed description of what was changed and why
   - Reference to any related issues

4. **Wait for review** and address any feedback

5. **Merge** once approved

### PR Checklist

- [ ] Code builds successfully
- [ ] Tests pass
- [ ] Documentation updated
- [ ] Commit messages are clear
- [ ] PR description explains the changes
- [ ] Related issues are referenced

## 🐛 Reporting Issues

### Bug Reports

When reporting bugs, please include:

- **Clear title** describing the issue
- **Steps to reproduce** the problem
- **Expected behavior** vs. actual behavior
- **Environment details** (Go version, OS, etc.)
- **Screenshots** if applicable
- **Error messages** or logs

### Feature Requests

For new features, please provide:

- **Clear description** of the proposed feature
- **Use case** or problem it solves
- **Implementation ideas** if you have them
- **Mockups** or examples if applicable

## 📞 Getting Help

- **Issues**: Use GitHub issues for bugs and feature requests
- **Discussions**: Use GitHub discussions for questions and general discussion
- **Documentation**: Check the README and this contributing guide first

## 🎯 Recognition

Contributors will be recognized in:
- The repository's contributor list
- Release notes for significant contributions
- Special mentions for major features or lesson additions

Thank you for contributing to the Go Learning Environment! Your efforts help make Go learning more accessible for everyone. 🚀
