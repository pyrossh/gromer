setup:
	go install github.com/mitranim/gow@latest
	go run github.com/playwright-community/playwright-go/cmd/playwright install --with-deps

dev:
	gow run main.go

test:
	go test -v ./...

build:
	podman build -f ../_example/Containerfile  -t example-app:develop ../

run:
	podman run -p 3000:3000 example-app:develop