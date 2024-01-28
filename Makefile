postgres:
	docker run --name postgres16 -p 5432:5432 -e POSTGRES_USERNAME=root -e POSTGRES_PASSWORD=secret -d postgres
dropdb:
	docker exec -it postgres16 dropdb simple_bank
createdb:
	docker exec -it postgres16 createdb --username=postgres --owner=postgres simple_bank
migrateup:
	migrate -path db/migration -database "postgresql://postgres:root@localhost:5432/simple_bank?sslmode=disable" -verbose up
migreatedown:
	migrate -path db/migration -database "postgresql://postgres:root@localhost:5432/simple_bank?sslmode=disable" -verbose down
sqlgen:
	docker run --rm -v ${pwd}:/src -w /src sqlc/sqlc generate
test:
	go test -v -cover ./...
.PHONY:postgres createdb dropdb migrateup migreatedown sqlgen test