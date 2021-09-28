package api

import (
	"log"
	"net/http"
)

func handleErr(w http.ResponseWriter, err error, statusError int) {
	log.Println(err)
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte("Error to create a new clocker schema"))
}
