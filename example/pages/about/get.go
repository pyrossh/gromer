package about

import (
	"context"

	. "github.com/pyros2097/gromer/handlebars"
)

func GET(c context.Context) (HtmlContent, int, error) {
	return Html(`
		{{#Page title="About me"}}
			<div class="flex flex-col justify-center items-center">
					{{#Header}}
						A new link is here
					{{/Header}}
					<h1>About Me</h1>
			</div>
		{{/Page}}
		`).Render()
}
