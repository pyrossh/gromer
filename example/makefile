run:
	go run *.go

local:
	sam local start-api

css:
	npx tailwindcss-cli@latest build -i assets/config.css -o assets/styles.css

build: export GOOS=linux
build: export GOARCH=amd64
build:
	go build -o main

deploy: export NODE_ENV=production
deploy:
	make css
	make build
	sam deploy
	aws s3 sync ./assets s3://wapp-example/assets --delete
	aws cloudfront create-invalidation --distribution-id E3C93YR5O58WG6 --paths "/*"
