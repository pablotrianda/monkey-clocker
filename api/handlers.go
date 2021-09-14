package api

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/pablotrianda/monkey-clocker/clocker"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello!!n.n!!!"))
	clocker.NewClocker()
}

func wichID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	w.Write([]byte(id))
}
