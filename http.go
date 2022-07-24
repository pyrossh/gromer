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
	"github.com/google/uuid"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/pyros2097/gromer/gsx"
	"github.com/rotisserie/eris"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/segmentio/go-camelcase"
	"xojoc.pw/useragent"
)

var (
	info                  *debug.BuildInfo
	IsCloundRun           bool
	pathParamsRegex                       = regexp.MustCompile(`{(.*?)}`)
	globalStatusComponent StatusComponent = nil
)

type StatusComponent func(c *gsx.Context, status int, err error) []*gsx.Tag

func init() {
	IsCloundRun = os.Getenv("K_REVISION") != ""
	info, _ = debug.ReadBuildInfo()
	zerolog.ErrorStackMarshaler = func(err error) interface{} {
		return eris.ToJSON(err, true)
	}
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
	gsx.RegisterFunc(GetAssetUrl)
}

func getFunctionName(temp interface{}) string {
	strs := strings.Split((runtime.FuncForPC(reflect.ValueOf(temp).Pointer()).Name()), ".")
	return strs[len(strs)-1]
}

func RespondError(w http.ResponseWriter, r *http.Request, status int, err error) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(status) // always write status last
	if status >= 500 {
		formattedStr := eris.ToCustomString(err, eris.StringFormat{
			Options: eris.FormatOptions{
				WithTrace:    true,
				InvertOutput: true,
				InvertTrace:  true,
			},
			MsgStackSep:  "\n",
			PreStackSep:  "\t",
			StackElemSep: " | ",
			ErrorSep:     "\n",
		})
		log.Error().Msg(err.Error() + "\n" + formattedStr)
	}
	c := createCtx(r, "Status")
	c.Set("funcName", "error")
	c.Set("error", err.Error())
	if r.Header.Get("HX-Request") == "true" || globalStatusComponent == nil {
		tags := c.Render(`
			<div style="color: red;">
				<h1>{error}</h1>
			</div>
		`)
		gsx.Write(c, w, tags)
		return
	}
	tags := globalStatusComponent(c, status, err)
	gsx.Write(c, w, tags)
}

func GetRouteParams(route string) []string {
	params := []string{}
	found := pathParamsRegex.FindAllString(route, -1)
	for _, v := range found {
		params = append(params, strings.Replace(strings.Replace(v, "}", "", 1), "{", "", 1))
	}
	return params
}

func PerformRequest(route string, h interface{}, c interface{}, w http.ResponseWriter, r *http.Request, isJson bool) {
	params := GetRouteParams(route)
	args := []reflect.Value{reflect.ValueOf(c)}
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
				RespondError(w, r, 400, err)
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
								RespondError(w, r, 400, err)
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
								RespondError(w, r, 400, err)
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
				RespondError(w, r, 400, err)
				return
			}
		} else {
			RespondError(w, r, 400, eris.Errorf("Illegal Content-Type found %s", contentType))
			return
		}
		if !isJson {
			c.(*gsx.Context).Set("params", instance.Elem().Interface())
		}
		args = append(args, instance.Elem())
	}
	values := reflect.ValueOf(h).Call(args)
	response := values[0].Interface()
	responseStatus := values[1].Interface().(int)
	responseError := values[2].Interface()
	if responseError != nil {
		RespondError(w, r, responseStatus, eris.Wrap(responseError.(error), "Render failed"))
		return
	}
	if isJson {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(responseStatus)
		data, err := json.Marshal(response)
		if err != nil {
			RespondError(w, r, responseStatus, eris.Wrap(responseError.(error), "Json Marshal failed"))
			return
		}
		w.Write(data)
		return
	}
	w.Header().Set("Content-Type", "text/html")
	// This has to be at end always
	w.WriteHeader(responseStatus)
	if responseStatus != 204 {
		gsx.Write(c.(*gsx.Context), w, response.([]*gsx.Tag))
	}
}

func LogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		url := r.URL.Path
		if r.URL.RawQuery != "" {
			url += "?" + r.URL.RawQuery
		}
		ip, _, _ := net.SplitHostPort(r.RemoteAddr)
		if len(ip) > 0 && ip[0] == '[' {
			ip = ip[1 : len(ip)-1]
		}
		ua := useragent.Parse(r.UserAgent()).Name
		defer func() {
			if err := recover(); err != nil {
				log.Error().Msgf("%s 599 %s %s", r.Method, ua, url)
				RespondError(w, r, 599, eris.Errorf("%+v", err))
			}
		}()
		m := httpsnoop.CaptureMetrics(next, w, r)
		log.Info().Msgf("%s %d %.2fkb %s %s %s", r.Method, m.Code, float64(m.Written)/1024.0, m.Duration.Round(time.Millisecond).String(),
			ua,
			url,
		)
	})
}

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

func CompressMiddleware(next http.Handler) http.Handler {
	return handlers.CompressHandler(next)
}

func CacheMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "public, max-age=2592000") // perma cache for 1 month
		next.ServeHTTP(w, r)
	})
}

func StaticRoute(router *mux.Router, path string, fs embed.FS) {
	router.PathPrefix(path).Methods("GET").Handler(http.StripPrefix(path, http.FileServer(http.FS(fs))))
}

func IconsRoute(router *mux.Router, path string, fs embed.FS) {
	router.PathPrefix(path).Methods("GET").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			RespondError(w, r, 400, err)
			return
		}
		data, err := fs.ReadFile(strings.TrimPrefix(r.URL.Path, "/"))
		if err != nil {
			RespondError(w, r, 404, err)
			return
		}
		fill := r.Form.Get("fill")
		color := gsx.GetColor(fill)
		svg := strings.ReplaceAll(string(data), "<svg", fmt.Sprintf(`<svg fill="%s" `, color))
		w.Header().Set("Content-Type", "image/svg+xml")
		w.WriteHeader(200)
		w.Write([]byte(svg))
	})
}

func ComponentStylesRoute(router *mux.Router, route string) {
	router.Path(route).Methods("GET").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/css")
		w.WriteHeader(200)
		w.Write([]byte(gsx.GetComponentStyles()))
	})
}

func createCtx(r *http.Request, route string) *gsx.Context {
	newCtx := context.WithValue(context.WithValue(r.Context(), "url", r.URL), "header", r.Header)
	var hx *gsx.HX
	if r.Header.Get("HX-Request") == "true" {
		hx = &gsx.HX{
			Boosted:     r.Header.Get("HX-Boosted") == "true",
			CurrentUrl:  r.Header.Get("HX-Current-URL"),
			Prompt:      r.Header.Get("HX-Prompt"),
			Target:      r.Header.Get("HX-Target"),
			TriggerName: r.Header.Get("HX-Trigger-Name"),
			TriggerID:   r.Header.Get("HX-Trigger"),
		}
	}
	c := gsx.NewContext(newCtx, hx)
	c.Set("funcName", camelcase.Camelcase(route))
	c.Set("requestId", uuid.NewString())
	c.Link("stylesheet", "/gromer/css/normalize@3.0.0.css", "", "")
	c.Link("stylesheet", GetComponentsStylesUrl(), "", "")
	c.Link("icon", "/assets/favicon.ico", "image/x-icon", "image")
	c.Script("/gromer/js/htmx@1.7.0.js", false)
	c.Script("/gromer/js/hyperscript@0.9.6.js", false)
	// c.Script("/gromer/js/alpinejs@3.9.6.js", true)
	return c
}

func RegisterStatusHandler(router *mux.Router, comp StatusComponent) {
	globalStatusComponent = comp
	router.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c := createCtx(r, "Status")
		tags := comp(c, 404, nil)
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(404)
		gsx.Write(c, w, tags)
	})
}

func PageRoute(router *mux.Router, route string, page, action interface{}) {
	router.HandleFunc(route, func(w http.ResponseWriter, r *http.Request) {
		c := createCtx(r, route)
		if r.Method == "GET" {
			PerformRequest(route, page, c, w, r, false)
		} else {
			PerformRequest(route, action, c, w, r, false)
		}
	}).Methods("GET", "POST")
}

func ApiRoute(router *mux.Router, method, route string, h interface{}) {
	router.HandleFunc(route, func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(context.WithValue(r.Context(), "url", r.URL), "header", r.Header)
		PerformRequest(route, h, ctx, w, r, true)
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

func GetComponentsStylesUrl() string {
	sum := getSum("components.css", func() [16]byte {
		return md5.Sum([]byte(gsx.GetComponentStyles()))
	})
	return fmt.Sprintf("/components.css?hash=%s", sum)
}
