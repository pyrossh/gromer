package app

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Counter(c *RenderContext) UI {
	count, _ := c.UseInt(0)
	return Col(
		Row(Css("yellow"),
			Text("Counter"),
		),
		Row(
			Div(
				Text("-"),
			),
			Div(
				Text(count()),
			),
			Div(
				Text("+"),
			),
		),
	)
}

func Route(c *RenderContext) UI {
	return Div(
		Div(),
		Counter(c),
	)
}

func TestCreatePage(t *testing.T) {
	page := bytes.NewBuffer(nil)
	page.WriteString("<!DOCTYPE html>\n")
	Html(
		Head(
			Title("Title"),
		),
		Body(Route(NewRenderContext())),
	).Html(page)
	assert.Equal(t, "<!DOCTYPE html>\n<html>\n    <head>\n        <meta charset=\"UTF-8\">\n        <meta http-equiv=\"Content-Type\" content=\"text/html;charset=utf-8\">\n        <meta http-equiv=\"encoding\" content=\"utf-8\">\n        <title>\n            Title\n        </title>\n    </head>\n    <body>\n        <div>\n            <div></div>\n            <div class=\"flex flex-row justify-center align-items-center\">\n                <div class=\"yellow\">\n                    Counter\n                </div>\n                <div class=\"flex flex-row justify-center align-items-center\">\n                    <div>\n                        -\n                    </div>\n                    <div>\n                        0\n                    </div>\n                    <div>\n                        +\n                    </div>\n                </div>\n            </div>\n        </div>\n    </body>\n</html>", page.String())
}
