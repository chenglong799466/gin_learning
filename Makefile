.PHONY: run build docker

run:
	go run cmd/main.go

build:
	go build -o bin/server cmd/main.go

docker:
	docker build -t gin_learning .
