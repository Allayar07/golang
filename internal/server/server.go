package server

import (
	"net/http"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:    ":" + port,
		Handler: handler,
		//MaxHeaderBytes:    2 << 20,
		//ReadHeaderTimeout: 10 * time.Second,
		//WriteTimeout:      10 * time.Second,
	}

	return s.httpServer.ListenAndServe()
}
