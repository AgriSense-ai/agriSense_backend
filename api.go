package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type APIServer struct {
	listenAddr string
	store      Storage
}

func NewAPIServer(listenAddr string, store Storage) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
		store:      store,
	}
}

func (s *APIServer) Run() {
	router := mux.NewRouter()

	router.HandleFunc("/account", makeHTTPHandler(s.handleAccount))
	router.HandleFunc("/account/{id}", makeHTTPHandler(s.handleGetAccountWithID))

	log.Printf("API server listening on %s", s.listenAddr)

	http.ListenAndServe(s.listenAddr, router)
}

func (s *APIServer) handleAccount(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return s.handleGetAccount(w)
	}
	if r.Method == "POST" {
		return s.handleCreateAccount(w, r)
	}
	if r.Method == "DELETE" {
		return s.handleDeleteAccount(w, r)
	}
	return fmt.Errorf("not implemented %s", r.Method)
}

func (s *APIServer) handleGetAccountWithID(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return s.handleGetAccountByID(w, r)
	}
	if r.Method == "DELETE" {
		return s.handleDeleteAccount(w, r)
	}
	return fmt.Errorf("not implemented %s", r.Method)
}

func (s *APIServer) handleGetAccount(w http.ResponseWriter) error {
	accounts, err := s.store.GetAccounts()
	if err != nil {
		return WriteJSON(w, http.StatusInternalServerError, APIError{Error: err.Error()})
	}

	return WriteJSON(w, http.StatusOK, accounts)
}

func (s *APIServer) handleGetAccountByID(w http.ResponseWriter, r *http.Request) error {
	id_valid, err := getID(r)
	if err != nil {
		return WriteJSON(w, http.StatusBadRequest, APIError{Error: err.Error()})
	}
	acc, err := s.store.GetAccountByID(id_valid)
	if err != nil {
		return WriteJSON(w, http.StatusNotFound, APIError{Error: err.Error()})
	}
	return WriteJSON(w, http.StatusOK, acc)

}

func (s *APIServer) handleCreateAccount(w http.ResponseWriter, r *http.Request) error {
	createAccountReq := new(CreateAccountRequest)
	if err := json.NewDecoder(r.Body).Decode(createAccountReq); err != nil {
		return err
	}

	account := NewAccount(createAccountReq.FirstName, createAccountReq.LastName)

	if err := s.store.CreateAccount(account); err != nil {
		return err
	}

	return WriteJSON(w, http.StatusCreated, account)
}

func (s *APIServer) handleDeleteAccount(w http.ResponseWriter, r *http.Request) error {
	id_valid, err := getID(r)
	if err != nil {
		return WriteJSON(w, http.StatusBadRequest, APIError{Error: err.Error()})
	}
	if err := s.store.DeleteAccount(id_valid); err != nil {
		return WriteJSON(w, http.StatusNotFound, APIError{Error: err.Error()})
	}

	return WriteJSON(w, http.StatusAccepted, map[string]int{"status": id_valid})
}

func (s *APIServer) handleTransfer(w http.ResponseWriter, r *http.Request) error {

	return nil
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

type apiFunc func(http.ResponseWriter, *http.Request) error

type APIError struct {
	Error string
}

func makeHTTPHandler(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJSON(w, http.StatusInternalServerError, APIError{Error: err.Error()})
		}
	}
}

func getID(r *http.Request) (int, error) {
	id := mux.Vars(r)["id"]
	id_valid, err := strconv.Atoi(id)
	if err != nil {
		return 0, err
	}
	return id_valid, nil
}
