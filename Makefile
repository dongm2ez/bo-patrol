.PHONY: build run test clean

build:
	go build -o bin/server cmd/server/main.go

run:
	go run cmd/server/main.go

test:
	go test -v ./...

clean:
	rm -rf bin/

migrate:
	go run cmd/migrate/main.go

docker-up:
	docker-compose up -d

docker-down:
	docker-compose down

build-web:
	cd web && npm run build

dev-web:
	cd web && npm run dev

.PHONY: build run test clean migrate docker-up docker-down build-web dev-web
