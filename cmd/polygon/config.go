package main

// Config конфигурация приложения
type Config struct {
	HttpPrivateListen string `long:"http-private-listen" description:"Listening host:port for private http-server" env:"CL_HTTP_PRIVATE_LISTEN" required:"true"`
}
