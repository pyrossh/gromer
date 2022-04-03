package about

import (
	"context"

	. "github.com/pyros2097/gromer"
)

func GET(c context.Context) (HtmlContent, int, error) {
	return Html(`
	<!DOCTYPE html>
	<html lang="en">
			<head>`, nil)
}
