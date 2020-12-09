run:
	go run *.go

go:
	go build

wasm:
	GOOS=js GOARCH=wasm go build -o assets/main.wasm

css:
	npx tailwindcss-cli@latest build assets/config.css -o assets/styles.css

lambda:
	sam local start-api

sync:
	aws s3 sync ./assets --delete public --dryrun s3://go-app-bucket-111