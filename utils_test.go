package app

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func stubRoute(c *RenderContext) UI {
	return Div(Text("Stub"))
}

var testRoutes = map[string]RenderFunc{
	"/about":     stubRoute,
	"/clock":     stubRoute,
	"/container": stubRoute,
	"/":          stubRoute,
}

func TestMatchPath(t *testing.T) {
	assert.Equal(t, true, matchPath("/", "/about"))
}
