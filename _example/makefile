setup:
	go install github.com/mitranim/gow@latest
	go install github.com/pyros2097/gromer/cmd/gromer@latest

update:
	gromer -pkg github.com/pyros2097/gromer/_example

run:
	gow run main.go

build: export GOOS=linux
build: export GOARCH=amd64
build:
	go build -o main

docker-build:
	docker build -f ../_example/Dockerfile  -t example-app:develop ../

docker-run:
	docker run -p 3000:3000 example-app:develop