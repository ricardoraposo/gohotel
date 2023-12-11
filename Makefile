build:
	@go build -o bin/gohotel
run: build
	@./bin/gohotel
test:
	@go test ./...
