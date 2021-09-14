package api

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi"
)

func TestApiRootRoute(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)

	if err != nil {
		t.Errorf("Error creating a new request: %v", err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(sayHello)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code. Expected: %d. Got: %d.", http.StatusOK, status)
	}

	if b, err := io.ReadAll(rr.Body); err == nil && string(b) != "HELLO" {
		t.Errorf("Handler returned body content. Expected: %s. Got: %s.", "HELLO", b)
	}
}

func TestApiGetByID(t *testing.T) {
	id := "112"
	path := fmt.Sprintf("/%s", id)
	req, err := http.NewRequest("GET", path, nil)

	if err != nil {
		t.Errorf("Error creating a new request: %v", err)
	}

	rr := httptest.NewRecorder()

	// Need to create a router that we can pass the request through so that the vars will be added to the context
	router := chi.NewRouter()
	router.HandleFunc("/{id}", wichID)
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code. Expected: %d. Got: %d.", http.StatusOK, status)
	}

	if b, err := io.ReadAll(rr.Body); err == nil && string(b) != id {
		t.Errorf("Handler returned body content. Expected: %s. Got: %s.", id, b)
	}
}
