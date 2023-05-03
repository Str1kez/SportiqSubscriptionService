.PHONY: build
.SILENT:

env:
	cp .env.example .env
build:
	go build -o .bin/subscription-service cmd/subscription-service/main.go
run: build
	.bin/subscription-service

upgrade:
	export $$(cat .env.dev); migrate -database "postgres://$$POSTGRES_USER:$$POSTGRES_PASSWORD@$$POSTGRES_HOST:$$POSTGRES_PORT/$$POSTGRES_DB" -path ./migrations up
downgrade:
	export $$(cat .env.dev); migrate -database "postgres://$$POSTGRES_USER:$$POSTGRES_PASSWORD@$$POSTGRES_HOST:$$POSTGRES_PORT/$$POSTGRES_DB" -path ./migrations down

db history:
	docker compose up -d --remove-orphans $@
build-image:
	docker build . -t subscription-service -f build/service/Dockerfile
up:
	docker compose up -d --remove-orphans
down:
	docker compose down
