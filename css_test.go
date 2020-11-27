package app

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreatePage(t *testing.T) {
	page := bytes.NewBuffer(nil)
	page.WriteString("<!DOCTYPE html>\n")
	Html(
		Body(Div(Css("abc"), Text("About Me"))),
	).Html(page)
	assert.Equal(t, "<!DOCTYPE html>\n<html>\n    <head>\n        <div class=\"abc\"></div>\n    </head>\n</html>", page.String())
}
