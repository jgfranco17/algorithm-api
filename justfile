# Default command
default:
    @just --list

run:
    @echo "Starting prototype server..."
    go run service/cmd/main.go