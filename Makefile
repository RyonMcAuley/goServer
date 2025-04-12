# Set variables for easy reuse
BINARY_NAME=goServer
MAIN_GO=cmd/app/main.go

# Default target, build the binary
all: build

# Build the Go binary
build:
	go build -o $(BINARY_NAME) $(MAIN_GO)

# Run the Go binary
run: build
	./$(BINARY_NAME)

# Clean the Go binary
clean:
	rm -f $(BINARY_NAME)

