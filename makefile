TARGET = "myfile.txt"
TEXT = "matheus"
DBNAME = fintech_db
CONTAINER_NAME = postgres-example
IMAGE = postgres:12.0-alpine
PORTS = 5432:5432
ENV_VARS = POSTGRES_PASSWORD=MP@TEST123 -e POSTGRES_USER=root
POSTGRESQL_URL = "postgresql://root:MP@TEST123@localhost:5432/fintech_db?sslmode=disable"

.PHONY: postgres




postgres:
	@var1=$$(docker ps -aq -f name=^$(CONTAINER_NAME)$$);\
	if [ "$$var1" ]; then \
		echo "Container $(CONTAINER_NAME) exist"; \
		if [ $$(docker ps -q -f name=^$(CONTAINER_NAME)$$) ]; then \
			echo "Docker $(CONTAINER_NAME) is running"; \
		fi \
	else \
		docker run --name $(CONTAINER_NAME) -p $(PORTS) -e $(ENV_VARS) -d $(IMAGE); \
	fi

createdb:
	docker exec -it $(CONTAINER_NAME) createdb --username=root --owner=root $(DBNAME)

dropdb:
	docker exec -it $(CONTAINER_NAME)  dropdb $(DBNAME)

migrationup:
	migrate -database $(POSTGRESQL_URL) -path db/migrations up

migrationdown:
	migrate -database $(POSTGRESQL_URL) -path db/migrations down

sqlc:
	sqlc generate

test:
	go test -v ./...	

