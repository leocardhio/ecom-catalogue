container_name=ecom-pgdb

createdb:
	docker exec -it $(container_name) psql -U postgres -c "CREATE DATABASE ecommerce;"

dropdb:
	docker exec -it $(container_name) psql -U postgres -c "DROP DATABASE ecommerce;"

create-migration:
	migrate create -dir ./db/migration -ext sql $(name)

migrateup:
	 migrate -path ./db/migration -database postgres://postgres:password@localhost:5432/ecommerce?sslmode=disable -verbose up

migratedown:
	migrate -path ./db/migration -database postgres://postgres:password@localhost:5432/ecommerce?sslmode=disable -verbose down

.PHONY: createdb dropdb create-migration migrateup migratedown