#!/bin/bash

# Go Learning Environment - Multi-platform Build Script
# This script builds binaries for Windows, Linux, and macOS

set -e  # Exit on any error

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Create bin directory
mkdir -p bin

echo -e "${BLUE}🚀 Building Go Learning Environment for multiple platforms...${NC}"
echo

# Function to build for a specific platform
build_platform() {
    local os=$1
    local arch=$2
    local output_name=$3
    local description=$4

    echo -e "${YELLOW}Building for ${description}...${NC}"

    if GOOS=$os GOARCH=$arch go build -o bin/$output_name ./cmd; then
        echo -e "${GREEN}✅ Successfully built: bin/$output_name${NC}"
    else
        echo -e "${RED}❌ Failed to build for ${description}${NC}"
        return 1
    fi
}

# Build for all platforms
echo -e "${BLUE}Building binaries...${NC}"

# Windows
build_platform "windows" "amd64" "go-learn-windows-amd64.exe" "Windows (amd64)"
build_platform "windows" "386" "go-learn-windows-386.exe" "Windows (386)"

# Linux
build_platform "linux" "amd64" "go-learn-linux-amd64" "Linux (amd64)"
build_platform "linux" "386" "go-learn-linux-386" "Linux (386)"
build_platform "linux" "arm64" "go-learn-linux-arm64" "Linux (arm64)"

# macOS
build_platform "darwin" "amd64" "go-learn-darwin-amd64" "macOS (Intel)"
build_platform "darwin" "arm64" "go-learn-darwin-arm64" "macOS (Apple Silicon)"

echo
echo -e "${GREEN}🎉 Build complete!${NC}"
echo -e "${BLUE}Binaries are located in the bin/ directory:${NC}"
echo

# List the built binaries
if command -v ls &> /dev/null; then
    ls -la bin/
else
    echo "Contents of bin/ directory:"
    for file in bin/*; do
        if [ -f "$file" ]; then
            echo "  $file"
        fi
    done
fi

echo
echo -e "${YELLOW}💡 Usage:${NC}"
echo "  Windows: bin/go-learn-windows-amd64.exe"
echo "  Linux:   ./bin/go-learn-linux-amd64"
echo "  macOS:   ./bin/go-learn-darwin-amd64"
echo
echo -e "${BLUE}📦 To create distribution archives:${NC}"
echo "  tar -czf releases/go-learn-linux-amd64.tar.gz bin/go-learn-linux-amd64"
echo "  zip releases/go-learn-windows-amd64.zip bin/go-learn-windows-amd64.exe"
