package not_found_404

import (
	"context"

	. "github.com/pyros2097/gromer/gsx"
)

func GET(h Html, c context.Context) (string, int, error) {
	return h.Render(`
		{{#Page title="Page Not Found"}}
			{{#Header}}{{/Header}}
			<main class="box center">
				<h1>Page Not Found</h1>
				<h2 class="mt-6">
					<a class="is-underlined" href="/">Go Back</a>
				</h1>
			</main>
		{{/Page}}
	`), 404, nil
}
