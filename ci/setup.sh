#!/usr/bin/env bash

  info() {
  echo -e "[\e[93mINFO\e[0m]: $1"
}

install_go_tool() {
  local name=$1
  local github_path=$2

  if ! command -v "$name" &> /dev/null; then
    info "$name could not be found. Installing..."
    go install "$github_path"
  else
    info "$name is already installed."
  fi
}

info "Installing Go dependencies..."
go mod download

install_go_tool "golangci-lint" "github.com/golangci/golangci-lint/v2/cmd/golangci-lint@v2.0.2"

if [ -z "$CI" ]; then
  info "Running locally."
  install_go_tool "gow" "go install github.com/mitranim/gow@latest"
else
  info "Running in CI environment."
fi

info "Setup complete!"