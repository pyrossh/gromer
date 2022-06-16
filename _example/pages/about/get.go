package about

import (
	"context"

	. "github.com/pyros2097/gromer/gsx"
)

func GET(h Html, c context.Context) (string, int, error) {
	return h.Render(`
		{{#Page title="About me"}}
			<div class="flex flex-col justify-center items-center">
					{{#Header}}
						A new link is here
					{{/Header}}
					<h1>About Me</h1>
			</div>
		{{/Page}}
	`), 200, nil
}
