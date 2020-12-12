package app

import (
	"io"
	"unsafe"

	"github.com/pyros2097/wapp/js"
)

var (
	dispatch Dispatcher = Dispatch
	uiChan              = make(chan func(), 512)
)

type Context struct {
	// The UI element tied to the context.
	Src UI

	// The JavaScript value of the element tied to the context. This is a
	// shorthand for:
	//  ctx.Src.JSValue()
	JSSrc js.Value
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
