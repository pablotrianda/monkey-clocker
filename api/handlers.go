package api

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/pablotrianda/monkey-clocker/clocker"
)

func newSchema(w http.ResponseWriter, r *http.Request) {
	var schema clocker.Clocker
	err := json.NewDecoder(r.Body).Decode(&schema)
	if err != nil {
		handleErr(w, err, http.StatusBadRequest)
	}
	clocker.NewClocker(schema)
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("HELLO"))
}

func wichID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	w.Write([]byte(id))
}
