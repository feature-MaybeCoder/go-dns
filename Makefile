build:
	@echo "Building binary"
	env GOOS=linux GOARCH=amd64 go build -o .bin/main ./cmd/server
	@echo "Starting docker build" 
	docker-compose up -d --build

up:
	@echo "Building binary"
	env GOOS=linux GOARCH=amd64 go build -o .bin/main ./cmd/server 
	@echo "Starting docker up"
	docker-compose up -d