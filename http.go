package gromer

import (
	"context"
	"crypto/md5"
	"embed"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"os"
	"reflect"
	"regexp"
	"runtime"
	"runtime/debug"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/felixge/httpsnoop"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/pyros2097/gromer/gsx"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
	"xojoc.pw/useragent"
)

const (
	gzipEncoding   = "gzip"
	flateEncoding  = "deflate"
	acceptEncoding = "Accept-Encoding"
)

var (
	info            *debug.BuildInfo
	IsCloundRun     bool
	routeDefs       []RouteDefinition
	pathParamsRegex = regexp.MustCompile(`{(.*?)}`)
)

type RouteDefinition struct {
	Pkg        string      `json:"pkg"`
	PkgPath    string      `json:"pkgPath"`
	Method     string      `json:"method"`
	Path       string      `json:"path"`
	PathParams []string    `json:"pathParams"`
	Params     interface{} `json:"params"`
}

func init() {
	IsCloundRun = os.Getenv("K_REVISION") != ""
	info, _ = debug.ReadBuildInfo()
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	if IsCloundRun {
		zerolog.LevelFieldName = "severity"
		zerolog.TimestampFieldName = "timestamp"
		zerolog.TimeFieldFormat = time.RFC3339Nano
	} else {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
		log.Logger = log.Output(zerolog.ConsoleWriter{
			Out:          os.Stdout,
			NoColor:      IsCloundRun,
			PartsExclude: []string{zerolog.TimestampFieldName},
		})
	}
	gsx.RegisterFunc(GetStylesUrl)
	gsx.RegisterFunc(GetAssetUrl)
}

func getFunctionName(temp interface{}) string {
	strs := strings.Split((runtime.FuncForPC(reflect.ValueOf(temp).Pointer()).Name()), ".")
	return strs[len(strs)-1]
}

func RespondError(w http.ResponseWriter, status int, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status) // always write status last
	merror := map[string]interface{}{
		"error": err.Error(),
	}
	if status >= 500 {
		merror["error"] = "Internal Server Error"
		stack := string(debug.Stack())
		println(stack)
		log.WithLevel(zerolog.ErrorLevel).Err(err).Bool("panic", status == 599).Str("stack", stack).Stack()
	}
	validationErrors, ok := err.(validator.ValidationErrors)
	if ok {
		merror["error"] = GetValidationError(validationErrors)
	}
	data, _ := json.Marshal(merror)
	w.Write(data)
}

func GetRouteParams(route string) []string {
	params := []string{}
	found := pathParamsRegex.FindAllString(route, -1)
	for _, v := range found {
		params = append(params, strings.Replace(strings.Replace(v, "}", "", 1), "{", "", 1))
	}
	return params
}

func PerformRequest(route string, h interface{}, ctx context.Context, w http.ResponseWriter, r *http.Request) {
	params := GetRouteParams(route)
	renderContext := gsx.NewContext(ctx, r.Header.Get("HX-Request") == "true")
	renderContext.Set("requestId", uuid.NewString())
	renderContext.Link("stylesheet", GetStylesUrl(), "", "")
	renderContext.Link("icon", "/assets/favicon.ico", "image/x-icon", "image")
	renderContext.Script("/gromer/js/htmx@1.7.0.js", false)
	renderContext.Script("/gromer/js/alpinejs@3.9.6.js", true)
	args := []reflect.Value{reflect.ValueOf(renderContext)}
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
		method := r.Method
		contentType := r.Header.Get("Content-Type")
		if method == "GET" || ((method == "POST" || method == "PUT" || method == "PATCH") && contentType == "application/x-www-form-urlencoded") {
			err := r.ParseForm()
			if err != nil {
				RespondError(w, 400, err)
				return
			}
			rv := instance.Elem()
			for i := 0; i < structType.NumField(); i++ {
				if f := rv.Field(i); f.CanSet() {
					jsonName := structType.Field(i).Tag.Get("json")
					jsonValue := ""
					if method == "GET" {
						jsonValue = r.URL.Query().Get(jsonName)
					} else {
						jsonValue = r.Form.Get(jsonName)
					}
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
						log.Fatal().Msgf("Uknown form param: '%s' '%s'", jsonName, jsonValue)
					}
				}
			}
		} else if (method == "POST" || method == "PUT" || method == "PATCH") && contentType == "application/json" {
			err := json.NewDecoder(r.Body).Decode(instance.Interface())
			if err != nil {
				RespondError(w, 400, err)
				return
			}
		} else {
			RespondError(w, 400, fmt.Errorf("Illegal Content-Type tag found %s", contentType))
			return
		}
		renderContext.Set("params", instance.Elem().Interface())
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
	if v, ok := response.(*gsx.Node); ok {
		w.Header().Set("Content-Type", "text/html")
		// This has to be at end always
		w.WriteHeader(responseStatus)
		v.Write(renderContext, w)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	// This has to be at end always
	w.WriteHeader(responseStatus)
	data, _ := json.Marshal(response)
	w.Write(data)
}

func LogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				RespondError(w, 599, fmt.Errorf("%+v", err))
			}
		}()
		m := httpsnoop.CaptureMetrics(next, w, r)
		ip, _, _ := net.SplitHostPort(r.RemoteAddr)
		if len(ip) > 0 && ip[0] == '[' {
			ip = ip[1 : len(ip)-1]
		}
		log.WithLevel(zerolog.InfoLevel).Msgf("%s %d %.2fkb %s %s %s", r.Method,
			m.Code,
			float64(m.Written)/1024.0,
			m.Duration.Round(time.Millisecond).String(),
			useragent.Parse(r.UserAgent()).Name,
			r.URL.Path,
		)
	})
}

func CompressMiddleware(next http.Handler) http.Handler {
	return handlers.CompressHandler(next)
}

func CacheMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "public, max-age=2592000") // perma cache for 1 month
		next.ServeHTTP(w, r)
	})
}

func StatusHandler(h interface{}) http.Handler {
	return LogMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(context.WithValue(r.Context(), "url", r.URL), "header", r.Header)
		renderContext := gsx.NewContext(ctx, r.Header.Get("HX-Request") == "true")
		values := reflect.ValueOf(h).Call([]reflect.Value{reflect.ValueOf(renderContext)})
		response := values[0].Interface()
		responseStatus := values[1].Interface().(int)
		responseError := values[2].Interface()
		if responseError != nil {
			RespondError(w, responseStatus, responseError.(error))
			return
		}
		w.Header().Set("Content-Type", "text/html")

		// This has to be at end always after headers are set
		w.WriteHeader(responseStatus)
		response.(*gsx.Node).Write(renderContext, w)
	})).(http.Handler)
}

func StaticRoute(router *mux.Router, path string, fs embed.FS) {
	router.PathPrefix(path).Methods("GET").Handler(http.StripPrefix(path, http.FileServer(http.FS(fs))))
}

func StylesRoute(router *mux.Router, path string) {
	router.Path(path).Methods("GET").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/css")
		w.WriteHeader(200)
		w.Write([]byte(gsx.GetStyles()))
	})
}

func Handle(router *mux.Router, method, route string, h interface{}) {
	router.HandleFunc(route, func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(context.WithValue(r.Context(), "url", r.URL), "header", r.Header)
		PerformRequest(route, h, ctx, w, r)
	}).Methods(method, "OPTIONS")
}

func GetUrl(ctx context.Context) *url.URL {
	return ctx.Value("url").(*url.URL)
}

func GetHeader(ctx context.Context) http.Header {
	return ctx.Value("header").(http.Header)
}

var sumCache = sync.Map{}

func getSum(k string, cb func() [16]byte) string {
	if v, ok := sumCache.Load(k); ok {
		return v.(string)
	}
	sum := fmt.Sprintf("%x", cb())
	sumCache.Store(k, sum)
	return sum
}

func GetAssetUrl(fs embed.FS, path string) string {
	sum := getSum(path, func() [16]byte {
		data, err := fs.ReadFile(path)
		if err != nil {
			panic(err)
		}
		return md5.Sum(data)
	})
	return fmt.Sprintf("/assets/%s?hash=%s", path, sum)
}

func GetStylesUrl() string {
	sum := getSum("styles.css", func() [16]byte {
		return md5.Sum([]byte(gsx.GetStyles()))
	})
	return fmt.Sprintf("/styles.css?hash=%s", sum)
}
