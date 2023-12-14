build:
	@go build -o bin/gohotel
run: build
	@export JWT_SECRET=1i1hlj1234l12nlqjlh123123l6h575jilj2l
	@./bin/gohotel
test:
	@go test -v ./...
seed:
	@go run scripts/seed.go

