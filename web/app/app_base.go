package app

import (
	"bytes"
	"context"
	"github.com/alexedwards/scs"
	"github.com/alexedwards/scs/stores/memstore"
	"github.com/goradd/gengen/pkg/maps"
	"github.com/goradd/goradd/pkg/base"
	"github.com/goradd/goradd/pkg/config"
	"github.com/goradd/goradd/pkg/goradd"
	"github.com/goradd/goradd/pkg/html"
	grlog "github.com/goradd/goradd/pkg/log"
	buf2 "github.com/goradd/goradd/pkg/pool"
	"github.com/goradd/goradd/pkg/resource"
	"github.com/goradd/goradd/pkg/session"
	strings2 "github.com/goradd/goradd/pkg/strings"
	"github.com/goradd/goradd/pkg/sys"
	"log"
	"path/filepath"
	"strings"
	"time"

	"github.com/goradd/goradd/pkg/messageServer"
	"github.com/goradd/goradd/pkg/page"
	"net/http"
	"os"
)

// StaticDirectoryPaths is a map of patterns to directory locations to serve statically.
// These can be registered at the command line or in the application
var StaticDirectoryPaths *maps.StringSliceMap

// StaticBlacklist is the list of file terminators that specify what files we always want to hide from view
// when a static file directory is searched. The default will always hide .go files. Add to it if you have
// other kinds of files in your static directories that you do not want to show. Do this only at startup.
var StaticBlacklist = []string{".go"}

type staticFileProcessor struct {
	ending    string
	processor StaticFileProcessorFunc
}

type StaticFileProcessorFunc func(file string, w http.ResponseWriter, r *http.Request)

// StaticFileProcessors is a map that connects file endings to processors that will process the content and return it
// to the output stream, bypassing other means of prcessing static files.
var staticFileProcessors []staticFileProcessor

// The application interface. A minimal set of commands that the main routine will ask the application to do.
// The main routine offers a way of creating mock applications, and alternate versions of the application from the default
type ApplicationI interface {
	Init()
	ServeHTTP(w http.ResponseWriter, r *http.Request)
	PutContext(*http.Request) *http.Request
	SetupErrorPageTemplate()
	SetupPageCaching()
	InitializeLoggers()
	SetupAssetDirectories()
	SetupSessionManager()
	WebSocketAuthHandler(next http.Handler) http.Handler
	SessionHandler(next http.Handler) http.Handler
	ServeRequest(w http.ResponseWriter, r *http.Request)
	ServeStaticFile(w http.ResponseWriter, r *http.Request) bool
	ServeApiRequest(w http.ResponseWriter, r *http.Request) bool
}

// The application base, to be embedded in your application
type Application struct {
	base.Base
}

func (a *Application) Init(self ApplicationI) {
	a.Base.Init(self)

	self.SetupErrorPageTemplate()
	self.SetupPageCaching()
	self.InitializeLoggers()
	self.SetupAssetDirectories()
	self.SetupSessionManager()

	page.DefaultCheckboxLabelDrawingMode = html.LabelAfter
}

func (a *Application) this() ApplicationI {
	return a.Self.(ApplicationI)
}

// SetupErrorPageTemplate sets the template that controls the output when an error happens during the processing of a
// page request, including any code that panics. By default, in debug mode, it will popup an error message in the browser with debug
// information when an error occurs. And in release mode it will popup a simple message that an error occurred and will log the
// error to the error log. You can implement this function in your local application object to override it and do something different.
func (a *Application) SetupErrorPageTemplate() {
	if config.Debug {
		page.ErrorPageFunc = page.DebugErrorPageTmpl
	} else {
		page.ErrorPageFunc = page.ReleaseErrorPageTmpl
	}
}

// SetupPageCaching sets up the service that saves pagestate information that reflects the state of a goradd form to
// our go code. The default sets up a one server-one process cache that does not scale, which works great for development, testing, and
// for moderate amounts of traffic. Override and replace the page cache with one that serializes the page state and saves
// it to a database to make it scalable.
func (a *Application) SetupPageCaching() {
	// Control how pages are cached. This will vary depending on whether you are using multiple machines to run your app,
	// and whether you are in development mode, etc. This default is for an in-memory store on one server and only one
	// process on that server. It basically does not serialize anything and leaves the entire formstate intact in memory.
	// This makes for a very fast server, but one that takes up quite a bit of RAM if you have a lot of simultaneous users.
	page.SetPageCache(page.NewFastPageCache(1000, 60*60*24))

	// Control how pages are serialized if a serialization cache is being used. This version uses the gob encoder.
	// You likely will not need to change this, but you might if your database cannot handle binary data.
	page.SetPageEncoder(page.GobPageEncoder{})
}

// InitializeLoggers sets up the various types of logs for various types of builds. By default, the DebugLog
// and FrameworkDebugLogs will be deactivated when the config.Debug variables are false. Otherwise, configure how you
// want, and simply remove a log if you don't want it to log anything.
func (a *Application) InitializeLoggers() {
	grlog.CreateDefaultLoggers()
}

// SetupAssetDirectories registers default directories that will contain web assets. These assets are served up in
// place in development mode, and served from a specified asset directory in release mode. This means the assets will
// need to be copied to a central location and moved to the release server. See the build directory for info.
func (a *Application) SetupAssetDirectories() {
	page.RegisterAssetDirectory(config.GoraddAssets(), config.AssetPrefix+"goradd")
	page.RegisterAssetDirectory(config.ProjectAssets(), config.AssetPrefix+"project")

	// If serving static html out of the root path, this will point it to the HtmlDirectory
	if dir := config.HtmlDirectory(); dir != "" {
		RegisterStaticPath("/", dir)
	}
}

// SetupSessionManager sets up the session manager. The session can be used to save data that is specific to a user
// and specific to the user's time on a browser. Sessions are often used to save login credentials so that you know
// the current user is logged in.
//
// The default uses a 3rd party session manager, and stores the session in memory, which is useful for development,
// testing, debugging, and for moderately used websites. The default does not scale, so replace it with a different
// storage mechanism is you are launching multiple copies of the app.
func (a *Application) SetupSessionManager() {
	// create the session manager. The default uses an in-memory storage engine. Change as you see fit.
	interval, _ := time.ParseDuration("24h")
	session.SetSessionManager(session.NewScsManager(scs.NewManager(memstore.New(interval))))
}

func (a *Application) PutContext(r *http.Request) *http.Request {
	return page.PutContext(r, os.Args[1:])
}

func (a *Application) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	pm := page.GetPageManager()
	if pm == nil {
		panic("No page manager defined")
	}

	ctx := r.Context()
	buf := page.OutputBuffer(ctx)
	if pm.IsPage(r.URL.Path) {
		headers, errCode := pm.RunPage(ctx, buf)
		if headers != nil {
			for k, v := range headers {
				// Multi-value headers can simply be separated with commas I believe
				w.Header().Set(k, v)
			}
		}
		if errCode != 0 {
			w.WriteHeader(errCode)
		}
	}
}

// MakeWebsocketMux creates the mux for the default websocket handler. The default handler provides session data to
// the web socket handler below, since its very common to need to get to session data to authenticate the user before
// responding to the request.
func (a *Application) MakeWebsocketMux() *http.ServeMux {
	mux := http.NewServeMux()

	mux.Handle("/ws", a.this().SessionHandler(a.this().WebSocketAuthHandler(messageServer.WebsocketHandler())))

	return mux
}

// WebSocketAuthHandler is the default authenticator of the web socket. This version simply makes sure the form
// has a pagestate, since if it doesn't, we should not be handling a request. If you want to authenticate using
// information out of the session, like to see whether the user is logged in, you should override this in your
// application instance.
func (a *Application) WebSocketAuthHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		pagestate := r.FormValue("id")

		if !page.GetPageManager().HasPage(pagestate) {
			// The page manager has no record of the pagestate, so either it is expired or never existed
			return // TODO: return error?
		}

		next.ServeHTTP(w, r)
	})
}

// MakeAppServer creates the handler chain that will handle http requests. There are a ton of ways to do this, 3rd party
// libraries to help with this, and middlewares you can use. This is a working example, and not a declaration of any
// "right" way to do this, since it can be very application specific. Generally you must make sure that
// PutContextHandler is called before ServeAppHandler in the chain.
func (a *Application) MakeAppServer() http.Handler {
	// the handler chain gets built in the reverse order of getting called
	buf := buf2.GetBuffer()
	defer buf2.PutBuffer(buf)

	// These handlers are called in reverse order
	h := a.ServeRequestHandler(buf)
	h = a.ServeStaticFileHandler(buf, h) // TODO: Speed this handler up by checking to see if the url is a goradd form before deciding to get context and session
	h = a.ServeAppHandler(buf, h)
	h = a.PutContextHandler(h)
	h = a.this().SessionHandler(h)
	h = a.BufferOutputHandler(h)

	return h
}

// SessionHandler initializes the global session handler. This default version uses the scs session handler. Feel
// free to replace it with the session handler of your choice.
func (a *Application) SessionHandler(next http.Handler) http.Handler {
	return session.Use(next)
}

// ServeRequestHandler is the last handler on the default call chain. It calls ServeRequest so the sub-class can handle it.
func (a *Application) ServeRequestHandler(buf *bytes.Buffer) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if page.OutputBuffer(r.Context()).Len() == 0 {
			a.this().ServeRequest(w, r)
		}
	}
	return http.HandlerFunc(fn)
}

// ServeStaticFileHandler serves up static files by calling ServeStaticFile.
func (a *Application) ServeStaticFileHandler(buf *bytes.Buffer, next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {

		if !a.this().ServeStaticFile(w, r) && next != nil {
			next.ServeHTTP(w, r)
		}
	}
	return http.HandlerFunc(fn)
}

// ServeAppHandler processes requests for goradd forms
func (a *Application) ServeAppHandler(buf *bytes.Buffer, next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		a.ServeHTTP(w, r)

		if next != nil && page.OutputBuffer(r.Context()).Len() == 0 {
			next.ServeHTTP(w, r)
		}
	}
	return http.HandlerFunc(fn)
}

// PutContextHandler is an http handler that adds the application context to the current context.
func (a *Application) PutContextHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		r = a.this().PutContext(r)
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

// BufferOutputHandler manages the buffering of http output. It must be the last item in the handler list.
func (a *Application) BufferOutputHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		// Setup the output buffer
		outBuf := buf2.GetBuffer()
		ctx := r.Context()
		ctx = context.WithValue(ctx, goradd.BufferContext, outBuf)
		r = r.WithContext(ctx)

		defer buf2.PutBuffer(outBuf)
		next.ServeHTTP(w, r)
		_, _ = w.Write(outBuf.Bytes())
	}
	return http.HandlerFunc(fn)
}

// ServeStaticFile serves up static html and other files found in registered directories.
// If the file is not found, it will return false.
func (a *Application) ServeStaticFile(w http.ResponseWriter, r *http.Request) bool {
	url := r.URL.Path
	var path string

	// StaticDirectoryPaths should be sorted longest to shortest at this point
	StaticDirectoryPaths.Range(func(pattern string, dir string) bool {
		if strings2.StartsWith(url, pattern) {
			fPath := strings.TrimPrefix(url, pattern)
			if fPath != "" && fPath[0:1] != "/" {
				// We only matched part of a directory, so not a match
				return true // go to next iteration
			}
			cleaned := strings.TrimPrefix(fPath, "..") // This prevents someone from hacking by using .. to refer to files outside of the directory
			cleaned = filepath.Clean(cleaned)
			path = filepath.Join(dir, cleaned)
			return false // stop iterating
		}
		return true
	})

	if path == "" {
		return false // not found
	}

	for _, bl := range StaticBlacklist {
		if strings2.EndsWith(path, bl) {
			return false // cannot show this kind of file
		}
	}

	if sys.IsDir(path) {
		path = filepath.Join(path, "index.html")
	}

	if sys.PathExists(path) {
		for _, p := range staticFileProcessors {
			if strings2.EndsWith(path, p.ending) {
				p.processor(path, w, r)
				return true
			}
		}

		http.ServeFile(w, r, path)
		return true
	}

	return false // indicates no static file was found
}

// ServeRequest is the place to serve up any files that have not been handled in any other way, either by a previously
// declared handler, or by the goradd app server. ServeRequest is only called when all
// the other methods have failed. Override it to handle other files,
// or to change the messaging when a bad url is attempted.
func (a *Application) ServeRequest(w http.ResponseWriter, r *http.Request) {
	if !resource.HandleRequest(w, r) {
		http.NotFound(w, r)
	}
}

// RegisterStaticPath registers the given url path such that it points to the given directory. For example, passing
// "/test", "/my/test/dir" will statically serve everything out of /my/test/dir whenever a url has /test in front of it.
// You can only call this during application startup. These directory paths take precedence over other similar paths that
// you have registered through goradd forms or through the html directory.
func RegisterStaticPath(path string, directory string) {
	if path[0:1] != "/" {
		log.Fatal("path must begin with a slash (must be a rooted path)")
	}

	if path[len(path)-1:] == "/" {
		// Strip ending slash so that we can handle both /a/b/ and /a/b urls as directories and treat them the same.
		path = path[:len(path)-1]
	}

	if StaticDirectoryPaths == nil {
		StaticDirectoryPaths = maps.NewStringSliceMap()
		// sort the directory paths longest to shortest so that when we iterate them, we won't short circuit
		// longer paths with shorter versions of the same path.
		StaticDirectoryPaths.SetSortFunc(func(key1,key2 string, val1, val2 string) bool {
			// order longest to shortest keys
			return len(key1) > len(key2)
		})
	}
	StaticDirectoryPaths.Set(path, directory)
}

// ServeApiHandler serves up an http API. This could be a REST api or something else.
func (a *Application) ServeApiHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if !a.this().ServeApiRequest(w, r) {
			next.ServeHTTP(w, r)
		}
	}
	return http.StripPrefix(config.ApiPrefix, http.HandlerFunc(fn))
}

// ServeApiRequest serves up an http api call. The prefix has been removed, so
// we just process the URL as if it were the command itself.
// This is currently just a stub to allow you to implement your own API. Eventually we hope this
// could be an auto-generated REST api or GraphQL api.
func (a *Application) ServeApiRequest(w http.ResponseWriter, r *http.Request) bool {
	// TODO
	//return rest.HandleRequest(w, r)	// indicates no static file was found
	return false
}

func RegisterStaticFileProcessor(ending string, processorFunc StaticFileProcessorFunc) {
	staticFileProcessors = append(staticFileProcessors, staticFileProcessor{ending, processorFunc})
}
