package about

import (
	"context"

	. "github.com/pyros2097/gromer"
)

func GET(c context.Context) (HtmlContent, int, error) {
	return Html(`
		{{#Page "gromer example"}}
			<div class="flex flex-col justify-center items-center">
					{{#Header "123"}}
						A new link is here
					{{/Header}}
					<h1>About Me</h1>
			</div>
		{{/Page}}
		`, M{})
}
