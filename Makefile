.PHONY: build
.SILENT:

env:
	cp .example.env .env
build-server:
	go build -o .bin/subscription-server cmd/subscription-server/main.go
run-server: build-server
	.bin/subscription-server
build-handler:
	go build -o .bin/subscription-message-handler cmd/subscription-message-handler/main.go
run-handler: build-handler
	.bin/subscription-message-handler

upgrade:
	export $$(cat .dev.env); migrate -database "postgres://$$POSTGRES_USER:$$POSTGRES_PASSWORD@$$POSTGRES_HOST:$$POSTGRES_PORT/$$POSTGRES_DB" -path ./migrations up
downgrade:
	export $$(cat .dev.env); migrate -database "postgres://$$POSTGRES_USER:$$POSTGRES_PASSWORD@$$POSTGRES_HOST:$$POSTGRES_PORT/$$POSTGRES_DB" -path ./migrations down

db history:
	docker compose up -d --remove-orphans $@
build-image-server:
	docker build . -t subscription-server -f build/server/Dockerfile
build-image-handler:
	docker build . -t subscription-message-handler -f build/handler/Dockerfile
up:
	docker compose up -d --remove-orphans
down:
	docker compose down
