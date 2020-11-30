package app

var (
	staticResourcesURL string
)

// Reload reloads the current page.
func Reload() {
	dispatch(func() {
		reload()
	})
}

// Run starts the wasm app and displays the UI node associated with the
// requested URL path.
//
// It panics if Go architecture is not wasm.
func Run(r RenderFunc) {
	run(r)
}

// Window returns the JavaScript "window" object.
func Window() BrowserWindow {
	return window
}
