setup:
	go install github.com/mitranim/gow@latest
	go install github.com/pyros2097/gromer/cmd/gromer@latest

update:
	gromer -pkg github.com/pyros2097/gromer/_example

dev:
	gow run main.go

build:
	podman build -f ../_example/Containerfile  -t example-app:develop ../

run:
	podman run -p 3000:3000 example-app:develop