package template

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
)

type Component func(map[string]interface{}) string

var htmlTags = []string{"ul", "li", "span", "div"}
var compMap = map[string]interface{}{}
var funcMap = map[string]interface{}{}

func getFunctionName(temp interface{}) string {
	strs := strings.Split((runtime.FuncForPC(reflect.ValueOf(temp).Pointer()).Name()), ".")
	return strs[len(strs)-1]
}

func RegisterComponent(f interface{}) {
	name := getFunctionName(f)
	compMap[name] = f
}

func RegisterFunc(f interface{}) {
	name := getFunctionName(f)
	funcMap[name] = f
}

func getAttribute(k string, kvs []*KeyValue) string {
	for _, param := range kvs {
		if param.Key == k {
			return strings.ReplaceAll(param.Value.Str, `"`, "")
		}
	}
	return ""
}

func Render(x *Xml, ctx map[string]interface{}) string {
	space, _ := ctx["_space"].(string)
	s := space + "<" + x.Name
	if len(x.Parameters) > 0 {
		s += " "
	}
	for i, param := range x.Parameters {
		s += param.Key + "=" + param.Value.Str
		if i < len(x.Parameters)-1 {
			s += " "
		}
	}
	s += ">\n"
	if x.Value != nil {
		if x.Value.Ref != "" {
			key := strings.ReplaceAll(strings.ReplaceAll(x.Value.Ref, "{", ""), "}", "")
			if f, ok := funcMap[key]; ok {
				s += f.(func() string)()
			} else {
				parts := strings.Split(key, ".")
				if len(parts) == 2 {
					if v, ok := ctx[parts[0]]; ok {
						s += reflect.ValueOf(v).Elem().FieldByName(parts[1]).Interface().(string)
					}
				}
			}
		}
	}
	if x.Name == "For" {
		ctxKey := getAttribute("key", x.Parameters)
		ctxName := getAttribute("name", x.Parameters)
		data := ctx[ctxKey]
		switch reflect.TypeOf(data).Kind() {
		case reflect.Slice:
			v := reflect.ValueOf(data)
			for i := 0; i < v.Len(); i++ {
				ctx["_space"] = space + "  "
				ctx[ctxName] = v.Index(i).Interface()
				s += Render(x.Children[0], ctx)
			}
		}
	} else {
		if comp, ok := compMap[x.Name]; ok {
			ctxKey := getAttribute("key", x.Parameters)
			result := reflect.ValueOf(comp).Call([]reflect.Value{reflect.ValueOf(ctx[ctxKey])})
			s += result[0].Interface().(string)
		} else {
			found := false
			for _, t := range htmlTags {
				if t == x.Name {
					found = true
				}
			}
			if !found {
				panic(fmt.Errorf("Comp not found %s", x.Name))
			}
		}
		for _, c := range x.Children {
			ctx["_space"] = space + "  "
			s += Render(c, ctx) + "\n"
		}
	}
	s += space + "</" + x.Name + ">"
	return s
}

func Html(ctx map[string]interface{}, tpl string) string {
	tree := &Module{}
	err := xmlParser.ParseBytes("filename", []byte(tpl), tree)
	if err != nil {
		panic(err)
	}
	o := ""
	for _, n := range tree.Nodes {
		v := Render(n, ctx)
		o += v
	}
	return o
}

// <script>
//     document.addEventListener('alpine:init', () => {
//         Alpine.store('todos', {
//         		list: [],
//						count: 0,
//         })
//     })
// </script>

// patch: {
// 	{ "op": "add", "path": "/todos/list", "value": { "id": "123", "text": "123" } },
// }
