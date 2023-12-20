.PHONY: compile
compile:
	go build -o stori cmd/cli/cli.go

.PHONY: install
install:
	go install

.PHONY: coverage
coverage:
	go test -coverprofile=coverage.out ./internal/calculator/infrastructure
	go tool cover -html=coverage.out
	go tool cover -func=coverage.out
