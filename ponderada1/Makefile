.PHONY: swag tidy run

tidy:
	go mod tidy

swag: tidy
	swag init -g cmd/server/main.go -o docs

run:
	go run ./cmd/server
