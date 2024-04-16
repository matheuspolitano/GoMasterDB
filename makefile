dbname = fintech_db
containerName = postgres-example
POSTGRESQL_URL = "postgresql://root:MP@TEST123@localhost:5432/fintech_db?sslmode=disable"

postgres:
	docker run --name $(containerName)  -p 5432:5432 -e POSTGRES_PASSWORD=MP@TEST123 -e POSTGRES_USER=root -d  postgres:12.0-alpine

createdb:
	docker exec -it $(containerName) createdb --username=root --owner=root $(dbname)

dropdb:
	docker exec -it $(containerName)  dropdb $(dbname)

migrationup:
	migrate -database $(POSTGRESQL_URL) -path db/migrations up

migrationdown:
	migrate -database $(POSTGRESQL_URL) -path db/migrations down

sqlc:
	sqlc generate

test:
	go test -v ./...	
