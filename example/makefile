setup:
	go install github.com/kyleconroy/sqlc/cmd/sqlc@v1.11.0
	go install github.com/amacneil/dbmate@v1.12.1
	go install github.com/mitranim/gow@latest
	go install github.com/pyros2097/gromer/cmd/gromer@v0.9.1
	docker pull postgres:14.1

start-db:
	docker run --name postgres141 -p 5432:5432 -e POSTGRES_PASSWORD=demo -d postgres:14.1

generate: export DATABASE_URL=postgres://postgres:demo@127.0.0.1:5432/postgres?sslmode=disable
generate:
	sqlc generate -f db/sqlc.yaml
	gromer -pkg github.com/pyros2097/gromer/example
	dbmate migrate

run: export PORT=3000
run: export DATABASE_URL=postgres://postgres:demo@127.0.0.1:5432/postgres?sslmode=disable
run:
	gow run main.go

build: export GOOS=linux
build: export GOARCH=amd64
build:
	go build -o main

docker-build-app:
	docker build -f ../example/Dockerfile  -t example-app:develop ../

docker-build-dbmate:
	docker build -f Dockerfile.dbmate  -t example-dbmate:develop .

docker-run-app: export DATABASE_URL=postgres://postgres:demo@docker.for.mac.host.internal:5432/postgres?sslmode=disable
docker-run-app:
	docker run -p 3000:3000 -e DATABASE_URL=$$DATABASE_URL example:latest

k8s-setup:
	kubectl apply -f https://raw.githubusercontent.com/reactive-tech/kubegres/v1.13/kubegres.yaml
	kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/controller-v1.0.0/deploy/static/provider/cloud/deploy.yaml

k8s-run:
	kubectl apply -k k8s/overlays/development

k8s-migrate:
	kubectl delete -f k8s/overlays/development/migration.yaml || true
	kubectl apply -f k8s/overlays/development/migration.yaml