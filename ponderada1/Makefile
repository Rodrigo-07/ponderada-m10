.PHONY: swag tidy run db down dev

tidy:
	go mod tidy

swag: tidy
	swag init -g cmd/server/main.go -o docs

run:
	go run ./cmd/server

db:
	docker compose up -d

down:
	docker compose down

dev: db swag run
