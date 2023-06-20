package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"

	"github.com/valyala/fasthttp"
)

const (
	// httpAddr адрес HTTP сервера
	httpAddr string = ":8080"
)

// main точка входа в приложение
func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("recovered from panic: %s\n", r)
		}
	}()

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()

		fmt.Printf("starting HTTP server:%s\n", httpAddr)
		server := fasthttp.Server{
			Handler: func(req *fasthttp.RequestCtx) {
				req.SuccessString("text/html; charset=utf-8", "success...")
			},
		}

		go func() {
			<-ctx.Done()
			err := server.Shutdown()
			if err != nil {
				fmt.Printf("error on shutdown HTTP server: %v\n", err)
			}
		}()

		err := server.ListenAndServe(httpAddr)
		if err != nil {
			fmt.Printf("error on listen and serve HTTP server: %v\n", err)
		}

		cancel()
	}()

	wg.Wait()
	fmt.Println("completed...")
}
