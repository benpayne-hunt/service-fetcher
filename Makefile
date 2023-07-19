build:
	@go build -o bin/service-fetcher

run:
	@go run main.go

test:
	@go test -v ./...
