package main

// Config конфигурация приложения
type Config struct {
	LogLevel          string `long:"log-level" description:"Log level: panic, fatal, warn or warning, info, debug" env:"CL_LOG_LEVEL" required:"true"`
	LogJSON           bool   `long:"log-json" description:"Enable force log format JSON" env:"CL_LOG_JSON"`
	HttpPrivateListen string `long:"http-private-listen" description:"Listening host:port for private http-server" env:"CL_HTTP_PRIVATE_LISTEN" required:"true"`
}
