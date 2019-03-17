package go_http

import (
	"net/http"
	"strconv"
	"time"
)

var SimpleServer = &http.Server{
	ReadTimeout:       600 * time.Second,
	ReadHeaderTimeout: 300 * time.Second,
	WriteTimeout:      600 * time.Second,
	IdleTimeout:       3600 * time.Second,
	MaxHeaderBytes:    http.DefaultMaxHeaderBytes,
}

type HttpServer interface {
	ListenAndServe() error
}

func NewSimpleHttpServer(port int, handler http.Handler) HttpServer {
	server := SimpleServer
	server.Addr = ":" + strconv.Itoa(port)
	server.Handler = handler

	return server
}

func SimpleStartHttpServer(port int, handler http.Handler) error {
	return NewSimpleHttpServer(port, handler).ListenAndServe()
}
