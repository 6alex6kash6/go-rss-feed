package httpserver

import (
	"net/http"

	"log/slog"

	"github.com/go-chi/chi/v5"
)

type Server struct {
	server *http.Server
}

func New(handler *chi.Mux, port string) *Server {
	httpServer := &http.Server{
		Addr:    ":" + port,
		Handler: handler,
	}
	return &Server{
		server: httpServer,
	}
}

func (s *Server) Start() {
	slog.Info("Server started on port:", s.server.Addr)
	s.server.ListenAndServe()
}
