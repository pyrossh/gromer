// +build !wasm

package app

import (
	"net/url"
	"os"
)

func getenv(k string) string {
	return os.Getenv(k)
}

func navigate(u *url.URL, updateHistory bool) error {
	panic("wasm required")
}

func reload() {
	panic("wasm required")
}

func run(r RenderFunc) {
	panic("wasm required")
}
