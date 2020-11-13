package app

import (
	"strings"
)

var (
	staticResourcesURL string
)

// KeepBodyClean prevents third-party Javascript libraries to add nodes to the
// body element.
func KeepBodyClean() (close func()) {
	return keepBodyClean()
}

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

// StaticResource makes a static resource path point to the right
// location whether the root directory is remote or not.
//
// Static resources are resources located in the web directory.
//
// This call is used internally to resolve paths within Cite, Data, Href, Src,
// and SrcSet. Paths already resolved are skipped.
func StaticResource(path string) string {
	if !strings.HasPrefix(path, "/web/") &&
		!strings.HasPrefix(path, "web/") {
		return path
	}

	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}

	return staticResourcesURL + path
}

// Window returns the JavaScript "window" object.
func Window() BrowserWindow {
	return window
}
