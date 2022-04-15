package not_found_404

import (
	"context"

	. "github.com/pyros2097/gromer/handlebars"
)

func GET(c context.Context) (HtmlContent, int, error) {
	return Html(`
		{{#Page title="Page Not Found"}}
			{{#Header}}{{/Header}}
			<main class="box center">
				<h1>Page Not Found</h1>
				<h2 class="mt-6">
					<a class="is-underlined" href="/">Go Back</a>
				</h1>
			</main>
		{{/Page}}
		`).RenderWithStatus(404)
}
