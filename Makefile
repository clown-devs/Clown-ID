.PHONY: build, run, all, check, swagger-gen

build: 
	go mod download && CGO_ENABLED=0 GOOS=linux go build -o ./bin/app ./cmd/main.go
	
run: check swagger-gen build
	docker-compose up --remove-orphans --build

check:
	go vet ./...

swagger-gen:
	swag init -g  ./cmd/main.go


.DEFAULT_GOAL := run