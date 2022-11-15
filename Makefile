run:
	go run main.go
	
run_db:
	docker start postgresNewsApp


run_migrate:
	@go run main.go -migrate=true


DB_URL=postgresql://root:secret@localhost:5435/newsApp?sslmode=disable

postgres:
	sudo docker run --name postgresNewsApp -p 5435:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

createdb:
	sudo docker exec -it postgresNewsApp createdb --username=root --owner=root newsApp

dropdb:
	sudo docker exec -it postgresNewsApp dropdb newsApp

migrateup:
	migrate -path migrations -database "$(DB_URL)" -verbose up

migratedown:
	migrate -path migrations -database "$(DB_URL)" -verbose down

docker_build:
	docker compose build --no-cache

docker_up:
	docker compose up

docker_down:
	docker compose down

.PHONY: run postgres createdb dropdb goose_reset goose_up goose_down goose_status docker_build docker_up docker_down