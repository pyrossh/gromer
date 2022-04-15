package gromer

import (
	"context"
	"embed"
	"encoding/json"
	"fmt"
	"html/template"
	"net"
	"net/http"
	"os"
	"reflect"
	"regexp"
	"runtime"
	"runtime/debug"
	"strconv"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/pyros2097/gromer/handlebars"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
	"xojoc.pw/useragent"
)

var info *debug.BuildInfo
var IsCloundRun bool

func init() {
	IsCloundRun = os.Getenv("K_REVISION") != ""
	info, _ = debug.ReadBuildInfo()
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	log.Logger = log.Output(zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: zerolog.TimeFormatUnix,
		NoColor:    IsCloundRun,
	})
}

var RouteDefs []RouteDefinition

type RouteDefinition struct {
	Pkg        string      `json:"pkg"`
	PkgPath    string      `json:"pkgPath"`
	Method     string      `json:"method"`
	Path       string      `json:"path"`
	PathParams []string    `json:"pathParams"`
	Params     interface{} `json:"params"`
}

func GetFunctionName(temp interface{}) string {
	strs := strings.Split((runtime.FuncForPC(reflect.ValueOf(temp).Pointer()).Name()), ".")
	return strs[len(strs)-1]
}

func RegisterComponent(fn any, props ...string) {
	name := GetFunctionName(fn)
	fnType := reflect.TypeOf(fn)
	fnValue := reflect.ValueOf(fn)
	handlebars.GlobalHelpers.Add(name, func(help handlebars.HelperContext) (template.HTML, error) {
		args := []reflect.Value{}
		var props any
		if fnType.NumIn() > 0 {
			structType := fnType.In(0)
			instance := reflect.New(structType)
			if structType.Kind() != reflect.Struct {
				log.Fatal().Msgf("component '%s' props should be a struct", name)
			}
			rv := instance.Elem()
			for i := 0; i < structType.NumField(); i++ {
				if f := rv.Field(i); f.CanSet() {
					jsonName := structType.Field(i).Tag.Get("json")
					if jsonName == "children" {
						s, err := help.Block()
						if err != nil {
							return "", err
						}
						f.Set(reflect.ValueOf(template.HTML(s)))
					} else {
						f.Set(reflect.ValueOf(help.Context.Get(jsonName)))
					}
				}
			}
			args = append(args, rv)
			props = rv.Interface()
		}
		res := fnValue.Call(args)
		tpl := res[0].Interface().(*handlebars.Template)
		tpl.Context.Set("props", props)
		s, _, err := tpl.Render()
		if err != nil {
			return "", err
		}
		return template.HTML(s), nil
	})
}

func RespondError(w http.ResponseWriter, status int, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status) // always write status last
	w.(*LogResponseWriter).SetError(err)
	merror := map[string]interface{}{
		"error": err.Error(),
	}
	if status >= 500 {
		merror["error"] = "Internal Server Error"
	}
	validationErrors, ok := err.(validator.ValidationErrors)
	if ok {
		merror["error"] = GetValidationError(validationErrors)
	}
	data, _ := json.Marshal(merror)
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

func addRouteDef(method, route string, h interface{}) {
	pathParams := GetRouteParams(route)
	var body any = nil
	funcType := reflect.TypeOf(h)
	if funcType.NumIn() > len(pathParams)+1 {
		structType := funcType.In(funcType.NumIn() - 1)
		instance := reflect.New(structType)
		if structType.Kind() != reflect.Struct {
			log.Fatal().Msgf("router  '%s' '%s' func final param should be a struct", method, route)
		}
		body = instance.Interface()
	}
	RouteDefs = append(RouteDefs, RouteDefinition{
		Method:     method,
		Path:       route,
		PathParams: pathParams,
		Params:     body,
	})
}

func PerformRequest(route string, h interface{}, ctx interface{}, w http.ResponseWriter, r *http.Request) {
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
			log.Fatal().Msgf("router '%s' func final param should be a struct", route)
		}
		if r.Method == "POST" || r.Method == "PUT" || r.Method == "PATCH" {
			err := json.NewDecoder(r.Body).Decode(instance.Interface())
			if err != nil {
				RespondError(w, 400, err)
				return
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
								return
							}
							f.SetInt(v)
						}
					} else if f.Kind() == reflect.Struct && f.Type().Name() == "Time" {
						if jsonValue == "" {
							f.Set(reflect.ValueOf(time.Time{}))
						} else {
							v, err := time.Parse(time.RFC3339, jsonValue)
							if err != nil {
								RespondError(w, 400, err)
								return
							}
							f.Set(reflect.ValueOf(v))
						}
					} else {
						log.Fatal().Msgf("Uknown query param: '%s' '%s'", jsonName, jsonValue)
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
		return
	}
	if v, ok := response.(handlebars.HtmlContent); ok {
		w.Header().Set("Content-Type", "text/html")
		// This has to be at end always
		w.WriteHeader(responseStatus)
		w.Write([]byte(v))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	// This has to be at end always
	w.WriteHeader(responseStatus)
	data, _ := json.Marshal(response)
	w.Write(data)
}

type writeCounter int64

func (wc *writeCounter) Write(p []byte) (n int, err error) {
	*wc += writeCounter(len(p))
	return len(p), nil
}
func headerSize(h http.Header) int64 {
	var wc writeCounter
	h.Write(&wc)
	return int64(wc) + 2 // for CRLF
}

type LogResponseWriter struct {
	http.ResponseWriter
	responseStatusCode    int
	responseContentLength int
	responseHeaderSize    int
	err                   error
}

func NewLogResponseWriter(w http.ResponseWriter) *LogResponseWriter {
	return &LogResponseWriter{ResponseWriter: w}
}

func (w *LogResponseWriter) WriteHeader(code int) {
	w.ResponseWriter.WriteHeader(code)
	w.responseStatusCode = code
	w.responseHeaderSize = int(headerSize(w.Header()))
}

func (w *LogResponseWriter) Write(body []byte) (int, error) {
	w.responseContentLength += len(body)
	return w.ResponseWriter.Write(body)
}

func (w *LogResponseWriter) SetError(err error) {
	w.err = err
}

var LogMiddleware = mux.MiddlewareFunc(func(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				stack := string(debug.Stack())
				RespondError(w, 599, fmt.Errorf("panic: %+v\n %s", err, stack))
			}
		}()
		logRespWriter := NewLogResponseWriter(w)
		next.ServeHTTP(logRespWriter, r)
		if IsCloundRun {
			return
		}
		startTime := time.Now()
		ip, _, _ := net.SplitHostPort(r.RemoteAddr)
		if len(ip) > 0 && ip[0] == '[' {
			ip = ip[1 : len(ip)-1]
		}
		logger := log.WithLevel(zerolog.InfoLevel)
		if logRespWriter.err != nil {
			logger = log.WithLevel(zerolog.ErrorLevel).Err(logRespWriter.err).Stack()
		}
		ua := useragent.Parse(r.UserAgent())
		logger.Msgf("%s %d %.2f KB %3s %s %s", r.Method,
			logRespWriter.responseStatusCode,
			float32(logRespWriter.responseContentLength)/1024.0,
			time.Since(startTime).Round(time.Millisecond).String(), ua.Name, r.URL.Path)

		// logger.
		// 	Str("method", r.Method).
		// 	Str("url", r.URL.String()).
		// 	Int("header_size", int(headerSize(r.Header))).
		// 	Int64("body_size", r.ContentLength).
		// 	Str("host", r.Host).
		// 	// Str("agent", r.UserAgent()).
		// 	Str("referer", r.Referer()).
		// 	Str("proto", r.Proto).
		// 	Str("remote_ip", ip).
		// 	Int("status", logRespWriter.responseStatusCode).
		// 	Int("resp_header_size", logRespWriter.responseHeaderSize).
		// 	Int("resp_body_size", logRespWriter.responseContentLength).
		// 	Str("latency", time.Since(startTime).String()).
		// 	Msgf("")
	})
})

func CorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		if r.Method == "OPTIONS" {
			w.WriteHeader(200)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func StatusHandler(h interface{}) http.Handler {
	return LogMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(context.WithValue(r.Context(), "url", r.URL), "header", r.Header)
		values := reflect.ValueOf(h).Call([]reflect.Value{reflect.ValueOf(ctx)})
		response := values[0].Interface()
		responseStatus := values[1].Interface().(int)
		responseError := values[2].Interface()
		if responseError != nil {
			RespondError(w, responseStatus, responseError.(error))
			return
		}
		w.Header().Set("Content-Type", "text/html")

		// This has to be at end always
		w.WriteHeader(responseStatus)
		w.Write([]byte(response.(handlebars.HtmlContent)))
	})).(http.Handler)
}

func WrapCache(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "public, max-age=2592000")
		h.ServeHTTP(w, r)
	})
}

func Static(router *mux.Router, path string, fs embed.FS) {
	router.PathPrefix(path).Handler(http.StripPrefix(path, WrapCache(http.FileServer(http.FS(fs)))))
}

func Handle(router *mux.Router, method, route string, h interface{}) {
	addRouteDef(method, route, h)
	router.HandleFunc(route, func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(context.WithValue(r.Context(), "url", r.URL), "header", r.Header)
		PerformRequest(route, h, ctx, w, r)
	}).Methods(method, "OPTIONS")
}
