make:
	go run main.go

test:
	go test ./... -short

build:
	go build -o gokedex main.go

install:
	go install