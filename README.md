# Go Learning Environment

[![Go Version](https://img.shields.io/badge/Go-1.25+-00ADD8?style=flat&logo=go)](https://golang.org/)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)

An interactive CLI application for learning Go programming concepts through hands-on practice. This project provides a structured learning experience with explanations, examples, and coding challenges.

## ✨ Features

- **Interactive Learning**: Step through Go concepts with guided explanations
- **Hands-on Practice**: Solve real coding problems in your editor
- **AI-Powered Explanations**: Optional AI-generated comprehensive explanations using Google Gemini
- **Progressive Difficulty**: 16 lessons covering fundamentals to advanced topics
- **Immediate Feedback**: Built-in testing for your solutions

## 📚 What You'll Learn

- Variables and Types
- Slices and Arrays
- Functions and Methods
- Structs and Interfaces
- Pointers and Memory Management
- Packages and Modules
- Control Flow (if/else, loops, switch)
- Data Structures (maps, custom types)
- Error Handling
- Concurrency (goroutines, channels)
- Generics
- Context and Cancellation
- Structured Logging
- JSON Processing
- Testing and Benchmarking

## 🚀 Quick Start

### Prerequisites

- **Go 1.25+** installed and available on your PATH
- **Text Editor** (VS Code, GoLand, Vim, etc.)

### Installation

```bash
# Clone the repository
git clone https://github.com/yourusername/go-learn.git
cd go-learn

# Run the application
go run ./cmd
```

### Optional: AI-Generated Explanations

Enhance your learning experience with AI-powered explanations:

1. Get a Google Gemini API key from [Google AI Studio](https://makersuite.google.com/app/apikey)
2. Set the environment variable:
   ```bash
   export GEMINI_API_KEY="your-api-key-here"
   ```
3. Restart the application

**Features:**
- Comprehensive explanations with practical examples
- Real-time content generation
- Automatic fallback to static explanations if AI is unavailable
- Beginner-friendly content tailored for Go learners

## 🎯 How It Works

1. **Choose a Topic**: Select from 16 available lessons
2. **Learn**: Read explanations and study examples
3. **Practice**: Solve coding challenges in your editor
4. **Test**: Run automated tests to verify your solutions
5. **Repeat**: Continue with the next topic

### Example Session

```
--- 🚀 Go Learning Environment 🚀 ---
Choose a topic to learn, or type 'q' to quit:
1. Variables and Types
2. Slices
...
> 1

--- Topic: Variables and Types ---
1. Explain 🧠
2. Show Example ✨
3. Solve Problem 💻
b. Back to main menu
> 1

--- AI-Generated Explanation ---
[Comprehensive explanation with examples and best practices]
```

## 🏗️ Project Structure

```
go-learn/
├── cmd/
│   └── main.go                 # CLI entry point
├── internal/
│   ├── lessons/
│   │   ├── lessons.go          # Lesson registry and runner
│   │   ├── 1_variables_lesson.go
│   │   ├── 2_slices_lesson.go
│   │   └── ...                 # Individual lesson files
│   └── problems/
│       ├── 1_variables.go
│       ├── 2_slices.go
│       └── ...                 # Practice problem files
├── go.mod
├── go.sum
├── README.md
├── CONTRIBUTING.md
└── LICENSE
```

### Adding New Lessons

See [CONTRIBUTING.md](CONTRIBUTING.md) for detailed instructions on adding new lessons.

## 🛠️ Development

### Building

```bash
# Build only the CLI (recommended)
go build -o go-learn ./cmd

# Build the entire workspace
go build ./...
```

### Running Tests (Enventual upgrade)

```bash
# Test the entire project
go test ./...

# Test specific packages
go test ./internal/lessons
```

## 🏗️ Building for Multiple Platforms

Go makes it easy to cross-compile for different operating systems and architectures. Here are the commands to build binaries for Windows, Linux, and macOS:

### Build Scripts

For convenience, we've included build scripts for different platforms:

#### Linux/macOS:
```bash
# Make the script executable and run it
chmod +x build.sh
./build.sh
```

#### Windows:
```powershell
# Run the PowerShell script
.\build.ps1

# Or with options
.\build.ps1 -Release  # Also create release archives
.\build.ps1 -Clean    # Clean before building
```

#### Using GNU Make on Windows

If you prefer using `make` commands on Windows, you can install GNU Make via winget:

```bash
# Install GNU Make
winget install ezwinports.make

# Restart your terminal/command prompt to refresh PATH
# Then use make commands normally:
make help          # Show available targets
make windows       # Build Windows binary
make linux         # Build Linux binary
make release       # Build all platforms + create archives
```

### Build Output

After running any build script, you'll find the binaries in the `bin/` directory:

```
bin/
├── go-learn-windows-amd64.exe    # Windows 64-bit
├── go-learn-windows-386.exe      # Windows 32-bit
├── go-learn-linux-amd64          # Linux 64-bit
├── go-learn-linux-386            # Linux 32-bit
├── go-learn-linux-arm64          # Linux ARM64
├── go-learn-darwin-amd64         # macOS Intel
└── go-learn-darwin-arm64         # macOS Apple Silicon
```

### Build Options Summary

| Method | Platforms | Use Case | Speed |
|--------|-----------|----------|-------|
| `go run ./cmd` | Current platform only | Development | Fastest |
| `go build ./cmd` | Current platform only | Local deployment | Fast |
| `GOOS=... go build` | Single target platform | Specific deployment | Medium |
| `./build.sh` | All platforms | Distribution | Slower |
| `make release` | All platforms + archives | Official releases | Slowest |

## 🤝 Contributing

We welcome contributions! Please see our [Contributing Guide](CONTRIBUTING.md) for details on:

- Adding new lessons
- Improving existing content
- Reporting bugs
- Feature requests
- Code style guidelines

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

- Built with ❤️ for the Go community
- Inspired by interactive learning platforms
- Powered by Google Gemini for AI explanations

## 📢 Roadmap

[ ] Add testing and proper coverage
[ ] Additional Go lessons eventually
[ ] Update Gemini to pull docs link for fully accurate data
[ ] Make problems solvable in the cli
[ ] 🤷 Who knows what else, may not even make it to these items this was a random idea that turned out pretty fun

---

**Happy Learning!** 🎉

*Found this helpful? ⭐ Star the repo and share with fellow Go developers!* 
