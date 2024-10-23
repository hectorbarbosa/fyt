# Create local database
.PHONY: createdb 
createdb:
	psql -h localhost -U postgres \
        -c "CREATE DATABASE fyt ENCODING='UTF8'";

# Drop local database
.PHONY: dropdb 
dropdb:
	psql -h localhost -U postgres \
        -c "DROP DATABASE IF EXISTS fyt";

# Create Docker container `fyt` tables
migrateup:
	migrate -path db/migrations/ -database postgres://postgres:postgres@localhost:5432/fyt?sslmode=disable up

# Drop Docker container `fyt` tables
migratedown:
	migrate -path db/migrations/ -database postgres://postgres:postgres@localhost:5432/fyt?sslmode=disable down 

.PHONY: sqlc 
sqlc:
	internal/storage/postgresql/sqlc generate 

.PHONY: build
build:
	go build -o bin/fyt -v ./cmd/fyt

.PHONY: run 
run:
	bin/fyt -env dev 

.PHONY: test
test:
	go test -v -race -timeout 30s ./...

.DEFAULT_GOAL := build