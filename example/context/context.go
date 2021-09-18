package context

import (
	"bytes"
	c "context"
	"fmt"
	"net/http"
)

type ReqContext struct {
	c.Context
	UserID string
	JS     *bytes.Buffer
	CSS    *bytes.Buffer
}

func NewReqContext(w http.ResponseWriter, r *http.Request) *ReqContext {
	return &ReqContext{r.Context(), "123", bytes.NewBuffer(nil), bytes.NewBuffer(nil)}
}

func UseData(ctx *ReqContext, name, data string) {
	ctx.JS.WriteString(fmt.Sprintf(`
		Alpine.data('%s', () => {
			return %s;
		});
	`, name, data))
}

// func HTML(html string, data map[string]interface{}) string {
// 	return fmt.Sprintf("")
// }
