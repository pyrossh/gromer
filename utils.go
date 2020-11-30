package app

import (
	"context"
	"io"
	"unsafe"
)

var (
	dispatch Dispatcher = Dispatch
	uiChan              = make(chan func(), 512)
)

// Context represents a context that is tied to a UI element. It is canceled
// when the element is dismounted.
//
// It implements the context.Context interface.
//  https://golang.org/pkg/context/#Context
type Context struct {
	context.Context

	// The UI element tied to the context.
	Src UI

	// The JavaScript value of the element tied to the context. This is a
	// shorthand for:
	//  ctx.Src.JSValue()
	JSSrc Value
}

// Dispatcher is a function that executes the given function on the goroutine
// dedicated to UI.
type Dispatcher func(func())

// Dispatch executes the given function on the UI goroutine.
func Dispatch(f func()) {
	uiChan <- f
}

func writeIndent(w io.Writer, indent int) {
	for i := 0; i < indent*4; i++ {
		w.Write(stob(" "))
	}
}

func ln() []byte {
	return stob("\n")
}

func btos(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func stob(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(&s))
}
