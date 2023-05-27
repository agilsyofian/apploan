DB_URL=root:root@tcp(127.0.0.1:3306)/kreditplus?charset=utf8mb4&parseTime=True&loc=Local

runmysql:
	docker run --rm --name database -p 3306:3306 -e MYSQL_ROOT_PASSWORD=root -d mysql:5.7

createdb:
	docker exec -it database mysql -uroot -proot -e "CREATE DATABASE kreditplus;"

dropdb:
	docker exec -it database mysql -uroot -proot -e "DROP DATABASE kreditplus;"

migrate:
	go run migrations/migrate.go

migrateup:
	go run deploy/migration.go

test:
	go test -v -cover -short ./...

server:
	go run main.go

.PHONY: runmysql createdb dropdb migrate migrateup test server
