package gromer

import (
	"encoding/json"
	"net/http"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"github.com/fatih/color"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func init() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
}

func RespondError(w http.ResponseWriter, status int, err error) {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	data, _ := json.Marshal(map[string]string{
		"error": err.Error(),
	})
	w.Write(data)
}

var pathParamsRegex = regexp.MustCompile(`{(.*?)}`)

func GetRouteParams(route string) []string {
	params := []string{}
	found := pathParamsRegex.FindAllString(route, -1)
	for _, v := range found {
		params = append(params, strings.Replace(strings.Replace(v, "}", "", 1), "{", "", 1))
	}
	return params
}

func PerformRequest(route string, h interface{}, ctx interface{}, w http.ResponseWriter, r *http.Request) (int, error) {
	params := GetRouteParams(route)
	args := []reflect.Value{reflect.ValueOf(ctx)}
	funcType := reflect.TypeOf(h)
	icount := funcType.NumIn()
	vars := mux.Vars(r)
	for _, k := range params {
		args = append(args, reflect.ValueOf(vars[k]))
	}
	if len(args) != icount {
		structType := funcType.In(icount - 1)
		instance := reflect.New(structType)
		if structType.Kind() != reflect.Struct {
			panic(route + " func final param should be a struct")
		}
		if r.Method == "POST" || r.Method == "PUT" || r.Method == "PATCH" {
			err := json.NewDecoder(r.Body).Decode(instance.Interface())
			if err != nil {
				RespondError(w, 400, err)
				return 400, err
			}
		} else if r.Method == "GET" {
			rv := instance.Elem()
			for i := 0; i < structType.NumField(); i++ {
				if f := rv.Field(i); f.CanSet() {
					jsonName := structType.Field(i).Tag.Get("json")
					jsonValue := r.URL.Query().Get(jsonName)
					if f.Kind() == reflect.String {
						f.SetString(jsonValue)
					} else if f.Kind() == reflect.Int64 || f.Kind() == reflect.Int32 || f.Kind() == reflect.Int {
						base := 64
						if f.Kind() == reflect.Int32 {
							base = 32
						}
						if jsonValue == "" {
							f.SetInt(0)
						} else {
							v, err := strconv.ParseInt(jsonValue, 10, base)
							if err != nil {
								RespondError(w, 400, err)
								return 400, err
							}
							f.SetInt(v)
						}
					} else {
						panic("Uknown query param: " + jsonName + " " + jsonValue)
					}
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
		return responseStatus, responseError.(error)
	}
	if v, ok := response.(HtmlPage); ok {
		w.Header().Set("Content-Type", "text/html")
		// This has to be at end always
		w.WriteHeader(responseStatus)
		v.WriteHtml(w)
		return 200, nil
	}
	w.Header().Set("Content-Type", "application/json")
	// This has to be at end always
	w.WriteHeader(responseStatus)
	data, _ := json.Marshal(response)
	w.Write(data)
	return 200, nil
}

func LogReq(status int, r *http.Request) {
	a := color.FgGreen
	if status >= 500 {
		a = color.FgRed
	} else if status >= 400 {
		a = color.FgYellow
	}
	m := color.FgCyan
	if r.Method == "POST" {
		m = color.FgYellow
	} else if r.Method == "PUT" {
		m = color.FgMagenta
	} else if r.Method == "DELETE" {
		m = color.FgRed
	}
	log.Info().Msgf("%3s %s %s", color.New(a).Sprint(status), color.New(m).Sprintf("%-4s", r.Method), color.WhiteString(r.URL.Path))
}
