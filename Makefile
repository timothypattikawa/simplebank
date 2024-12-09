postgres17:
	docker run -d --name postgres17 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -v ~/docker/postgres-data:/var/lib/postgresql/data -p 5432:5432 postgres:17-alpine
createdb:
	docker exec -it postgres17 createdb --username=root --owner=root simple_bank
dropdb:
	docker exec -it postgres17 dropdb --username=root --owner=root simple_bank
migrateup:
	migrate -path script/migrations -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up
migratedown:
	migrate -path script/migrations -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down
sqlc:
	sqlc ganarate 
test:
	go test -v -cover ./...

.PHONY: postgres17 createdb dropdb migrateup migratedown sqlc