binary_name := "kei"
main_package := "./main.go"
coverage_file := "coverage.out"
coverage_html := "coverage.html"

# Default recipe to display help
default:
    @just --list

# Run the application
run:
    go run {{main_package}}

# Run with hot reload using gow
dev:
    @echo "Running with hot reload..."
    @command -v gow >/dev/null 2>&1 || { echo "❌ gow not installed. Run 'just install-tools'"; exit 1; }
    gow run {{main_package}}

# Build the application
build:
    @echo "Building {{binary_name}}..."
    go build -o bin/{{binary_name}} {{main_package}}
    @echo "Build complete: bin/{{binary_name}}"

# Install the application to $GOPATH/bin
install:
    @echo "Installing {{binary_name}}..."
    go install {{binary_name}}
    @echo "Install complete: {{binary_name}} installed to $(go env GOPATH)/bin"

# Run tests
test:
    @echo "Running tests..."
    go test -v -race ./...

# Run tests with coverage report
test-coverage:
    @echo "Running tests with coverage..."
    go test -v -race -coverprofile={{coverage_file}} -covermode=atomic ./...
    go tool cover -html={{coverage_file}} -o {{coverage_html}}
    @echo "Coverage report generated: {{coverage_html}}"

# Display test coverage in terminal
coverage:
    @echo "Running tests and displaying coverage..."
    go test -v -race -coverprofile={{coverage_file}} -covermode=atomic ./...
    go tool cover -func={{coverage_file}}

# Format code with gofumpt and organize imports
fmt:
    @echo "Formatting code with gofumpt..."
    @command -v gofumpt >/dev/null 2>&1 || { echo "❌ gofumpt not installed. Refer flake.nix for installation"; exit 1; }
    gofumpt -l -w .
    @echo "Organizing imports with goimports..."
    @command -v goimports-reviser >/dev/null 2>&1 || { echo "❌ goimports-reviser not installed. Refer flake.nix for installation"; exit 1; }
    goimports-reviser -rm-unused ./...
    @echo "Running golangci-lint..."
    @command -v golangci-lint >/dev/null 2>&1 || { echo "❌ golangci-lint not installed. Refer flake.nix for installation"; exit 1; }
    golangci-lint fmt ./...
    @echo "✓ Formatting complete"

# Run golangci-lint
lint:
    @echo "Running golangci-lint..."
    @command -v golangci-lint >/dev/null 2>&1 || { echo "❌ golangci-lint not installed. Refer flake.nix for installation"; exit 1; }
    golangci-lint run --timeout 5m ./...

# Run golangci-lint with auto-fix
lint-fix:
    @echo "Running golangci-lint with auto-fix..."
    @command -v golangci-lint >/dev/null 2>&1 || { echo "❌ golangci-lint not installed. Refer flake.nix for installation"; exit 1; }
    golangci-lint run --fix --timeout 5m ./...

# Run go vet
vet:
    @echo "Running go vet..."
    go vet ./...

# Run go mod tidy
tidy:
    @echo "Tidying go modules..."
    go mod tidy
    go mod verify

# Run go generate
generate:
    @echo "Running go generate..."
    go generate ./...

# Run security checks with govulncheck
security:
    @echo "Running security checks..."
    @command -v govulncheck >/dev/null 2>&1 || { echo "❌ govulncheck not installed. Refer flake.nix for installation"; exit 1; }
    govulncheck ./...

# Run all quality checks
check: fmt vet lint test
    @echo "✓ All checks passed!"

# Run all checks and generate coverage
check-all: fmt vet lint test-coverage security
    @echo "✓ All checks and coverage complete!"

# Clean build artifacts, cache, and modules
clean:
    @echo "Cleaning..."
    rm -rf bin/
    rm -f {{coverage_file}} {{coverage_html}}
    go clean -cache -testcache
    go clean -modcache
    @echo "✓ Clean complete"

# Download dependencies
deps:
    @echo "Downloading dependencies..."
    go mod download

# Update dependencies
deps-up:
    @echo "Updating dependencies..."
    go get -u ./...
    go mod tidy
