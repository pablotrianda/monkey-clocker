package api

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

type server struct {
	*http.Server
}

func (s *server) startServer() {
	log.Printf("starting server... on: %s", s.Addr)

	if err := s.ListenAndServe(); err != nil {
		log.Println("Error to start the server... :(")
		log.Println(err.Error())
	}
}

func newServer(port string, mux *chi.Mux) *server {
	s := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	return &server{s}
}
