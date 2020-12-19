run:
	go run *.go

wasm: export GOOS=js
wasm: export GOARCH=wasm
wasm:
	go build -o assets/main.wasm

css:
	npx tailwindcss-cli@latest build assets/config.css -o assets/styles.css

deploy: export NODE_ENV=production
deploy:
	npx tailwindcss-cli@latest build assets/config.css -o assets/styles.css
	go build -o main
	make wasm
	sam deploy
	aws s3 sync ./assets s3://wapp.pyros2097.dev/assets --delete
	aws cloudfront create-invalidation --distribution-id E17XVZYYZ1JXEU --paths "/*"

local:
	sam local start-api
