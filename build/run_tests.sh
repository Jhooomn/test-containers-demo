#!/bin/bash

# Set the script to exit immediately on error
set -e

# Navigate to the build folder from the project root
cd "$(dirname "$0")"

# Go up to the project root where the go.mod file is located
cd ..

# Run the go test command with the -v (verbose) flag on all packages
echo "Running Go tests..."
go test ./... -v

echo "Tests completed."