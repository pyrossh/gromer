package about

import (
	"context"

	. "github.com/pyros2097/gromer/gsx"
)

func GET(h Context, c context.Context) (*Node, int, error) {
	return h.Render(`
		<Page title="About me">
			<div class="flex flex-col justify-center items-center">
				A new link is here
				P<h1>About Me</h1>
			</div>
		</Page>
	`), 200, nil
}
