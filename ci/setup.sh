#!/usr/bin/env bash

echo "Installing Go dependencies..."
go mod download

if ! command -v golangci-lint &> /dev/null
then
  echo "golangci-lint could not be found. Installing..."
  go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
else
  echo "golangci-lint is already installed."
fi

echo "Setup complete!"