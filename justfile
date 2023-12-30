# Default command
default:
    @just --list

run-debug port:
    @echo "Starting prototype server..."
    go run service/cmd/main.go --port={{port}} --dev=true

run-prod port:
    @echo "Starting production server..."
    go run service/cmd/main.go --port {{port}} --dev=false

test:
    @echo "Running unit tests!"
    go clean -testcache
    go test -cover github.com/jgfranco17/algorithm-api/...

# Build Docker image
build:
	@echo "Building Docker image..."
	docker build -t algorithm-api:latest -f ./Dockerfile .
	@echo "Docker image built successfully!"
