package main

// Config конфигурация приложения
type Config struct {
	LogLevel          string `long:"log-level" description:"Log level: panic, fatal, warn or warning, info, debug" env:"CL_LOG_LEVEL" required:"true"`
	LogJSON           bool   `long:"log-json" description:"Enable force log format JSON" env:"CL_LOG_JSON"`
	HttpServiceListen string `long:"http-service-listen" description:"Listening host:port for service http-server" env:"CL_HTTP_SERVICE_LISTEN" required:"true"`
	EnablePprof       bool   `long:"enable-pprof" description:"Enable pprof server" env:"CL_ENABLE_PPROF"`
}
