# Database
# --- for Macos
#migrate:
#	migrate -database 'postgres://bookkeeper_user:bookkeeper_2023@localhost:5432/bookkeeper?sslmode=disable' -path db/migrations up
#migrate-rollback:
#	migrate -database 'postgres://bookkeeper_user:bookkeeper_2023@localhost:5432/bookkeeper?sslmode=disable' -path db/migrations down

# ---- for linux
# Database
create-migration:
	migrate create -ext sql -dir db/migrations -tz Local $(name)
migrate:
	migrate -database 'postgresql://bookkeeper_user:bookkeeper_2023@localhost:5432/bookkeeper?sslmode=disable' -path db/migrations up
migrate-rollback:
	migrate -database 'postgresql://bookkeeper_user:bookkeeper_2023@localhost:5432/bookkeeper?sslmode=disable' -path db/migrations down

dev:
	go run cmd/main/app.go
build:
	go build cmd/main/app.go
