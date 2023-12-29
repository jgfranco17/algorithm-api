# Default command
default:
    @just --list

run:
    @echo "Starting prototype server..."
    go run service/cmd/main.go

test:
    @echo "Running unit tests!"
    go test github.com/jgfranco17/algorithm-api/...
