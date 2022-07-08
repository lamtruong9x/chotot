mysql:
	docker run --name mysql8 -e MYSQL_ROOT_PASSWORD=pass -p 3030:3306 -d mysql:latest

migrate:
	migrate create -ext sql -dir ./internal/db/migration -seq init_schema

migrateup:
	migrate -path internal/db/migration -database "mysql://root:pass@tcp(localhost:3030)/chotot?multiStatements=true" -verbose up

migratedown:
	migrate -path internal/db/migration -database "mysql://root:pass@tcp(localhost:3030)/chotot?multiStatements=true" -verbose down

.PHONY: mysql migrate migrateup migratedown