package api

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/go-chi/chi"
)

type server struct {
	*http.Server
}

func (s *server) startServer() {
	log.Printf("starting server... on: %s", s.Addr)

	idleConnsClosed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint

		// We received an interrupt signal, shut down.
		if err := s.Shutdown(context.Background()); err != nil {
			// Error from closing listeners, or context timeout:
			log.Printf("HTTP server Shutdown: %v", err)
		}

		close(idleConnsClosed)
	}()

	// Start the server
	if err := s.ListenAndServe(); err != http.ErrServerClosed {
		log.Println("Error to start the server... :(")
	}

	<-idleConnsClosed
}

func newServer(port string, mux *chi.Mux) *server {
	s := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	return &server{s}
}
