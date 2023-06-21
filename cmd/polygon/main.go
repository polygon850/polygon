package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"sync"
	"time"

	"github.com/jessevdk/go-flags"
	"github.com/valyala/fasthttp"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// main точка входа в приложение
func main() {
	var cfg = new(Config)
	parser := flags.NewParser(cfg, flags.HelpFlag|flags.PassDoubleDash)
	_, err := parser.Parse()
	if err != nil {
		log.Fatal(fatalJsonLog("failed to parse config", err))
	}

	logger, err := initLogger()
	if err != nil {
		log.Fatal(fatalJsonLog("failed to init logger", err))
	}

	defer func() {
		if r := recover(); r != nil {
			logger.Error("recovered from panic", zap.Error(err))
		}
	}()

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()

		logger.Info("starting HTTP private server", zap.String("address", cfg.HttpPrivateListen))
		server := fasthttp.Server{
			Handler: func(req *fasthttp.RequestCtx) {
				req.SuccessString("text/html; charset=utf-8", "success...")
			},
		}

		go func() {
			<-ctx.Done()
			err := server.Shutdown()
			if err != nil {
				logger.Error("error on shutdown HTTP private server", zap.Error(err))
			}
		}()

		err := server.ListenAndServe(cfg.HttpPrivateListen)
		if err != nil {
			logger.Error("error on listen and serve HTTP private server", zap.Error(err))
		}

		cancel()
	}()

	wg.Wait()
	logger.Info("application completed")
}

// initLogger создает и настраивает новый экземпляр логгера
func initLogger() (*zap.Logger, error) {
	var cfg zap.Config
	cfg = zap.NewProductionConfig()
	cfg.EncoderConfig.EncodeTime = zapcore.RFC3339NanoTimeEncoder

	return cfg.Build()
}

// fatalJsonLog стилизует сообщение ошибки
func fatalJsonLog(msg string, err error) string {
	escape := func(s string) string {
		return strings.ReplaceAll(strings.ReplaceAll(s, `\`, `\\`), `"`, `\"`)
	}
	errString := ""
	if err != nil {
		errString = err.Error()
	}
	return fmt.Sprintf(
		`{"level":"fatal","ts":"%s","msg":"%s","error":"%s"}`,
		time.Now().Format(time.RFC3339Nano),
		escape(msg),
		escape(errString),
	)
}
