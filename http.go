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
)

var info *debug.BuildInfo
var IsCloundRun bool

var RouteDefs []RouteDefinition

type RouteDefinition struct {
	Pkg        string      `json:"pkg"`
	PkgPath    string      `json:"pkgPath"`
	Method     string      `json:"method"`
	Path       string      `json:"path"`
	PathParams []string    `json:"pathParams"`
	Params     interface{} `json:"params"`
}

type HtmlContent string
type HandlersTemplate struct {
	text string
	ctx  *handlebars.Context
}

func Html(tpl string) *HandlersTemplate {
	return &HandlersTemplate{text: tpl, ctx: handlebars.NewContext()}
}

func (t *HandlersTemplate) Prop(key string, v any) *HandlersTemplate {
	t.ctx.Set(key, v)
	return t
}

func (t *HandlersTemplate) Props(args ...any) *HandlersTemplate {
	for i := 0; i < len(args); i += 2 {
		key := fmt.Sprintf("%s", args[i])
		t.ctx.Set(key, args[i+1])
	}
	return t
}

func (t *HandlersTemplate) Render(args ...any) (HtmlContent, int, error) {
	s, err := handlebars.Render(t.text, t.ctx)
	if err != nil {
		return HtmlContent("Server Erorr"), 500, err
	}
	return HtmlContent(s), 200, nil
}

func HtmlErr(status int, err error) (HtmlContent, int, error) {
	return HtmlContent("ErrorPage/AccessDeniedPage/NotFoundPage based on status code"), status, err
}

func GetFunctionName(temp interface{}) string {
	strs := strings.Split((runtime.FuncForPC(reflect.ValueOf(temp).Pointer()).Name()), ".")
	return strs[len(strs)-1]
}

func RegisterComponent(fn interface{}, props ...string) {
	name := GetFunctionName(fn)
	fnType := reflect.TypeOf(fn)
	fnValue := reflect.ValueOf(fn)
	handlebars.GlobalHelpers.Add(name, func(title string, c handlebars.HelperContext) (template.HTML, error) {
		s, err := c.Block()
		if err != nil {
			return "", err
		}
		args := []reflect.Value{}
		if fnType.NumIn() > 0 {
			args = append(args, reflect.ValueOf(title))
			if s != "" {
				args = append(args, reflect.ValueOf(template.HTML(s)))
			}
		}
		res := fnValue.Call(args)
		tpl := res[0].Interface().(*HandlersTemplate)
		comp, err := handlebars.Render(tpl.text, tpl.ctx)
		if err != nil {
			return "", err
		}
		return template.HTML(comp), nil
	})
}

func init() {
	IsCloundRun = os.Getenv("K_REVISION") != ""
	info, _ = debug.ReadBuildInfo()
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	if !IsCloundRun {
		log.Logger = log.Output(zerolog.ConsoleWriter{
			Out:        os.Stdout,
			TimeFormat: zerolog.TimeFormatUnix,
		})
	}
}

func RespondError(w http.ResponseWriter, status int, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status) // always write status last
	merror := map[string]interface{}{
		"error": err.Error(),
	}
	if status >= 500 {
		log.Error().Str("type", "panic").Msg(err.Error())
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
	if v, ok := response.(HtmlContent); ok {
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

var LogMiddleware = mux.MiddlewareFunc(func(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				stack := string(debug.Stack())
				RespondError(w, 500, fmt.Errorf("panic: %+v\n %s", err, stack))
			}
		}()
		startTime := time.Now()
		logRespWriter := NewLogResponseWriter(w)
		next.ServeHTTP(logRespWriter, r)
		ip, _, _ := net.SplitHostPort(r.RemoteAddr)
		if len(ip) > 0 && ip[0] == '[' {
			ip = ip[1 : len(ip)-1]
		}
		log.Info().
			Str("method", r.Method).
			Str("url", r.URL.String()).
			Int("header_size", int(headerSize(r.Header))).
			Int64("body_size", r.ContentLength).
			Str("host", r.Host).
			// Str("agent", r.UserAgent()).
			Str("referer", r.Referer()).
			Str("proto", r.Proto).
			Str("remote_ip", ip).
			Int("status", logRespWriter.responseStatusCode).
			Int("resp_header_size", logRespWriter.responseHeaderSize).
			Int("resp_body_size", logRespWriter.responseContentLength).
			Str("latency", time.Since(startTime).String()).
			Msg("")
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

var NotFoundHandler = LogMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	RespondError(w, 404, fmt.Errorf("path '%s' not found", r.URL.String()))
}))

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
