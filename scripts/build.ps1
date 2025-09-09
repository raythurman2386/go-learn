# Go Learning Environment - Multi-platform Build Script (Windows)
# This script builds binaries for Windows, Linux, and macOS

param(
    [switch]$Clean,
    [switch]$Release
)

# Colors for output (PowerShell)
$Green = "Green"
$Yellow = "Yellow"
$Red = "Red"
$Blue = "Blue"
$Cyan = "Cyan"

# Create bin directory
if (!(Test-Path bin)) {
    New-Item -ItemType Directory -Path bin | Out-Null
}

# Clean build directory if requested
if ($Clean) {
    Write-Host "[CLEAN] Cleaning bin directory..." -ForegroundColor $Yellow
    Remove-Item bin/* -Force -ErrorAction SilentlyContinue
}

Write-Host "[BUILD] Building Go Learning Environment for multiple platforms..." -ForegroundColor $Blue
Write-Host ""

# Function to build for a specific platform
function New-PlatformBinary {
    param(
        [string]$OS,
        [string]$Arch,
        [string]$OutputName,
        [string]$Description
    )

    Write-Host "Building for $Description..." -ForegroundColor $Yellow

    $env:GOOS = $OS
    $env:GOARCH = $Arch

    try {
        go build -o bin/$OutputName ./cmd
        Write-Host "[OK] Successfully built: bin/$OutputName" -ForegroundColor $Green
        return $true
    }
    catch {
        Write-Host "[ERROR] Failed to build for $Description" -ForegroundColor $Red
        Write-Host "   Error: $($_.Exception.Message)" -ForegroundColor $Red
        return $false
    }
}

# Build for all platforms
Write-Host "Building binaries..." -ForegroundColor $Blue
$buildCount = 0
$successCount = 0

# Windows builds
if (New-PlatformBinary "windows" "amd64" "go-learn-windows-amd64.exe" "Windows (amd64)") { $successCount++ }; $buildCount++
if (New-PlatformBinary "windows" "386" "go-learn-windows-386.exe" "Windows (386)") { $successCount++ }; $buildCount++

# Linux builds
if (New-PlatformBinary "linux" "amd64" "go-learn-linux-amd64.exe" "Linux (amd64)") { $successCount++ }; $buildCount++
if (New-PlatformBinary "linux" "386" "go-learn-linux-386.exe" "Linux (386)") { $successCount++ }; $buildCount++
if (New-PlatformBinary "linux" "arm64" "go-learn-linux-arm64.exe" "Linux (arm64)") { $successCount++ }; $buildCount++

# macOS builds
if (New-PlatformBinary "darwin" "amd64" "go-learn-darwin-amd64.exe" "macOS (Intel)") { $successCount++ }; $buildCount++
if (New-PlatformBinary "darwin" "arm64" "go-learn-darwin-arm64.exe" "macOS (Apple Silicon)") { $successCount++ }; $buildCount++

Write-Host ""
Write-Host "[DONE] Build complete!" -ForegroundColor $Green
Write-Host "Results: $successCount/$buildCount builds successful" -ForegroundColor $(if ($successCount -eq $buildCount) { $Green } else { $Yellow })
Write-Host ""
Write-Host "Binaries are located in the bin/ directory:" -ForegroundColor $Blue
Write-Host ""

# List the built binaries
Get-ChildItem bin/ | ForEach-Object {
    Write-Host "  $($_.Name)" -ForegroundColor $Cyan
}

Write-Host ""
Write-Host "[USAGE] Usage:" -ForegroundColor $Yellow
Write-Host "  Windows: bin\go-learn-windows-amd64.exe" -ForegroundColor $Cyan
Write-Host "  Linux:   bin\go-learn-linux-amd64.exe" -ForegroundColor $Cyan
Write-Host "  macOS:   bin\go-learn-darwin-amd64.exe" -ForegroundColor $Cyan
Write-Host ""

if ($Release) {
    Write-Host "[PACKAGE] Creating release archives..." -ForegroundColor $Blue

    # Create releases directory
    if (!(Test-Path releases)) {
        New-Item -ItemType Directory -Path releases | Out-Null
    }

    # Create zip files for each platform
    $platforms = @(
        @{Name="windows-amd64"; Ext=".exe"},
        @{Name="windows-386"; Ext=".exe"},
        @{Name="linux-amd64"; Ext=".exe"},
        @{Name="linux-386"; Ext=".exe"},
        @{Name="linux-arm64"; Ext=".exe"},
        @{Name="darwin-amd64"; Ext=".exe"},
        @{Name="darwin-arm64"; Ext=".exe"}
    )

    foreach ($platform in $platforms) {
        $fileName = "go-learn-$($platform.Name)$($platform.Ext)"
        $zipName = "go-learn-$($platform.Name).zip"

        if (Test-Path "bin/$fileName") {
            try {
                Compress-Archive -Path "bin/$fileName" -DestinationPath "releases/$zipName" -Force
                Write-Host "  [OK] Created: releases/$zipName" -ForegroundColor $Green
            }
            catch {
                Write-Host "  [ERROR] Failed to create: releases/$zipName" -ForegroundColor $Red
            }
        }
    }

    Write-Host ""
    Write-Host "[DIR] Release archives created in releases/ directory" -ForegroundColor $Blue
}

Write-Host ""
Write-Host "For more information, see the README.md file" -ForegroundColor $Blue
