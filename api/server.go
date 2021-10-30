package api

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-chi/chi"
)

type server struct {
	*http.Server
}

func (s *server) startServer() {
	log.Printf("starting server... on: %s", s.Addr)

	go func() {
		if err := s.ListenAndServe(); err != nil {
			log.Println("Error to start the server... :(")
			log.Println(err.Error())
		}
	}()

	s.gracefulShutdown()
}

func newServer(port string, mux *chi.Mux) *server {
	s := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	return &server{s}
}

func (s *server) gracefulShutdown() {
	quit := make(chan os.Signal, 1)

	signal.Notify(quit, os.Interrupt)
	sig := <-quit
	log.Printf("cmd is shutting down %s", sig.String())

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	s.SetKeepAlivesEnabled(false)
	if err := s.Shutdown(ctx); err != nil {
		log.Fatalf("could not gracefully shutdown the cmd %s", err.Error())
	}
	log.Printf("cmd stopped")
}
