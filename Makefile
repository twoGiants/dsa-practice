CMD=go
BUILD=$(CMD) build
CLEAN=$(CMD) clean
TEST=$(CMD) test
BINARY_NAME=dsa
BUILD_DIR=build
BINARY_PATH=$(BUILD_DIR)/$(BINARY_NAME)
BINARY_UNIX=$(BINARY_PATH)_unix

.PHONY: all setup build run test watch clean build-linux lint

all: setup lint test build

setup:
	./ci/setup.sh

build:
	@mkdir -p $(BUILD_DIR)
	$(BUILD) -o $(BINARY_PATH) -v

run:
	$(CMD) run .

test:
	$(TEST) -v ./...

watch:
	gow test ./...

clean:
	$(CLEAN)
	rm -rf $(BUILD_DIR)

build-linux:
	@mkdir -p $(BUILD_DIR)
	GOOS=linux GOARCH=amd64 $(BUILD) -o $(BINARY_UNIX) -v

lint:
	golangci-lint run