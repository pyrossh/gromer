package cli

// import (
// 	"bytes"
// 	"fmt"
// 	"io/ioutil"
// 	"log"
// 	"net/http"
// 	"os"
// 	"os/exec"
// 	"path/filepath"
// 	"plugin"
// 	"strconv"
// 	"strings"
// 	"text/template"
// 	"time"

// 	"github.com/akrylysov/algnhsa"
// 	"github.com/markbates/pkger"
// 	app "github.com/pyros2097/wapp"
// 	"gopkg.in/fsnotify.v1"
// )

// const wasmExecTemplate = `const enosys=()=>{const a=new Error("not implemented");return a.code="ENOSYS",a};let outputBuf="";window.fs={constants:{O_WRONLY:-1,O_RDWR:-1,O_CREAT:-1,O_TRUNC:-1,O_APPEND:-1,O_EXCL:-1},writeSync(a,b){outputBuf+=decoder.decode(b);const c=outputBuf.lastIndexOf("\n");return-1!=c&&(console.log(outputBuf.substr(0,c)),outputBuf=outputBuf.substr(c+1)),b.length},write(a,b,c,d,e,f){if(0!==c||d!==b.length||null!==e)return void f(enosys());const g=this.writeSync(a,b);f(null,g)}};const encoder=new TextEncoder("utf-8"),decoder=new TextDecoder("utf-8");class Go{constructor(){this.argv=["js"],this.env={},this.exit=a=>{0!==a&&console.warn("exit code:",a)},this._exitPromise=new Promise(a=>{this._resolveExitPromise=a}),this._pendingEvent=null,this._scheduledTimeouts=new Map,this._nextCallbackTimeoutID=1;const a=(a,b)=>{this.mem.setUint32(a+0,b,!0),this.mem.setUint32(a+4,Math.floor(b/4294967296),!0)},b=a=>{const b=this.mem.getUint32(a+0,!0),c=this.mem.getInt32(a+4,!0);return b+4294967296*c},c=a=>{const b=this.mem.getFloat64(a,!0);if(0!==b){if(!isNaN(b))return b;const c=this.mem.getUint32(a,!0);return this._values[c]}},d=(a,b)=>{const c=2146959360;if("number"==typeof b)return isNaN(b)?(this.mem.setUint32(a+4,2146959360,!0),void this.mem.setUint32(a,0,!0)):0===b?(this.mem.setUint32(a+4,2146959360,!0),void this.mem.setUint32(a,1,!0)):void this.mem.setFloat64(a,b,!0);switch(b){case void 0:return void this.mem.setFloat64(a,0,!0);case null:return this.mem.setUint32(a+4,c,!0),void this.mem.setUint32(a,2,!0);case!0:return this.mem.setUint32(a+4,c,!0),void this.mem.setUint32(a,3,!0);case!1:return this.mem.setUint32(a+4,c,!0),void this.mem.setUint32(a,4,!0);}let d=this._ids.get(b);d===void 0&&(d=this._idPool.pop(),d===void 0&&(d=this._values.length),this._values[d]=b,this._goRefCounts[d]=0,this._ids.set(b,d)),this._goRefCounts[d]++;let e=1;switch(typeof b){case"string":e=2;break;case"symbol":e=3;break;case"function":e=4;}this.mem.setUint32(a+4,2146959360|e,!0),this.mem.setUint32(a,d,!0)},e=a=>{const c=b(a+0),d=b(a+8);return new Uint8Array(this._inst.exports.mem.buffer,c,d)},f=d=>{const e=b(d+0),f=b(d+8),g=Array(f);for(let a=0;a<f;a++)g[a]=c(e+8*a);return g},g=a=>{const c=b(a+0),d=b(a+8);return decoder.decode(new DataView(this._inst.exports.mem.buffer,c,d))},h=Date.now()-performance.now();this.importObject={go:{"runtime.wasmExit":a=>{const b=this.mem.getInt32(a+8,!0);this.exited=!0,delete this._inst,delete this._values,delete this._goRefCounts,delete this._ids,delete this._idPool,this.exit(b)},"runtime.wasmWrite":a=>{const c=b(a+8),d=b(a+16),e=this.mem.getInt32(a+24,!0);fs.writeSync(c,new Uint8Array(this._inst.exports.mem.buffer,d,e))},"runtime.resetMemoryDataView":()=>{this.mem=new DataView(this._inst.exports.mem.buffer)},"runtime.nanotime1":b=>{a(b+8,1e6*(h+performance.now()))},"runtime.walltime1":b=>{const c=new Date().getTime();a(b+8,c/1e3),this.mem.setInt32(b+16,1e6*(c%1e3),!0)},"runtime.scheduleTimeoutEvent":a=>{const c=this._nextCallbackTimeoutID;this._nextCallbackTimeoutID++,this._scheduledTimeouts.set(c,setTimeout(()=>{for(this._resume();this._scheduledTimeouts.has(c);)console.warn("scheduleTimeoutEvent: missed timeout event"),this._resume()},b(a+8)+1)),this.mem.setInt32(a+16,c,!0)},"runtime.clearTimeoutEvent":a=>{const b=this.mem.getInt32(a+8,!0);clearTimeout(this._scheduledTimeouts.get(b)),this._scheduledTimeouts.delete(b)},"runtime.getRandomData":a=>{crypto.getRandomValues(e(a+8))},"syscall/js.finalizeRef":a=>{const b=this.mem.getUint32(a+8,!0);if(this._goRefCounts[b]--,0===this._goRefCounts[b]){const a=this._values[b];this._values[b]=null,this._ids.delete(a),this._idPool.push(b)}},"syscall/js.stringVal":a=>{d(a+24,g(a+8))},"syscall/js.valueGet":a=>{const b=Reflect.get(c(a+8),g(a+16));a=this._inst.exports.getsp(),d(a+32,b)},"syscall/js.valueSet":a=>{Reflect.set(c(a+8),g(a+16),c(a+32))},"syscall/js.valueDelete":a=>{Reflect.deleteProperty(c(a+8),g(a+16))},"syscall/js.valueIndex":a=>{d(a+24,Reflect.get(c(a+8),b(a+16)))},"syscall/js.valueSetIndex":a=>{Reflect.set(c(a+8),b(a+16),c(a+24))},"syscall/js.valueCall":a=>{try{const b=c(a+8),e=Reflect.get(b,g(a+16)),h=f(a+32),i=Reflect.apply(e,b,h);a=this._inst.exports.getsp(),d(a+56,i),this.mem.setUint8(a+64,1)}catch(b){d(a+56,b),this.mem.setUint8(a+64,0)}},"syscall/js.valueInvoke":a=>{try{const b=c(a+8),e=f(a+16),g=Reflect.apply(b,void 0,e);a=this._inst.exports.getsp(),d(a+40,g),this.mem.setUint8(a+48,1)}catch(b){d(a+40,b),this.mem.setUint8(a+48,0)}},"syscall/js.valueNew":a=>{try{const b=c(a+8),e=f(a+16),g=Reflect.construct(b,e);a=this._inst.exports.getsp(),d(a+40,g),this.mem.setUint8(a+48,1)}catch(b){d(a+40,b),this.mem.setUint8(a+48,0)}},"syscall/js.valueLength":b=>{a(b+16,parseInt(c(b+8).length))},"syscall/js.valuePrepareString":b=>{const e=encoder.encode(c(b+8)+"");d(b+16,e),a(b+24,e.length)},"syscall/js.valueLoadString":a=>{const b=c(a+8);e(a+16).set(b)},"syscall/js.valueInstanceOf":a=>{this.mem.setUint8(a+24,c(a+8)instanceof c(a+16))},"syscall/js.copyBytesToGo":b=>{const d=e(b+8),f=c(b+32);if(!(f instanceof Uint8Array))return void this.mem.setUint8(b+48,0);const g=f.subarray(0,d.length);d.set(g),a(b+40,g.length),this.mem.setUint8(b+48,1)},"syscall/js.copyBytesToJS":b=>{const d=c(b+8),f=e(b+16);if(!(d instanceof Uint8Array))return void this.mem.setUint8(b+48,0);const g=f.subarray(0,d.length);d.set(g),a(b+40,g.length),this.mem.setUint8(b+48,1)},debug:a=>{console.log(a)}}}}async run(a){this._inst=a,this.mem=new DataView(this._inst.exports.mem.buffer),this._values=[NaN,0,null,!0,!1,window,this],this._goRefCounts=[],this._ids=new Map,this._idPool=[],this.exited=!1;let b=4096;const c=a=>{const c=b,d=encoder.encode(a+"\0");return new Uint8Array(this.mem.buffer,b,d.length).set(d),b+=d.length,0!=b%8&&(b+=8-b%8),c},d=this.argv.length,e=[];this.argv.forEach(a=>{e.push(c(a))}),e.push(0);const f=Object.keys(this.env).sort();f.forEach(a=>{e.push(c(a+"="+this.env[a]))}),e.push(0);const g=b;e.forEach(a=>{this.mem.setUint32(b,a,!0),this.mem.setUint32(b+4,0,!0),b+=8}),this._inst.exports.run(d,g),this.exited&&this._resolveExitPromise(),await this._exitPromise}_resume(){if(this.exited)throw new Error("Go program has already exited");this._inst.exports.resume(),this.exited&&this._resolveExitPromise()}_makeFuncWrapper(a){const b=this;return function(){const c={id:a,this:this,args:arguments};return b._pendingEvent=c,b._resume(),c.result}}}const go=new Go;WebAssembly.instantiateStreaming(fetch("__path__"),go.importObject).then(a=>go.run(a.instance)).catch(a=>console.error("could not load wasm",a));`

// func wasmExecJs(path string) string {
// 	return strings.Replace(wasmExecTemplate, "__path__", path, 1)
// }

// var watchDelta = 1000 * time.Millisecond

// type Watcher struct {
// 	rootdir     string
// 	watcher     *fsnotify.Watcher
// 	watchVendor bool
// 	update      chan string
// }

// // MustRegisterWatcher creates a new Watcher and starts listening to
// // given folders
// func MustRegisterWatcher() *Watcher {
// 	w := &Watcher{
// 		update:      make(chan string),
// 		watchVendor: false,
// 	}
// 	var err error
// 	w.watcher, err = fsnotify.NewWatcher()
// 	if err != nil {
// 		log.Fatalf("Could not register watcher: %s", err)
// 	}
// 	w.watchFolders()
// 	return w
// }

// // Watch listens file updates, and sends signal to
// // update channel when .go and .tmpl files are updated
// func (w *Watcher) Watch() {
// 	eventSent := false
// 	for {
// 		select {
// 		case event := <-w.watcher.Events:
// 			// discard chmod events
// 			if event.Op&fsnotify.Chmod != fsnotify.Chmod {
// 				// test files do not need a rebuild
// 				if isTestFile(event.Name) {
// 					continue
// 				}
// 				if !isWatchedFileType(event.Name) {
// 					continue
// 				}
// 				if eventSent {
// 					continue
// 				}
// 				eventSent = true
// 				// prevent consequent builds
// 				go func() {
// 					w.update <- event.Name
// 					time.Sleep(watchDelta)
// 					eventSent = false
// 				}()

// 			}
// 		case err := <-w.watcher.Errors:
// 			if err != nil {
// 				log.Fatalf("Watcher error: %s", err)
// 			}
// 			return
// 		}
// 	}
// }

// func isTestFile(fileName string) bool {
// 	return strings.HasSuffix(filepath.Base(fileName), "_test.go")
// }

// func isWatchedFileType(fileName string) bool {
// 	ext := filepath.Ext(fileName)

// 	return ext == ".go"
// }

// // Close closes the fsnotify watcher channel
// func (w *Watcher) Close() {
// 	w.watcher.Close()
// 	close(w.update)
// }

// // watchFolders recursively adds folders that will be watched against the changes,
// // starting from the working directory
// func (w *Watcher) watchFolders() {
// 	wd, err := os.Getwd()

// 	if err != nil {
// 		log.Fatalf("Could not get root working directory: %s", err)
// 	}

// 	filepath.Walk(wd, func(path string, info os.FileInfo, err error) error {
// 		// skip files
// 		if info == nil {
// 			log.Fatalf("wrong watcher package: %s", path)
// 		}

// 		if !info.IsDir() {
// 			return nil
// 		}

// 		if !w.watchVendor {
// 			// skip vendor directory
// 			vendor := fmt.Sprintf("%s/vendor", wd)
// 			if strings.HasPrefix(path, vendor) {
// 				return filepath.SkipDir
// 			}
// 		}

// 		// skip hidden folders
// 		if len(path) > 1 && strings.HasPrefix(filepath.Base(path), ".") {
// 			return filepath.SkipDir
// 		}

// 		w.addFolder(path)

// 		return err
// 	})
// }

// // addFolder adds given folder name to the watched folders, and starts
// // watching it for further changes
// func (w *Watcher) addFolder(name string) {
// 	println("watch: " + name)
// 	if err := w.watcher.Add(name); err != nil {
// 		log.Fatalf("Could not watch folder: %s", err)
// 	}
// }

// type RouteInfo struct {
// 	Path       string
// 	FuncName   string
// 	RenderFunc func(*app.RenderContext) app.UI
// }

// var routesMap = map[string]RouteInfo{}

// func getSOPath(p string) string {
// 	return "build/" + filepath.Base(filepath.Dir(p)) + ".so"
// }

// func getWasmPath(p string) string {
// 	return "build/" + filepath.Base(filepath.Dir(p)) + ".wasm"
// }

// func buildSo(path string) (string, error) {
// 	soPath := getSOPath(path)
// 	out, err := exec.Command("go", "build", "-buildmode=plugin", "-o", soPath, path).CombinedOutput()
// 	if err != nil {
// 		println(string(out))
// 		println(err.Error())
// 		return "", err
// 	}
// 	fmt.Printf("wrote: %s\n", soPath)
// 	return soPath, nil
// }

// func buildWasm(path string) (string, error) {
// 	wasmPath := getWasmPath(path)
// 	cmd := exec.Command("go", "build", "-o", wasmPath, path)
// 	cmd.Env = os.Environ()
// 	cmd.Env = append(cmd.Env, "GOOS=js", "GOARCH=wasm")
// 	out, err := cmd.CombinedOutput()
// 	if err != nil {
// 		println(string(out))
// 		println(err.Error())
// 		return "", err
// 	}
// 	fmt.Printf("wrote: %s\n", wasmPath)
// 	return wasmPath, nil
// }

// func getRoute(basePath, p string) (string, string) {
// 	clean := strings.Replace(strings.Replace(p, basePath, "", 1), ".go", "", -1)
// 	return strings.Replace(clean, "index", "", -1), strings.Title(strings.Replace(clean, "/", "", -1))
// }

// func writeMainFile(basePath string) {
// 	tpl, err := template.New("writeMain").Parse(`// GENERATED FILE DO NOT EDIT
// package main

// import (
// 	. "github.com/pyros2097/wapp"
// 	"github.com/pyros2097/wapp/js"
// )

// func main() {
// 	{{range $key, $element := .}}
// 		if js.Window.URL().Path == "{{.Path}}" {
// 			Run({{.FuncName}})
// 		}
// 	{{end}}
// }
// `)
// 	if err != nil {
// 		panic(err)
// 	}
// 	buf := bytes.NewBuffer(nil)
// 	err = tpl.Execute(buf, routesMap)
// 	if err != nil {
// 		panic(err)
// 	}
// 	err = ioutil.WriteFile("pages/main.go", buf.Bytes(), 0644)
// 	if err != nil {
// 		panic(err)
// 	}
// }

// func loadRoutes(p *plugin.Plugin, basePath string, dry bool) {
// 	err := filepath.Walk(basePath, func(path string, info os.FileInfo, err error) error {
// 		if err != nil {
// 			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
// 			return err
// 		}
// 		if !info.IsDir() {
// 			routePath, routeFunc := getRoute(basePath, path)
// 			if routePath == "/main" {
// 				return nil
// 			}
// 			routesMap[routePath] = RouteInfo{
// 				Path:     routePath,
// 				FuncName: routeFunc,
// 			}
// 			if dry {
// 				// println("Dry")
// 				return nil
// 			}
// 			fmt.Printf("route: %s routeFunc: %s\n", routePath, routeFunc)
// 			renderFn, err := p.Lookup(routeFunc)
// 			if err != nil {
// 				return err
// 			}
// 			routesMap[routePath] = RouteInfo{
// 				Path:       routePath,
// 				FuncName:   routeFunc,
// 				RenderFunc: renderFn.(func(*app.RenderContext) app.UI),
// 			}
// 			// println(createPage(routesMap[routePath](app.NewRenderContext())).String())
// 		}
// 		return nil
// 	})
// 	if err != nil {
// 		fmt.Printf("error walking the path %q: %v\n", basePath, err)
// 		panic(err)
// 	}
// }

// func CreatePage(title string, ui app.UI, wasmPath string) *bytes.Buffer {
// 	page := bytes.NewBuffer(nil)
// 	page.WriteString("<!DOCTYPE html>\n")
// 	elems := app.FilterUIElems(ui)
// 	app.Html(
// 		app.Head(
// 			app.Title(title),
// 			app.Meta("author", "pyros2097"),
// 			app.Meta("description", "Description"),
// 			app.Meta("keywords", ""),
// 			app.Meta("theme-color", ""),
// 			app.Meta("viewport", "width=device-width, initial-scale=1, maximum-scale=1, user-scalable=0, viewport-fit=cover"),
// 			app.Link("icon", "/assets/icon.png"),
// 			app.Link("apple-touch-icon", "/assets/icon.png"),
// 			app.Link("stylesheet", "/assets/styles.css"),
// 			app.Link("manifest", "manifest"),
// 			app.Script(`const enosys = () => { const a = new Error("not implemented"); return a.code = "ENOSYS", a }; let outputBuf = ""; window.fs = { constants: { O_WRONLY: -1, O_RDWR: -1, O_CREAT: -1, O_TRUNC: -1, O_APPEND: -1, O_EXCL: -1 }, writeSync(a, b) { outputBuf += decoder.decode(b); const c = outputBuf.lastIndexOf("\n"); return -1 != c && (console.log(outputBuf.substr(0, c)), outputBuf = outputBuf.substr(c + 1)), b.length }, write(a, b, c, d, e, f) { if (0 !== c || d !== b.length || null !== e) return void f(enosys()); const g = this.writeSync(a, b); f(null, g) } }; const encoder = new TextEncoder("utf-8"), decoder = new TextDecoder("utf-8"); class Go { constructor() { this.argv = ["js"], this.env = {}, this.exit = a => { 0 !== a && console.warn("exit code:", a) }, this._exitPromise = new Promise(a => { this._resolveExitPromise = a }), this._pendingEvent = null, this._scheduledTimeouts = new Map, this._nextCallbackTimeoutID = 1; const a = (a, b) => { this.mem.setUint32(a + 0, b, !0), this.mem.setUint32(a + 4, Math.floor(b / 4294967296), !0) }, b = a => { const b = this.mem.getUint32(a + 0, !0), c = this.mem.getInt32(a + 4, !0); return b + 4294967296 * c }, c = a => { const b = this.mem.getFloat64(a, !0); if (0 !== b) { if (!isNaN(b)) return b; const c = this.mem.getUint32(a, !0); return this._values[c] } }, d = (a, b) => { const c = 2146959360; if ("number" == typeof b) return isNaN(b) ? (this.mem.setUint32(a + 4, 2146959360, !0), void this.mem.setUint32(a, 0, !0)) : 0 === b ? (this.mem.setUint32(a + 4, 2146959360, !0), void this.mem.setUint32(a, 1, !0)) : void this.mem.setFloat64(a, b, !0); switch (b) { case void 0: return void this.mem.setFloat64(a, 0, !0); case null: return this.mem.setUint32(a + 4, c, !0), void this.mem.setUint32(a, 2, !0); case !0: return this.mem.setUint32(a + 4, c, !0), void this.mem.setUint32(a, 3, !0); case !1: return this.mem.setUint32(a + 4, c, !0), void this.mem.setUint32(a, 4, !0); }let d = this._ids.get(b); d === void 0 && (d = this._idPool.pop(), d === void 0 && (d = this._values.length), this._values[d] = b, this._goRefCounts[d] = 0, this._ids.set(b, d)), this._goRefCounts[d]++; let e = 1; switch (typeof b) { case "string": e = 2; break; case "symbol": e = 3; break; case "function": e = 4; }this.mem.setUint32(a + 4, 2146959360 | e, !0), this.mem.setUint32(a, d, !0) }, e = a => { const c = b(a + 0), d = b(a + 8); return new Uint8Array(this._inst.exports.mem.buffer, c, d) }, f = d => { const e = b(d + 0), f = b(d + 8), g = Array(f); for (let a = 0; a < f; a++)g[a] = c(e + 8 * a); return g }, g = a => { const c = b(a + 0), d = b(a + 8); return decoder.decode(new DataView(this._inst.exports.mem.buffer, c, d)) }, h = Date.now() - performance.now(); this.importObject = { go: { "runtime.wasmExit": a => { const b = this.mem.getInt32(a + 8, !0); this.exited = !0, delete this._inst, delete this._values, delete this._goRefCounts, delete this._ids, delete this._idPool, this.exit(b) }, "runtime.wasmWrite": a => { const c = b(a + 8), d = b(a + 16), e = this.mem.getInt32(a + 24, !0); fs.writeSync(c, new Uint8Array(this._inst.exports.mem.buffer, d, e)) }, "runtime.resetMemoryDataView": () => { this.mem = new DataView(this._inst.exports.mem.buffer) }, "runtime.nanotime1": b => { a(b + 8, 1e6 * (h + performance.now())) }, "runtime.walltime1": b => { const c = new Date().getTime(); a(b + 8, c / 1e3), this.mem.setInt32(b + 16, 1e6 * (c % 1e3), !0) }, "runtime.scheduleTimeoutEvent": a => { const c = this._nextCallbackTimeoutID; this._nextCallbackTimeoutID++, this._scheduledTimeouts.set(c, setTimeout(() => { for (this._resume(); this._scheduledTimeouts.has(c);)console.warn("scheduleTimeoutEvent: missed timeout event"), this._resume() }, b(a + 8) + 1)), this.mem.setInt32(a + 16, c, !0) }, "runtime.clearTimeoutEvent": a => { const b = this.mem.getInt32(a + 8, !0); clearTimeout(this._scheduledTimeouts.get(b)), this._scheduledTimeouts.delete(b) }, "runtime.getRandomData": a => { crypto.getRandomValues(e(a + 8)) }, "syscall/js.finalizeRef": a => { const b = this.mem.getUint32(a + 8, !0); if (this._goRefCounts[b]--, 0 === this._goRefCounts[b]) { const a = this._values[b]; this._values[b] = null, this._ids.delete(a), this._idPool.push(b) } }, "syscall/js.stringVal": a => { d(a + 24, g(a + 8)) }, "syscall/js.valueGet": a => { const b = Reflect.get(c(a + 8), g(a + 16)); a = this._inst.exports.getsp(), d(a + 32, b) }, "syscall/js.valueSet": a => { Reflect.set(c(a + 8), g(a + 16), c(a + 32)) }, "syscall/js.valueDelete": a => { Reflect.deleteProperty(c(a + 8), g(a + 16)) }, "syscall/js.valueIndex": a => { d(a + 24, Reflect.get(c(a + 8), b(a + 16))) }, "syscall/js.valueSetIndex": a => { Reflect.set(c(a + 8), b(a + 16), c(a + 24)) }, "syscall/js.valueCall": a => { try { const b = c(a + 8), e = Reflect.get(b, g(a + 16)), h = f(a + 32), i = Reflect.apply(e, b, h); a = this._inst.exports.getsp(), d(a + 56, i), this.mem.setUint8(a + 64, 1) } catch (b) { d(a + 56, b), this.mem.setUint8(a + 64, 0) } }, "syscall/js.valueInvoke": a => { try { const b = c(a + 8), e = f(a + 16), g = Reflect.apply(b, void 0, e); a = this._inst.exports.getsp(), d(a + 40, g), this.mem.setUint8(a + 48, 1) } catch (b) { d(a + 40, b), this.mem.setUint8(a + 48, 0) } }, "syscall/js.valueNew": a => { try { const b = c(a + 8), e = f(a + 16), g = Reflect.construct(b, e); a = this._inst.exports.getsp(), d(a + 40, g), this.mem.setUint8(a + 48, 1) } catch (b) { d(a + 40, b), this.mem.setUint8(a + 48, 0) } }, "syscall/js.valueLength": b => { a(b + 16, parseInt(c(b + 8).length)) }, "syscall/js.valuePrepareString": b => { const e = encoder.encode(c(b + 8) + ""); d(b + 16, e), a(b + 24, e.length) }, "syscall/js.valueLoadString": a => { const b = c(a + 8); e(a + 16).set(b) }, "syscall/js.valueInstanceOf": a => { this.mem.setUint8(a + 24, c(a + 8) instanceof c(a + 16)) }, "syscall/js.copyBytesToGo": b => { const d = e(b + 8), f = c(b + 32); if (!(f instanceof Uint8Array)) return void this.mem.setUint8(b + 48, 0); const g = f.subarray(0, d.length); d.set(g), a(b + 40, g.length), this.mem.setUint8(b + 48, 1) }, "syscall/js.copyBytesToJS": b => { const d = c(b + 8), f = e(b + 16); if (!(d instanceof Uint8Array)) return void this.mem.setUint8(b + 48, 0); const g = f.subarray(0, d.length); d.set(g), a(b + 40, g.length), this.mem.setUint8(b + 48, 1) }, debug: a => { console.log(a) } } } } async run(a) { this._inst = a, this.mem = new DataView(this._inst.exports.mem.buffer), this._values = [NaN, 0, null, !0, !1, window, this], this._goRefCounts = [], this._ids = new Map, this._idPool = [], this.exited = !1; let b = 4096; const c = a => { const c = b, d = encoder.encode(a + "\0"); return new Uint8Array(this.mem.buffer, b, d.length).set(d), b += d.length, 0 != b % 8 && (b += 8 - b % 8), c }, d = this.argv.length, e = []; this.argv.forEach(a => { e.push(c(a)) }), e.push(0); const f = Object.keys(this.env).sort(); f.forEach(a => { e.push(c(a + "=" + this.env[a])) }), e.push(0); const g = b; e.forEach(a => { this.mem.setUint32(b, a, !0), this.mem.setUint32(b + 4, 0, !0), b += 8 }), this._inst.exports.run(d, g), this.exited && this._resolveExitPromise(), await this._exitPromise } _resume() { if (this.exited) throw new Error("Go program has already exited"); this._inst.exports.resume(), this.exited && this._resolveExitPromise() } _makeFuncWrapper(a) { const b = this; return function () { const c = { id: a, this: this, args: arguments }; return b._pendingEvent = c, b._resume(), c.result } } } const go = new Go; WebAssembly.instantiateStreaming(fetch("/assets/main.wasm"), go.importObject).then(a => go.run(a.instance)).catch(a => console.error("could not load wasm", a));`),
// 		),
// 		app.Body(elems[0]),
// 	).Html(page)
// 	return page
// }

// func serve(wd string, isDev bool) {
// 	assetsFileServer := http.FileServer(pkger.Dir(filepath.Join(wd, "assets")))
// 	buildFileServer := http.FileServer(pkger.Dir(filepath.Join(wd, "build")))
// 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		println("path: " + r.URL.Path)
// 		if routeInfo, ok := routesMap[r.URL.Path]; ok {
// 			wasmPath := "/build/" + filepath.Base(wd) + ".wasm"
// 			page := CreatePage(routeInfo.RenderFunc(app.NewRenderContext()), wasmPath)
// 			w.Header().Set("Content-Length", strconv.Itoa(page.Len()))
// 			w.Header().Set("Content-Type", "text/html")
// 			w.WriteHeader(http.StatusOK)
// 			w.Write(page.Bytes())
// 		} else if strings.Contains(r.URL.Path, "/build") {
// 			r.URL.Path = strings.Replace(r.URL.Path, "/build", "", 1)
// 			buildFileServer.ServeHTTP(w, r)
// 		} else {
// 			r.URL.Path = strings.Replace(r.URL.Path, "/assets", "", 1)
// 			assetsFileServer.ServeHTTP(w, r)
// 		}
// 	})
// 	if isDev {
// 		println("Serving on HTTP port: 1234")
// 		http.ListenAndServe(":1234", nil)
// 	} else {
// 		println("algnhsa serving default mux")
// 		algnhsa.ListenAndServe(http.DefaultServeMux, nil)
// 	}
// }

// func buildAll(basePath string) {
// 	loadRoutes(nil, basePath, true)
// 	writeMainFile(basePath)
// 	buildWasm(basePath)
// 	soPath, err := buildSo(basePath)
// 	if err != nil {
// 		println("could not build")
// 		panic(err)
// 	}
// 	p, err := plugin.Open(soPath)
// 	if err != nil {
// 		println("could not load so plugin")
// 		panic(err)
// 	}
// 	// fmt.Printf("%+v\n", p)
// 	loadRoutes(p, basePath, false)
// }

// func Watch(isDev bool) {
// 	wd, err := os.Getwd()
// 	if err != nil {
// 		fmt.Printf("could not get wd")
// 		return
// 	}
// 	if isDev {
// 		basePath := filepath.Join(wd, "pages")
// 		buildAll(basePath)
// 		watcher := MustRegisterWatcher()
// 		go watcher.Watch()
// 		go serve(wd, isDev)
// 		for file := range watcher.update {
// 			println("changed: " + file)
// 			buildAll(basePath)
// 		}
// 	} else {
// 		serve(wd, isDev)
// 	}
// }
