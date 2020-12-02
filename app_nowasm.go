// +build !wasm

package app

func Reload() {
	panic("wasm required")
}

func Run(r RenderFunc) {
	panic("wasm required")
}
