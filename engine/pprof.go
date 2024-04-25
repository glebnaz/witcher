package engine

import (
	"net/http"
	"net/http/pprof"
)

func wrapPProf(mux *http.ServeMux) {
	routers := []struct {
		Path    string
		Handler http.HandlerFunc
	}{
		{"/debug/pprof/", pprof.Index},
		{"/debug/pprof/heap", pprof.Handler("heap").ServeHTTP},
		{"/debug/pprof/goroutine", pprof.Handler("goroutine").ServeHTTP},
		{"/debug/pprof/block", pprof.Handler("block").ServeHTTP},
		{"/debug/pprof/threadcreate", pprof.Handler("threadcreate").ServeHTTP},
		{"/debug/pprof/cmdline", pprof.Cmdline},
		{"/debug/pprof/profile", pprof.Profile},
		{"/debug/pprof/symbol", pprof.Symbol},
		{"/debug/pprof/trace", pprof.Trace},
		{"/debug/pprof/mutex", pprof.Handler("mutex").ServeHTTP},
		{"/debug/pprof/allocs", pprof.Handler("allocs").ServeHTTP},
	}

	for _, r := range routers {
		mux.HandleFunc(r.Path, r.Handler)
	}
}
