package api

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/pablotrianda/monkey-clocker/clocker"
)

func newSchema(w http.ResponseWriter, r *http.Request) {
	res, err := io.ReadAll(r.Body)
	handleErr(err, w, http.StatusBadRequest)

	newClocker, err := clocker.NewClocker(res)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newClocker)

}

func sayHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("HELLO"))
}

func wichID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	w.Write([]byte(id))
}

func loginUser(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Header)
	w.Write([]byte("Heeellooo"))
}
