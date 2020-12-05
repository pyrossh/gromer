run:
	go run *.go

build:
	go build

build-wasm:
	GOOS=js GOARCH=wasm go build -o assets/main.wasm

build-css:
	npx tailwindcss-cli@latest build assets/config.css -o assets/styles.css