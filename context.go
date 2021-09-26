package wapp

import (
	"bytes"
	c "context"
	"encoding/json"
	"net/http"
	"reflect"

	"github.com/gobuffalo/velvet"
	"github.com/gorilla/mux"
)

type WappContext struct {
	c.Context
	PathParams map[string]string
	JS         *bytes.Buffer
	CSS        *bytes.Buffer
}

func NewWappContext(r *http.Request) (WappContext, error) {
	pathParams := mux.Vars(r)
	return WappContext{
		Context:    r.Context(),
		PathParams: pathParams,
		JS:         bytes.NewBuffer(nil),
		CSS:        bytes.NewBuffer(nil),
	}, nil
}

func (r WappContext) UseData(name string, data map[string]interface{}) {
	v := velvet.NewContext()
	v.Set("name", name)
	generatedMap := map[string]interface{}{}
	for k, v := range data {
		if reflect.TypeOf(v).Kind() == reflect.Func {
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
	r.JS.WriteString(s)
}

func RespondError(w http.ResponseWriter, status int, err error) {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	data, _ := json.Marshal(map[string]string{
		"error": err.Error(),
	})
	w.Write(data)
}

func PerformRequest(h interface{}, ctx interface{}, w http.ResponseWriter, r *http.Request) error {
	args := []reflect.Value{reflect.ValueOf(ctx)}
	funcType := reflect.TypeOf(h)
	icount := funcType.NumIn()
	if icount == 2 {
		structType := funcType.In(1)
		instance := reflect.New(structType)
		if r.Method == "POST" || r.Method == "PUT" || r.Method == "PATCH" {
			err := json.NewDecoder(r.Body).Decode(instance.Interface())
			if err != nil {
				RespondError(w, 500, err)
				return err
			}
		} else if r.Method == "GET" {
			rv := instance.Elem()
			for i := 0; i < structType.NumField(); i++ {
				if f := rv.Field(i); f.CanSet() {
					jsonName := structType.Field(i).Tag.Get("json")
					jsonValue := r.URL.Query().Get(jsonName)
					f.SetString(jsonValue)
				}
			}
		}
		args = append(args, instance.Elem())
	}
	values := reflect.ValueOf(h).Call(args)
	response := values[0].Interface()
	responseStatus := values[1].Interface().(int)
	responseError := values[2].Interface()
	if responseError != nil {
		RespondError(w, responseStatus, responseError.(error))
		return responseError.(error)
	}
	if v, ok := response.(*Element); ok {
		w.WriteHeader(responseStatus)
		w.Header().Set("Content-Type", "text/html")
		v.WriteHtml(w)
		return nil
	}
	w.WriteHeader(responseStatus)
	w.Header().Set("Content-Type", "application/json")
	data, _ := json.Marshal(response)
	w.Write(data)
	return nil
}
