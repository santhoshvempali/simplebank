postgres:
	docker run --name postgres -p 5433:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=root -d postgres
createdb:
	docker exec -it postgres createdb --username=root --owner=root simple_bank
dropdb:
	docker exec -it postgres dropdb   --username=root --owner=root simple_bank
migrateup:
	migrate -path db/migrations -database "postgresql://root:root@localhost:5434/simple_bank?sslmode=disable" up
migratedown:
		migrate -path db/migrations -database "postgresql://root:root@localhost:5434/simple_bank?sslmode=disable" down
sqlc:
	sqlc generate
test:
	go test -v -cover ./...
.PHONY: postgres createdb dropdb  migrateup migratedown sqlc test