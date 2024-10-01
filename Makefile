launch:
	env GOOS=linux GOARCH=amd64 go build -o .bin/main ./cmd/server 
	docker-compose up -d 