# Makefile for GAuth CLI Application

# Go-related variables
APP_NAME = gauth
MAIN = main.go
BIN_DIR = ./bin
BIN = $(BIN_DIR)/$(APP_NAME)

# Default target
all: build

# Build the binary
build:
	@echo "Building $(APP_NAME)..."
	@mkdir -p $(BIN_DIR)
	@go build -o $(BIN) $(SRC)
	@echo "Build complete: $(BIN)"

# Install the binary to /usr/local/bin
install:
	@echo "Installing $(APP_NAME)..."
	@cp $(BIN) /usr/local/bin/
	@echo "Installed $(APP_NAME) to /usr/local/bin"

# Clean up build artifacts
clean:
	@echo "Cleaning up..."
	@rm -rf $(BIN_DIR)
	@echo "Clean complete."

# Uninstall the binary from /usr/local/bin
uninstall:
	@echo "Uninstalling $(APP_NAME)..."
	@rm -f /usr/local/bin/$(APP_NAME)
	@echo "Uninstall complete."

# Help information
help:
	@echo "Makefile for $(APP_NAME)"
	@echo ""
	@echo "Usage:"
	@echo "  make [target]"
	@echo ""
	@echo "Targets:"
	@echo "  build      Build the application binary"
	@echo "  install    Install the application to /usr/local/bin"
	@echo "  clean      Clean up build artifacts"
	@echo "  uninstall  Uninstall the application from /usr/local/bin"
	@echo "  help       Show this help message"
