container_name=catalog-db

createdb:
	docker exec -it $(container_name) psql -U postgres -c "CREATE DATABASE catalog;"

dropdb:
	docker exec -it $(container_name) psql -U postgres -c "DROP DATABASE catalog;"

create-migration:
	migrate create -dir ./db/migration -ext sql $(name)

#! Need Unix-based terminal to authenticate with docker
migrateup:
	migrate -path ./db/migration -database postgres://postgres:password@localhost:5432/catalog?sslmode=disable -verbose up 3

migratedown:
	migrate -path ./db/migration -database postgres://postgres:password@localhost:5432/catalog?sslmode=disable -verbose down 3

start:
	go run main.go -env=$(env)

.PHONY: createdb dropdb create-migration migrateup migratedown start