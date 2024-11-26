run:
	go run cmd/app/main.go

build:
	go build -o bin/GymLens ./cmd/app

test:
	go test ./...
