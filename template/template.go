package template

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"

	"github.com/alecthomas/repr"
)

type ComponentFunc struct {
	Func interface{}
	Args []string
}

type Html func(string) string

var htmlTags = []string{"ul", "li", "span", "div"}
var compMap = map[string]ComponentFunc{}
var funcMap = map[string]interface{}{}

func getFunctionName(temp interface{}) string {
	strs := strings.Split((runtime.FuncForPC(reflect.ValueOf(temp).Pointer()).Name()), ".")
	return strs[len(strs)-1]
}

func RegisterComponent(f interface{}, args ...string) {
	name := getFunctionName(f)
	compMap[name] = ComponentFunc{
		Func: f,
		Args: args,
	}
}

func RegisterFunc(f interface{}) {
	name := getFunctionName(f)
	funcMap[name] = f
}

func getAttribute(k string, kvs []*Attribute) string {
	for _, param := range kvs {
		if param.Key == k {
			return strings.ReplaceAll(param.Value.Str, `"`, "")
		}
	}
	return ""
}

func subsRef(ctx map[string]interface{}, ref string) interface{} {
	if f, ok := funcMap[ref]; ok {
		return f.(func() string)()
	} else {
		parts := strings.Split(strings.ReplaceAll(ref, "!", ""), ".")
		if len(parts) == 2 {
			if v, ok := ctx[parts[0]]; ok {
				i := reflect.ValueOf(v).Elem().FieldByName(parts[1]).Interface()
				switch iv := i.(type) {
				case bool:
					if strings.Contains(ref, "!") {
						return !iv
					} else {
						return iv
					}
				case string:
					return iv
				}
			}
		}
	}
	return nil
}

func Render(x *Xml, ctx map[string]interface{}) string {
	space, _ := ctx["_space"].(string)
	s := space + "<" + x.Name
	if len(x.Attributes) > 0 {
		s += " "
	}
	for i, param := range x.Attributes {
		if len(param.Value.KV) != 0 {
			values := []string{}
			for _, kv := range param.Value.KV {
				if subsRef(ctx, kv.Value) == true {
					values = append(values, kv.Key)
				}
			}
			s += param.Key + `="` + strings.Join(values, "") + `"`
			repr.Println(s)
		} else if param.Value.Ref != "" {
			s += param.Key + `="` + subsRef(ctx, param.Value.Ref).(string) + `"`
		} else {
			s += param.Key + "=" + param.Value.Str
		}
		if i < len(x.Attributes)-1 {
			s += " "
		}
	}
	s += ">\n"
	if x.Value != nil {
		if x.Value.Ref != "" {
			s += space + "  " + subsRef(ctx, x.Value.Ref).(string) + "\n"
		}
	}
	if x.Name == "For" {
		ctxKey := getAttribute("key", x.Attributes)
		ctxName := getAttribute("itemKey", x.Attributes)
		data := ctx[ctxKey]
		switch reflect.TypeOf(data).Kind() {
		case reflect.Slice:
			v := reflect.ValueOf(data)
			for i := 0; i < v.Len(); i++ {
				ctx["_space"] = space + "  "
				ctx[ctxName] = v.Index(i).Interface()
				s += Render(x.Children[0], ctx) + "\n"
			}
		}
	} else {
		if comp, ok := compMap[x.Name]; ok {
			ctx["_space"] = space + "  "
			h := HtmlFunc(ctx)
			args := []reflect.Value{reflect.ValueOf(h)}
			for _, k := range comp.Args {
				if v, ok := ctx[k]; ok {
					args = append(args, reflect.ValueOf(v))
				} else {
					v := getAttribute(k, x.Attributes)
					args = append(args, reflect.ValueOf(v))
				}
			}
			result := reflect.ValueOf(comp.Func).Call(args)
			s += result[0].Interface().(string) + "\n"
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

func HtmlFunc(ctx map[string]interface{}) Html {
	return func(tpl string) string {
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
