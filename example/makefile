setup:
	brew install sqlc
	brew install dbmate
	go install github.com/mitranim/gow
	go install github.com/pyros2097/gromer/cmd/gromer
	docker pull postgres:14.1

start-db:
	docker run --name postgres141 -p 5432:5432 -e POSTGRES_PASSWORD=demo -d postgres:14.1

generate:
	sqlc generate -f db/sqlc.yaml
	gromer -pkg github.com/pyros2097/gromer/example

migrate: export DATABASE_URL=postgres://postgres:demo@127.0.0.1:5432/postgres?sslmode=disable
migrate:
	dbmate migrate

run: export DATABASE_URL=postgres://postgres:demo@127.0.0.1:5432/postgres?sslmode=disable
run:
	gow run main.go

build: export GOOS=linux
build: export GOARCH=amd64
build:
	go build -o main

build-image:
	cd .. && docker build -t example:develop .

run-image: export DATABASE_URL=postgres://postgres:demo@docker.for.mac.host.internal:5432/postgres?sslmode=disable
run-image:
	docker run -p 3000:3000 -e DATABASE_URL=$$DATABASE_URL example:develop

k8s-local:
	kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/controller-v1.0.0/deploy/static/provider/cloud/deploy.yaml
	kubectl apply -k k8s/development