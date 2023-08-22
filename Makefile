.PHONY: build, run, all, check, swagger-gen

build: 
	go mod download && CGO_ENABLED=0 GOOS=linux go build -o ./bin/app ./cmd/main.go

start:
	docker-compose up --remove-orphans

check:
	go vet ./...

swagger-gen:
	swag init -g  ./cmd/main.go

run: check build start

run-with-swag: run swagger-gen
.DEFAULT_GOAL := run