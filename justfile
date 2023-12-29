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
    go test github.com/jgfranco17/algorithm-api/...
