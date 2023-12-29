# Default command
default:
    @just --list

run port:
    @echo "Starting prototype server..."
    go run service/cmd/main.go --port {{port}}

test:
    @echo "Running unit tests!"
    go test github.com/jgfranco17/algorithm-api/...
