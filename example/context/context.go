package context

import (
	"bytes"
	c "context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"reflect"

	"github.com/gobuffalo/velvet"
	"github.com/gorilla/mux"
	. "github.com/pyros2097/wapp"
	"golang.org/x/net/html"
)

type ReqContext struct {
	c.Context
	Path        string
	PathParams  map[string]string
	QueryParams map[string]string
	Body        []byte
	JS          *bytes.Buffer
	CSS         *bytes.Buffer
	UserID      string
}

func NewReqContext(r *http.Request) (ReqContext, error) {
	pathParams := mux.Vars(r)
	var body []byte
	var err error
	if r.Method == "POST" || r.Method == "PUT" || r.Method == "PATCH" {
		body, err = ioutil.ReadAll(r.Body)
		if err != nil {
			return ReqContext{}, err
		}
	}
	return ReqContext{
		Path:        r.URL.Path,
		QueryParams: map[string]string{},
		PathParams:  pathParams,
		Body:        body,
		JS:          bytes.NewBuffer(nil),
		CSS:         bytes.NewBuffer(nil),
		UserID:      "123",
	}, nil
}

type Handler func(c ReqContext) (interface{}, int, error)

func Wrap(h Handler) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, err := NewReqContext(r)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			data, _ := json.Marshal(M{
				"error": err.Error(),
			})
			w.Write(data)
			return
		}
		value, status, err := h(ctx)
		w.WriteHeader(status)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			data, _ := json.Marshal(M{
				"error": err.Error(),
			})
			w.Write(data)
			return
		}
		if v, ok := value.(*Element); ok {
			w.Header().Set("Content-Type", "text/html")
			v.WriteHtml(w)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		data, _ := json.Marshal(value)
		w.Write(data)
	}
}

func IsFunc(v interface{}) bool {
	return reflect.TypeOf(v).Kind() == reflect.Func
}

func UseData(ctx ReqContext, name string, data map[string]interface{}) {
	v := velvet.NewContext()
	v.Set("name", name)
	generatedMap := M{}
	for k, v := range data {
		if IsFunc(v) {
			f := v.(func() string)
			generatedMap[k] = "() {" + f() + "}"
		} else {
			generatedMap[k+":"] = v
		}
	}
	v.Set("data", generatedMap)
	s, err := velvet.Render(`
		Alpine.data('{{ name }}', () => {
			return {
				{{#each data}}
				{{ @key }}{{ @value }},
				{{/each}}
			};
		});
	`, v)
	if err != nil {
		panic(err)
	}
	ctx.JS.WriteString(s)
}

type Component func(ReqContext) string

var components = map[string]Component{}

func RegisterComponent(k string, v Component) {
	components[k] = v
}

func Html2(ctx ReqContext, input string, data map[string]interface{}) string {
	vctx := velvet.NewContext()
	for k, v := range data {
		vctx.Set(k, v)
	}
	doc, err := html.Parse(bytes.NewBufferString(input))
	if err != nil {
		panic(err)
	}
	var f func(*html.Node)
	f = func(n *html.Node) {
		if c, ok := components[n.Data]; ok {
			txt := c(ctx)
			println(n.Data, txt)
			// newNode, err := html.Parse(bytes.NewBufferString(txt))
			// if err != nil {
			// 	panic(err)
			// }
			println(n.NextSibling.Type)
			// n.RemoveChild(n.NextSibling)
			// n.RemoveChild(n.PrevSibling)
			// n.AppendChild(newNode)
		}
		if n.Type == html.ElementNode && n.Data == "h1" {
			println("h1")
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
	s, err := velvet.Render(input, vctx)
	if err != nil {
		panic(err)
	}
	return s
}
