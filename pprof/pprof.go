package pprof

import (
	"net/http/pprof"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp/fasthttpadaptor"
)

// Set pprof adaptors
var (
	pprofIndex        = fasthttpadaptor.NewFastHTTPHandlerFunc(pprof.Index)
	pprofCmdline      = fasthttpadaptor.NewFastHTTPHandlerFunc(pprof.Cmdline)
	pprofProfile      = fasthttpadaptor.NewFastHTTPHandlerFunc(pprof.Profile)
	pprofSymbol       = fasthttpadaptor.NewFastHTTPHandlerFunc(pprof.Symbol)
	pprofTrace        = fasthttpadaptor.NewFastHTTPHandlerFunc(pprof.Trace)
	pprofAllocs       = fasthttpadaptor.NewFastHTTPHandlerFunc(pprof.Handler("allocs").ServeHTTP)
	pprofBlock        = fasthttpadaptor.NewFastHTTPHandlerFunc(pprof.Handler("block").ServeHTTP)
	pprofGoroutine    = fasthttpadaptor.NewFastHTTPHandlerFunc(pprof.Handler("goroutine").ServeHTTP)
	pprofHeap         = fasthttpadaptor.NewFastHTTPHandlerFunc(pprof.Handler("heap").ServeHTTP)
	pprofMutex        = fasthttpadaptor.NewFastHTTPHandlerFunc(pprof.Handler("mutex").ServeHTTP)
	pprofThreadcreate = fasthttpadaptor.NewFastHTTPHandlerFunc(pprof.Handler("threadcreate").ServeHTTP)
)

// New creates a new middleware handler
func New(config ...Config) fiber.Handler {
	// Set default config
	cfg := configDefault(config...)

	pprofIndexPath := cfg.DefaultPath + "/"
	pprofCmdlinePath := cfg.DefaultPath + "/cmdline"
	pprofProfilePath := cfg.DefaultPath + "/profile"
	pprofSymbolPath := cfg.DefaultPath + "/symbol"
	pprofTracePath := cfg.DefaultPath + "/trace"
	pprofAllocsPath := cfg.DefaultPath + "/allocs"
	pprofBlockPath := cfg.DefaultPath + "/block"
	pprofGoroutinePath := cfg.DefaultPath + "/goroutine"
	pprofHeapPath := cfg.DefaultPath + "/heap"
	pprofMutexPath := cfg.DefaultPath + "/mutex"
	pprofThreadcreatePath := cfg.DefaultPath + "/threadcreate"

	minSieze := len(cfg.DefaultPath)

	// Return new handler
	return func(c *fiber.Ctx) error {
		path := c.Path()
		// We are only interested in cfg.DefaultPath routes
		if len(path) < minSieze || !strings.HasPrefix(path, cfg.DefaultPath) {
			return c.Next()
		}
		// Switch to original path without stripped slashes
		switch path {
		case pprofIndexPath:
			pprofIndex(c.Context())
		case pprofCmdlinePath:
			pprofCmdline(c.Context())
		case pprofProfilePath:
			pprofProfile(c.Context())
		case pprofSymbolPath:
			pprofSymbol(c.Context())
		case pprofTracePath:
			pprofTrace(c.Context())
		case pprofAllocsPath:
			pprofAllocs(c.Context())
		case pprofBlockPath:
			pprofBlock(c.Context())
		case pprofGoroutinePath:
			pprofGoroutine(c.Context())
		case pprofHeapPath:
			pprofHeap(c.Context())
		case pprofMutexPath:
			pprofMutex(c.Context())
		case pprofThreadcreatePath:
			pprofThreadcreate(c.Context())
		default:
			// pprof index only works with trailing slash
			return c.Redirect(cfg.DefaultPath+"/", 302)
		}
		return nil
	}
}
