package httpServer

import (
	"fmt"
	"net/http"
)

type HTTPServer struct {
	addr    string
	handler http.Handler
}

func NewHTTPServer(addr string, handler http.Handler) *HTTPServer {
	return &HTTPServer{
		addr:    addr,
		handler: handler,
	}
}

func (h *HTTPServer) Run() error {
	fmt.Printf("HTTP server listening on %s\n", h.addr)
	return http.ListenAndServe(h.addr, h.handler)
}
