package httpserver

import (
	"context"
	"github.com/inspectorvitya/note-storage/internal/application"
	"net"
	"net/http"
	"time"
)

type Server struct {
	App        *application.App
	HTTPServer *http.Server
	router     *http.ServeMux
}

func New(port string, app *application.App) *Server {
	mux := http.NewServeMux()
	middleware := Logging(mux)
	server := &Server{
		HTTPServer: &http.Server{
			Addr:         net.JoinHostPort("", port),
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 5 * time.Second,
			Handler:      middleware,
		},
		router: mux,
		App:    app,
	}
	return server
}

func (s *Server) Start() error {
	s.router.HandleFunc("/", s.Main)
	s.router.HandleFunc("/last", s.GetLast)
	s.router.HandleFunc("/first", s.GetFirs)
	return s.HTTPServer.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	err := s.HTTPServer.Shutdown(ctx)
	if err != nil {
		return err
	}
	return nil
}
