build:
	@go build -o bin/goBank

run: build
	@./bin/goBank

test:
	@go test -v ./...