# Go Learning Environment - Makefile for Multi-platform Builds
# This Makefile provides targets for building binaries for different platforms

.PHONY: all clean windows linux darwin build-all help

# Default target
all: build-all

# Create bin directory
bin:
	mkdir bin

# Build for Windows (amd64)
windows: bin
	@echo "Building for Windows (amd64)..."
	set GOOS=windows&& set GOARCH=amd64&& go build -o bin/go-learn-windows-amd64.exe ./cmd
	@echo "[DONE] Windows binary: bin/go-learn-windows-amd64.exe"

# Build for Windows (386)
windows-386: bin
	@echo "Building for Windows (386)..."
	set GOOS=windows&& set GOARCH=386&& go build -o bin/go-learn-windows-386.exe ./cmd
	@echo "[DONE] Windows 386 binary: bin/go-learn-windows-386.exe"

# Build for Linux (amd64)
linux: bin
	@echo "Building for Linux (amd64)..."
	set GOOS=linux&& set GOARCH=amd64&& go build -o bin/go-learn-linux-amd64 ./cmd
	@echo "[DONE] Linux binary: bin/go-learn-linux-amd64"

# Build for Linux (386)
linux-386: bin
	@echo "Building for Linux (386)..."
	set GOOS=linux&& set GOARCH=386&& go build -o bin/go-learn-linux-386 ./cmd
	@echo "[DONE] Linux 386 binary: bin/go-learn-linux-386"

# Build for Linux (arm64)
linux-arm64: bin
	@echo "Building for Linux (arm64)..."
	set GOOS=linux&& set GOARCH=arm64&& go build -o bin/go-learn-linux-arm64 ./cmd
	@echo "[DONE] Linux ARM64 binary: bin/go-learn-linux-arm64"

# Build for macOS (Intel)
darwin-amd64: bin
	@echo "Building for macOS (Intel)..."
	set GOOS=darwin&& set GOARCH=amd64&& go build -o bin/go-learn-darwin-amd64 ./cmd
	@echo "[DONE] macOS Intel binary: bin/go-learn-darwin-amd64"

# Build for macOS (Apple Silicon)
darwin-arm64: bin
	@echo "Building for macOS (Apple Silicon)..."
	set GOOS=darwin&& set GOARCH=arm64&& go build -o bin/go-learn-darwin-arm64 ./cmd
	@echo "[DONE] macOS Apple Silicon binary: bin/go-learn-darwin-arm64"

# Build all platforms
build-all: windows windows-386 linux linux-386 linux-arm64 darwin-amd64 darwin-arm64
	@echo ""
	@echo "[WOOPWOOP] All platforms built successfully!"
	@echo "Binaries are located in the bin/ directory:"
	dir bin /b

# Clean build artifacts
clean:
	@echo "[CLEANING] Cleaning build artifacts..."
	if exist bin rmdir /s /q bin
	if exist releases rmdir /s /q releases
	@echo "[DONE] Clean complete"

# Build optimized binaries (smaller size)
optimized: bin
	@echo "Building optimized binaries..."
	set GOOS=linux&& set GOARCH=amd64&& go build -ldflags="-s -w" -o bin/go-learn-linux-amd64-optimized ./cmd
	set GOOS=windows&& set GOARCH=amd64&& go build -ldflags="-s -w" -o bin/go-learn-windows-amd64-optimized.exe ./cmd
	set GOOS=darwin&& set GOARCH=amd64&& go build -ldflags="-s -w" -o bin/go-learn-darwin-amd64-optimized ./cmd
	set GOOS=darwin&& set GOARCH=arm64&& go build -ldflags="-s -w" -o bin/go-learn-darwin-arm64-optimized ./cmd
	@echo "[DONE] Optimized binaries built"

# Build static binaries (Linux only)
static: bin
	@echo "Building static Linux binary..."
	set CGO_ENABLED=0&& set GOOS=linux&& set GOARCH=amd64&& go build -a -ldflags '-extldflags "-static"' -o bin/go-learn-linux-amd64-static ./cmd
	@echo "[DONE] Static binary: bin/go-learn-linux-amd64-static"

# Create release archives
release: build-all
	@echo "[DONE] Creating release archives..."
	if not exist releases mkdir releases
	for %%f in (bin\go-learn-*.exe) do ( \
		powershell "Compress-Archive -Path '%%f' -DestinationPath 'releases\%%~nf.zip' -Force" \
	)
	for %%f in (bin\go-learn-*) do ( \
		if not "%%~xf"==".exe" ( \
			powershell "Compress-Archive -Path '%%f' -DestinationPath 'releases\%%~nf.zip' -Force" \
		) \
	)
	@echo "[DONE] Release archives created in releases/"

# Show available targets
help:
	@echo "Go Learning Environment - Build Targets"
	@echo ""
	@echo "Available targets:"
	@echo "  all          - Build all platforms (default)"
	@echo "  windows      - Build Windows (amd64) binary"
	@echo "  windows-386  - Build Windows (386) binary"
	@echo "  linux        - Build Linux (amd64) binary"
	@echo "  linux-386    - Build Linux (386) binary"
	@echo "  linux-arm64  - Build Linux (arm64) binary"
	@echo "  darwin-amd64 - Build macOS (Intel) binary"
	@echo "  darwin-arm64 - Build macOS (Apple Silicon) binary"
	@echo "  build-all    - Build all platforms"
	@echo "  optimized    - Build optimized (smaller) binaries"
	@echo "  static       - Build static Linux binary"
	@echo "  release      - Build all and create release archives"
	@echo "  clean        - Remove build artifacts"
	@echo "  help         - Show this help message"
	@echo ""
	@echo "Examples:"
	@echo "  make linux           # Build Linux binary"
	@echo "  make release         # Build all and create releases"
	@echo "  make clean && make   # Clean and rebuild all"

# Development helpers
run:
	go run ./cmd

test:
	go test ./...

fmt:
	go fmt ./...

vet:
	go vet ./...

# Development target with all checks
dev: fmt vet test run
