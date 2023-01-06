package router

import (
	"fmt"
	"log"
	"net/http"
	"time"

	logger "binancebot/log"
)

type stdRouter struct{}

type Settings struct {
	SslCert  string `json:"sslCert"`
	SslKey   string `json:"sslKey"`
	UseHTTP2 bool   `json:"useHTTP2" default:"false"`
}

var (
	settings Settings
)

func NewStdRouter(httpSettings Settings) Router {
	settings = httpSettings
	return &stdRouter{}
}

var (
	stdRouterDispatch = http.NewServeMux()
)

func (*stdRouter) HandleReq(uri string, proxyHandler func(writer http.ResponseWriter, request *http.Request)) {
	stdRouterDispatch.HandleFunc(uri, proxyHandler)
}

func (*stdRouter) Serve(port string) {
	server := &http.Server{
		Addr:         ":" + port,
		Handler:      stdRouterDispatch,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  15 * time.Second,
	}
	server.SetKeepAlivesEnabled(true)
	logger.Print(fmt.Sprintf("Listening on %s\n", server.Addr), "info")
	if settings.UseHTTP2 {
		log.Fatal(server.ListenAndServeTLS("./ssl/server.crt", "./ssl/server.key"))
		return
	}
	log.Fatal(server.ListenAndServe())
}
