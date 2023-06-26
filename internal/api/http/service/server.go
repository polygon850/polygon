package service

import (
	"context"
	"net/http/pprof"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttpadaptor"
	"go.uber.org/zap"
)

const pprofUrlPrefix = "/debug/pprof"

// ListenAndServe стартует служебный HTTP сервер
func ListenAndServe(ctx context.Context, addr string, enablePprof bool, logger *zap.Logger) error {
	r := router.New()
	r.GET("/ping", func(ctx *fasthttp.RequestCtx) {
		ctx.SuccessString("text/html; charset=utf-8", "PONG")
	})
	if enablePprof {
		for _, path := range []string{"/", "/allocs", "/block", "/goroutine", "/heap", "/mutex", "/threadcreate"} {
			r.GET(pprofUrlPrefix+path, fasthttpadaptor.NewFastHTTPHandlerFunc(pprof.Index))
		}

		r.GET(pprofUrlPrefix+"/cmdline", fasthttpadaptor.NewFastHTTPHandlerFunc(pprof.Cmdline))
		r.GET(pprofUrlPrefix+"/profile", fasthttpadaptor.NewFastHTTPHandlerFunc(pprof.Profile))
		r.GET(pprofUrlPrefix+"/symbol", fasthttpadaptor.NewFastHTTPHandlerFunc(pprof.Symbol))
		r.GET(pprofUrlPrefix+"/trace", fasthttpadaptor.NewFastHTTPHandlerFunc(pprof.Trace))
		logger.Warn("pprof routes registered", zap.String("prefix", pprofUrlPrefix))
	}

	server := fasthttp.Server{
		Handler: r.Handler,
	}

	go func() {
		<-ctx.Done()
		err := server.Shutdown()
		if err != nil {
			logger.Error("error on shutdown HTTP private server", zap.Error(err))
		}
	}()

	return server.ListenAndServe(addr)
}
