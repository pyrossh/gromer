setup:
	go install github.com/kyleconroy/sqlc/cmd/sqlc@v1.13.0
	go install github.com/amacneil/dbmate@v1.12.1
	go install github.com/mitranim/gow@latest
	go install github.com/pyros2097/gromer/cmd/gromer@latest
	docker pull postgres:14.1
	docker run --name postgres141 -p 5432:5432 -e POSTGRES_PASSWORD=demo -d postgres:14.1

generate: export DATABASE_URL=postgres://postgres:demo@127.0.0.1:5432/postgres?sslmode=disable
generate:
	sqlc generate -f db/sqlc.yaml
	gromer -pkg github.com/pyros2097/gromer/_example
	dbmate migrate

run: export PORT=3000
run: export DATABASE_URL=postgres://postgres:demo@127.0.0.1:5432/postgres?sslmode=disable
run:
	gow run main.go

build: export GOOS=linux
build: export GOARCH=amd64
build:
	go build -o main

docker-build:
	docker build -f ../_example/Dockerfile  -t example-app:develop ../

docker-run: export DATABASE_URL=postgres://postgres:demo@docker.for.mac.host.internal:5432/postgres?sslmode=disable
docker-run:
	docker run -p 3000:3000 -e DATABASE_URL=$$DATABASE_URL example:latest