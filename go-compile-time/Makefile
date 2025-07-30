.PHONY: build run clean help

# Default target
help:
	@echo "Available targets:"
	@echo "  build - Build with instrumentation"
	@echo "  run   - Run instrumented version"
	@echo "  clean - Remove build artifacts"

# Build with instrumentation
build:
	@echo "Building with integrated instrumentation (OpenTelemetry style)..."
	go build -o build-tool build-tool.go
	./build-tool build -o fibonacci app/fibonacci.go
	@echo "Done! Run with: ./fibonacci"

# Run the instrumented version
run: build
	./fibonacci

# Clean build artifacts
clean:
	rm -f fibonacci
	rm -f build-tool 
