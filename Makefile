.PHONY: build, run, all, check

build: 
	go mod download && CGO_ENABLED=0 GOOS=linux go build -o ./bin/app ./cmd/main.go
	
run: check build
	docker-compose up --remove-orphans --build 

check:
	go vet ./...

	
.DEFAULT_GOAL := run