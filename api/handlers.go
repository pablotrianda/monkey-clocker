package api

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/pablotrianda/monkey-clocker/clocker"
)

type Response struct {
	Name      string
	BusyHours string `json:"busy_hours"`
}

func newSchema(w http.ResponseWriter, r *http.Request) {
	var res Response
	err := json.NewDecoder(r.Body).Decode(&res)
	handleErr(err, w, http.StatusBadRequest)
	busyHours, err := convertStringArrayToIntArray(res.BusyHours)
	handleErr(err, w, http.StatusInternalServerError)

	newClocker, err := clocker.NewClocker(res.Name, busyHours)
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
