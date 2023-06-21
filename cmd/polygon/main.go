package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"

	"github.com/jessevdk/go-flags"
	"github.com/valyala/fasthttp"
)

// main точка входа в приложение
func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("recovered from panic: %s\n", r)
		}
	}()

	var cfg Config
	parser := flags.NewParser(&cfg, flags.Default)
	_, err := parser.Parse()
	if err != nil {
		fmt.Printf("error on parse config: %v\n", err)
		os.Exit(1)
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()

		fmt.Printf("starting HTTP private server:%s\n", cfg.HttpPrivateListen)
		server := fasthttp.Server{
			Handler: func(req *fasthttp.RequestCtx) {
				req.SuccessString("text/html; charset=utf-8", "success...")
			},
		}

		go func() {
			<-ctx.Done()
			err := server.Shutdown()
			if err != nil {
				fmt.Printf("error on shutdown HTTP private server: %v\n", err)
			}
		}()

		err := server.ListenAndServe(cfg.HttpPrivateListen)
		if err != nil {
			fmt.Printf("error on listen and serve HTTP private server: %v\n", err)
		}

		cancel()
	}()

	wg.Wait()
	fmt.Println("completed...")
}
