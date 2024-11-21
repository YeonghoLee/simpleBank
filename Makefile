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

server :
	go run main.go

mock :
	mockgen -package mockdb -destination db/mock/store.go github.com/go_dev/simplebank/db/sqlc Store

.PHONY: createdb dropdb migrateup migratedown sqlc server mock