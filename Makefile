createdb :
	docker exec -it postgres12 createdb --username=root --owner=root simple_bank

dropdb :
	docker exec -it postgres12 dropdb simple_bank

migrateup :
	migrate -path db/migration -database "postgresql://root:secret@localhost:1234/simple_bank?sslmode=disable" -verbose up

migratedown :
	migrate -path db/migration -database "postgresql://root:secret@localhost:1234/simple_bank?sslmode=disable" -verbose down

sqlc :
	sqlc generate

.PHONY: createdb dropdb migrateup migratedown sqlc