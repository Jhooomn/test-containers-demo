#!/bin/bash

# Exit script if any command fails
set -e

# Clean up the Go modules and remove any unused modules.
echo "Tidying up the Go modules..."
go mod tidy

# Download all the dependencies required for your Go project.
echo "Getting all the dependencies..."
go get ./...

# Copy all dependencies to the vendor directory.
echo "Vendoring all dependencies..."
go mod vendor

echo "Dependencies have been updated and vendored successfully."