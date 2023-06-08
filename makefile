init-schema:
	migrate create -ext sql -dir database/migrations -seq init_schema

migrations-local:
	migrate -path database/migrations -database "mysql://root:secret@tcp(localhost:3306)/db_user?query" -verbose up

mocks:
	$(pwd) cd internal/handler  && mockery --all
	$(pwd) cd internal/usecase  && mockery --all
	$(pwd) cd internal/repository  && mockery --all

start:
	go run main.go