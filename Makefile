DB_URL=root:root@tcp(127.0.0.1:3306)/kreditplus?charset=utf8mb4&parseTime=True&loc=Local

runmysql:
	docker run --rm --name dbmysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=root -d mysql:5.7

createdb:
	docker exec -it dbmysql mysql -uroot -proot -e "CREATE DATABASE kreditplus;"

dropdb:
	docker exec -it dbmysql mysql -uroot -proot -e "DROP DATABASE kreditplus;"

migrate:
	go run migrations/migrate.go

sqlc:
	sqlc generate

test:
	go test -v -cover -short ./...

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/agilsyofian/minibank/db/sqlc Store

.PHONY: runmysql createdb dropdb migrate test server mock
