build:
	@go build -o bin/transpogen .
run: build
	@./bin/transpogen
test:
	@go test -v ./...