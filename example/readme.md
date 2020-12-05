# wapp-example

An example to demostrate the wapp framework

1. Compile wasm frontend,
`GOOS=js GOARCH=wasm go build -o assets/main.wasm`

2. Run the server,
`go run *.go`

These commands are also available in the makefile for convenience