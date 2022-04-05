package rest

import (
	"log"
	"net/http"
	"net/http/pprof"

	"github.com/dinel13/thesis-ac/auth/domain"
	"github.com/julienschmidt/httprouter"
)

// middleware is used to intercept incoming HTTP calls and apply general functions upon them.
func middleware(n httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		log.Printf("HTTP request sent to %s from %s", r.URL.Path, r.RemoteAddr)

		// call registered handler
		n(w, r, ps)
	}
}

// wrapper will wrap an http.Handler function with the middleware function.
func wrapper(h http.Handler) httprouter.Handle {
	return middleware(func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		h.ServeHTTP(w, r)
	})
}

func Routes(ah domain.AuthRestHandlers) http.Handler {
	r := httprouter.New()

	r.HandlerFunc(http.MethodPost, "/verify", ah.Verify)
	r.HandlerFunc(http.MethodPost, "/login", ah.Login)
	r.HandlerFunc(http.MethodPost, "/signup", ah.Signup)

	r.GET("/debug/pprof/", wrapper(http.HandlerFunc(pprof.Index)))
	r.GET("/debug/pprof/cmdline", wrapper(http.HandlerFunc(pprof.Cmdline)))
	r.GET("/debug/pprof/profile", wrapper(http.HandlerFunc(pprof.Profile)))
	r.GET("/debug/pprof/symbol", wrapper(http.HandlerFunc(pprof.Symbol)))
	r.GET("/debug/pprof/trace", wrapper(http.HandlerFunc(pprof.Trace)))
	r.GET("/debug/pprof/allocs", wrapper(pprof.Handler("allocs")))
	r.GET("/debug/pprof/mutex", wrapper(pprof.Handler("mutex")))
	r.GET("/debug/pprof/goroutine", wrapper(pprof.Handler("goroutine")))
	r.GET("/debug/pprof/heap", wrapper(pprof.Handler("heap")))
	r.GET("/debug/pprof/threadcreate", wrapper(pprof.Handler("threadcreate")))
	r.GET("/debug/pprof/block", wrapper(pprof.Handler("block")))

	return r
}
