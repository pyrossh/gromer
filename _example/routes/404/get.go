package not_found_404

import (
	"context"

	. "github.com/pyros2097/gromer/gsx"
)

func GET(h Context, c context.Context) (*Node, int, error) {
	return h.Render(`
		<Page title="Page Not Found">
			<main class="box center">
				<h1>Page Not Found</h1>
				<h2 class="mt-6">
					<a class="is-underlined" href="/">Go Back</a>
				</h2>
			</main>
		</Page>
	`), 404, nil
}
