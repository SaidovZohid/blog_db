-include .env
.SILENT:
DB_URL=postgresql://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@$(POSTGRES_HOST):$(POSTGRES_PORT)/$(POSTGRES_DATABASE)?sslmode=disable

run:
	go run cmd/main.go
	
print:
	echo $(DB_URL)

swag-init:
	swag init -g api/api.go -o api/docs

composeup:
	docker compose --env-file ./.env.docker up

migrateup:
	migrate -path migrations -database "$(DB_URL)" -verbose up

migratedown:
	migrate -path migrations -database "$(DB_URL)" -verbose down