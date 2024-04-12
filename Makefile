.DEFAULT_GOAL := build

BINARY_NAME := flac-finder
BUILD_DIR := build

dependencies:
	asdf install
	go mod tidy

build_x64:
	GOOS=darwin go build -o $(BUILD_DIR)/$(BINARY_NAME)_mac_x64 .
	GOOS=linux GOARCH=amd64 go build -o $(BUILD_DIR)/$(BINARY_NAME)_linux_x64 .
	GOOS=windows go build -o $(BUILD_DIR)/$(BINARY_NAME)_win_x64.exe .

build_arm:
	GOOS=linux GOARCH=arm GOARM=6 go build -o $(BUILD_DIR)/$(BINARY_NAME)_linux_armv6 .
	GOOS=linux GOARCH=arm GOARM=7 go build -o $(BUILD_DIR)/$(BINARY_NAME)_linux_armv7 .
	GOOS=darwin GOARCH=arm64 go build -o $(BUILD_DIR)/$(BINARY_NAME)_mac_arm64 .

build: build_x64 build_arm
clean:
	@echo "Cleaning build directory"
	@rm -rf $(BUILD_DIR)
