build:
	@go build -o bin/gohotel
run: build
	@./bin/gohotel
test:
	@go test -v ./...
seed:
	@go run scripts/seed.go

