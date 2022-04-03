package gromer

import (
	"encoding/json"
)

func ApiExplorer() (HtmlContent, int, error) {
	_, err := json.Marshal(RouteDefs)
	if err != nil {
		return HtmlErr(500, err)
	}
	return Html("", nil)
}
