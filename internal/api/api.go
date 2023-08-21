package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Abdulrahman-02/Go-bank/internal/types"
	"github.com/gorilla/mux"
)

type APIserver struct {
	listenAddr string
}

type APIerror struct {
	Error string
}

func writeJSON(w http.ResponseWriter, status int, v any) error {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}

// apiFunc is a function that handles an API request and returns an error
type apiFunc func(http.ResponseWriter, *http.Request) error

// makeHTTPHandleFunc creates an HTTP handler from an API function
func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			// handle the error
			writeJSON(w, http.StatusBadRequest, APIerror{Error: err.Error()})
		}
	}
}

// NewAPIServer creates a new API server
func NewAPIserver(listenAddr string) *APIserver {
	return &APIserver{listenAddr: listenAddr}
}

// Run starts the API server
func (s *APIserver) Run() {
	router := mux.NewRouter()
	router.HandleFunc("/account", makeHTTPHandleFunc(s.handleAccount))

	log.Println("JSON API server running on port: ", s.listenAddr)
	http.ListenAndServe(s.listenAddr, router)
}

// handleAccount handles the account API
func (s *APIserver) handleAccount(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return s.handleGetAccount(w, r)
	}

	if r.Method == "POST" {
		return s.handleCreateAccount(w, r)
	}

	if r.Method == "DELETE" {
		return s.handleDeleteAccount(w, r)
	}

	return fmt.Errorf("method not allowed %s", r.Method)
}

// handleGetAccount handles the get account API
func (s *APIserver) handleGetAccount(w http.ResponseWriter, r *http.Request) error {
	account := types.NewAccount("Abdou", "Ben")
	return writeJSON(w, http.StatusOK, account)
}

// handleCreateAccount handles the create account API
func (s *APIserver) handleCreateAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

// handleDeleteAccount handles the delete account API
func (s *APIserver) handleDeleteAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

// handleTransfer handles the transfer API
func (s *APIserver) handleTransfer(w http.ResponseWriter, r *http.Request) error {
	return nil
}
